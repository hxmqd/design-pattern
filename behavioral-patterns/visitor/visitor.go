// 访问者模式（Visitor Pattern）使用了一个访问者类，改变了元素类的执行算法。通过这种方式，元素的执行算法可以随着访问者改变而改变。
// 这种类型的设计模式属于行为型模式。根据模式，元素对象已接受访问者对象，这样访问者对象就可以处理元素对象上的操作。
// 意图：主要将数据结构与数据操作分离。
// 主要解决：稳定的数据结构和易变的操作耦合问题。
// 何时使用：需要对一个对象结构中的对象进行很多不同的并且不相关的操作， 而需要避免让这些操作"污染"这些对象的类，使用访问者模式将这些封装到类中。
// 优点： 1、符合单一职责原则。 2、优秀的扩展性。 3、灵活性。
// 缺点： 1、具体元素对访问者公布细节，违反了迪米特原则。 2、具体元素变更比较困难。 3、违反了依赖倒置原则，依赖了具体类，没有依赖抽象。
package visitor

import "fmt"

type IComputerPartVisitor interface {
	ComputerVisit(computer *Computer)
	MouseVisit(mouse *Mouse)
	KeyboardVisit(Keyboard *Keyboard)
	MonitorVisit(monitor  *Monitor)
}

type ComputerPartVisitor struct {
}

func (c *ComputerPartVisitor) ComputerVisit(computer *Computer){
	fmt.Println("Displaying Computer.")
}

func (c *ComputerPartVisitor) MouseVisit(mouse *Mouse){
	fmt.Println("Displaying mouse.")
}

func (c *ComputerPartVisitor) KeyboardVisit(Keyboard *Keyboard){
	fmt.Println("Displaying keyboard.")
}

func (c *ComputerPartVisitor) MonitorVisit(monitor *Monitor){
	fmt.Println("Displaying monitor.")
}

type ComputerPart interface {
	Accept(computerPartVisitor IComputerPartVisitor)
}

type Keyboard struct {}

func (k *Keyboard) Accept(computerPartVisitor IComputerPartVisitor){
	computerPartVisitor.KeyboardVisit(k)
}

type Monitor struct {}

func (m *Monitor) Accept(computerPartVisitor IComputerPartVisitor){
	computerPartVisitor.MonitorVisit(m)
}

type Mouse struct {}

func (m *Mouse) Accept(computerPartVisitor IComputerPartVisitor){
	computerPartVisitor.MouseVisit(m)
}

type Computer struct {
	parts []ComputerPart
}

func NewComputer(parts []ComputerPart) *Computer{
	return &Computer{parts: parts}
}
func (c *Computer) Accept(computerPartVisitor IComputerPartVisitor){
	for i := 0; i < len(c.parts); i++ {
		c.parts[i].Accept(computerPartVisitor)
	}
	computerPartVisitor.ComputerVisit(c)
}







