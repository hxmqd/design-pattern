// 适配器模式（Adapter Pattern）作为两个不兼容的接口之间的桥梁， 将一个类的接口转换成客户希望的另外一个接口
// 适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
package main

import "log"

type Logger interface {
	Info(s string)
	Warn(s string)
}

// 假设你写了一个框架,框架里需要打日志，具体是个什么样的日志由使用者决定
type Frame struct {
	name string
	log Logger
}

func (f *Frame) start(){
	f.log.Info("frame start success")
}

// 但是问题来了，使用者用了一个这样的日志库
type RealLog struct {
}

func (l *RealLog) WARN(s string) {
	log.Print(s)
}

func (l *RealLog) INFO(s string) {
	log.Print(s)
}

// 可以看到，真正使用的日志与框架定义的日志接口不一样，如此我们就需要适配
type AdaptLog struct {
	relLog *RealLog
}

func (a *AdaptLog) Warn(s string) {
	a.relLog.WARN(s)
}

func (a *AdaptLog) Info(s string) {
	a.relLog.INFO(s)
}

// 有了这个适配器，我们就可以很好的在框架里使用RealLog了
func main()  {
	adaptLog := AdaptLog{&RealLog{}}
	frame := Frame{name:"frame", log:&adaptLog}
	frame.start()
}






