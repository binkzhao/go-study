package parallel

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

type Token struct{}

// StartGame 数据传递“击鼓传花”的游戏======================================
func StartGame() {
	newWorker := func(id int, ch chan Token, nextCh chan Token) {
		for {
			token := <-ch
			fmt.Println(id + 1)
			time.Sleep(time.Second / 10)
			nextCh <- token
		}
	}
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
	for i, ch := range chs {
		go newWorker(i, ch, chs[(i+1)%4])
	}
	chs[0] <- Token{}
	select {}
}

// Notify 信号通知======================================
// 有时候，doCleanup 可能是一个很耗时的操作，比如十几分钟才能完成，如果程序退出需要等待这么长时间，用户是不能接受的，
// 所以，在实践中，我们需要设置一个最长的等待时间。只要超过了这个时间，程序就不再等待，可以直接退出。
// 所以，退出的时候分为两个阶段：closing，代表程序退出，但是清理工作还没做；closed，代表清理工作已经做完。
func Notify() {
	doCleanup := func(closed chan struct{}) {
		time.Sleep((time.Minute))
		close(closed)
	}

	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

// Mutex 使用chan实现互斥锁======================================
type Mutex struct {
	ch chan struct{}
}

// NewMutex 使用锁需要初始化
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// Lock 请求锁，直到获取到
func (m *Mutex) Lock() {
	<-m.ch
}

// Unlock 解锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// LockTimeout 加入一个超时的设置
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

// IsLocked 锁是否已被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func demoMutex() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}

// Or-Done 模式
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()

	return orDone
}

func or2(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		reflect.Select(cases)
	}()

	return orDone
}

// fanInReflect 扇入模式==================================================
// 在软件工程中，模块的扇入是指有多少个上级模块调用它。
// 而对于我们这里的 Channel 扇入模式来说，就是指有多个源 Channel 输入、一个目的 Channel 输出的情况。
// 扇入比就是源 Channel 数量比 1。每个源 Channel 的元素都会发送给目标 Channel，相当于目标 Channel 的 receiver
// 只需要监听目标 Channel，就可以接收所有发送给源 Channel 的数据。
func fanInReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}

// fanInRec2 递归模式也是在 Channel 大于 2 时，采用二分法递归 merge
func fanInRec2(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec2(chans[:m]...),
			fanInRec2(chans[m:]...))
	}
}

func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil { //只要还有可读的chan
			select {
			case v, ok := <-a:
				if !ok { // a 已关闭，设置为nil
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok { // b 已关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

// 扇出模式=================================================
// 扇出模式只有一个输入源 Channel，有多个目标 Channel，扇出比就是 1 比目标 Channel 数的值，经常用在设计模式中的观察者模式中
//（观察者设计模式定义了对象间的一种一对多的组合关系。这样一来，一个对象的状态发生变化时，所有依赖于它的对象都会得到通知并自动刷新）。
// 在观察者模式中，数据变动后，多个观察者都会收到这个变更信号。
// 下面是一个扇出模式的实现。从源 Channel 取出一个数据后，依次发送给目标 Channel。在发送给目标 Channel 的时候，可以同步发送，也可以异步发送
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { //异步
					go func() {
						out[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}
