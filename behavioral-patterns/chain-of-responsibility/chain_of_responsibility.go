// 责任链模式为请求创建了一个接收者对象链。这种模式给予请求的类型，对请求的发送者和接收者进行解耦。 这种类型的设计模式属于行为型模式。
// 避免请求发送者与接收者耦合在一起，让多个对象都有可能接收请求，将这些对象连接成一条链，并且沿着这条链传递请求，直到有对象处理它为止
// 主要解决：职责链上的处理者负责处理请求，客户只需要将请求发送到职责链上即可，无须关心请求的处理细节和请求的传递， 所以职责链将请求的发送者和请求的处理者解耦了。

package chain_of_responsibility

import "fmt"

const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

type Logger interface {
	SetNextLogger(nextLogger Logger)
	LogMessage(level int, message string)
	Write(message string)
}

// 抽象记录器结构体
type AbstractLogger struct {
	level int
	nextLogger Logger
}

func (a *AbstractLogger)setNextLogger(nextLogger Logger){
	a.nextLogger = nextLogger
}

func (a *AbstractLogger)LogMessage(level int, message string){
	if a.level < level {
		a.Write(message)
	}
	if a.nextLogger != nil {
		a.nextLogger.LogMessage(level, message)
	}
}

func (a *AbstractLogger)Write(message string){}

// 控制台日志记录器
type ConsoleLogger struct {
	AbstractLogger
}

func (a *ConsoleLogger)Write(message string){
	fmt.Println("standard console :: logger:"+message)
}

func NewConsoleLogger(level int) *ConsoleLogger{
	return &ConsoleLogger{AbstractLogger{level:level}}
}

// 错误日志记录器
type ErrorLogger struct {
	AbstractLogger
}

func (a *ErrorLogger)Write(message string){
	fmt.Println("error console :: logger:"+message)
}

func NewErrorLogger(level int) *ErrorLogger{
	return &ErrorLogger{AbstractLogger{level:level}}
}

// 错误日志记录器
type FileLogger struct {
	AbstractLogger
}

func (a *FileLogger)Write(message string){
	fmt.Println("file :: logger:"+message)
}

func NewFileLogger(level int) *FileLogger{
	return &FileLogger{AbstractLogger{level:level}}
}





