// 在状态模式（State Pattern）中，类的行为是基于它的状态改变的。这种类型的设计模式属于行为型模式。
// 在状态模式中，我们创建表示各种状态的对象和一个行为随着状态对象改变而改变的 context 对象。
// 意图：允许对象在内部状态发生改变时改变它的行为，对象看起来好像修改了它的类。
// 主要解决：对象的行为依赖于它的状态（属性），并且可以根据它的状态改变而改变它的相关行为。

package state

import "fmt"

type State interface {
	DoAction(ctx *Context)
	ToString() string
}

type StartState struct {
}

func (s *StartState) DoAction(ctx *Context){
	fmt.Println("player is in start state")
	ctx.SetState(s)
}

func (s *StartState) ToString() string{
	return "stop state"
}

type StopState struct {
}

func (s *StopState) DoAction(ctx *Context){
	fmt.Println("player is in stop state")
	ctx.SetState(s)
}

func (s *StopState) ToString() string{
	return "stop state"
}

type Context struct {
	state State
}

func (ctx *Context) SetState(state State){
	ctx.state = state
}

func (ctx *Context) GetState() State{
	return ctx.state
}
