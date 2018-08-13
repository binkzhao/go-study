package patterns

import "sync"

// 线程安全的单例模式
type single struct {
	O interface{}
}

var instantiated *single
var once sync.Once // Once对象在全局只会执行一次，来保证线程安全

func New() *single {
	// 如果有多个once.Do,也是只会执行一次
	once.Do(func() {
		instantiated = &single{}
	})
	return instantiated
}