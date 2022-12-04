package parallel

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"sync"
	"time"
)

// 信号量：Weighted
// 我们创建和 CPU 核数一样多的 Worker，让它们去处理一个 4 倍数量的整数 slice。每个 Worker 一次只能处理一个整数，处理完之后，才能处理下一个
var (
	maxWorkers = runtime.GOMAXPROCS(0)                    // worker数量
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)                // 任务数，是worker的四倍
)

func WeightedDemo() {
	ctx := context.Background()

	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		// 启动worker goroutine
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}

	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v", err)
	}

	fmt.Println(task)
}

// Semaphore 数据结构，并且还实现了Locker接口
// 在初始化这个信号量的时候，我们设置它的初始容量，代表有多少个资源可以使用。它使用 Lock 和 Unlock 方法实现请求资源和释放资源，正好实现了 Locker 接口
type Semaphore struct {
	sync.Locker
	ch chan struct{}
}

// NewSemaphore 创建一个新的信号量
func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &Semaphore{ch: make(chan struct{}, capacity)}
}

// Lock 请求一个资源
func (s *Semaphore) Lock() {
	s.ch <- struct{}{}
}

// Unlock 释放资源
func (s *Semaphore) Unlock() {
	<-s.ch
}
