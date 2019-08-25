// 在策略模式（Strategy Pattern）中，一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。
// 意图：定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。
// 主要解决：在有多种算法相似的情况下，使用 if else 所带来的复杂和难以维护。
// 何时使用：一个系统有许多许多类，而区分它们的只是他们直接的行为。

// 状态模式的类图和策略模式类似，并且都是能够动态改变对象的行为。但是状态模式是通过状态转移来改变 Context 所组合的 State 对象，
// 而策略模式是通过 Context 本身的决策来改变组合的 Strategy 对象。
// 所谓的状态转移，是指 Context 在运行过程中由于一些条件发生改变而使得 State 对象发生改变。
// 状态模式主要是用来解决状态转移的问题，当状态发生转移了，那么 Context 对象就会改变它的行为。
// 策略模式主要是用来封装一组可以互相替代的算法族，并且可以根据需要动态地去替换 Context 使用的算法。

package strategy

type Strategy interface {
	DoOperation(num1, num2 int) int
}

type OperationAdd struct {
}

func (o OperationAdd) DoOperation(num1, num2 int) int{
	return num1 + num2
}

type OperationSubstract struct {
}

func (o OperationSubstract) DoOperation(num1, num2 int) int{
	return num1 - num2
}

type OperationMultiply struct {
}

func (o OperationMultiply) DoOperation(num1, num2 int) int{
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context{
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

