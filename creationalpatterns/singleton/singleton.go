// 单例模式
package singleton

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Single struct {
	i int
}

func (s *Single) ShowMessage() {
	fmt.Println("hello world!","i=",s.i)
}

func newSingle(i int) *Single{
	time.Sleep(time.Duration(rand.Intn(100)))
	return &Single{i}
}

// 懒汉式，协程不安全
var single1 *Single

func GetInstance1(i int) *Single{
	if single1 == nil {
		single1 = newSingle(i)
	}
	return single1
}

// 懒汉式，协程安全
var single2 *Single
var mux sync.Mutex
func GetInstance2(i int) *Single {
	mux.Lock()
	if single1 == nil {
		single1 = newSingle(i)
	}
	mux.Unlock()
	return single1
}

// 饿汉式, 协程安全
 var single3 = new(Single)
 func GetInstance3() *Single{
 	return single3
 }

 // 双检锁/双重校验锁(double-checked locking)
var single4 *Single
var mux1 sync.Mutex
func GetInstance4(i int) *Single {
	if single4 == nil {
		mux1.Lock()
		if single4 == nil {
			single4 = newSingle(i)
		}
		mux1.Unlock()
	}
	return single4
}

// sync.Once实现
var single5 *Single
var once sync.Once
func GetInstance5() *Single {
	once.Do(func(){
		single5 = &Single{}
	})
	return single5
}




