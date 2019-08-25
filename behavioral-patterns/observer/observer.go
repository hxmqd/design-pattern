// 当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。 比如，当一个对象被修改时，则会自动通知它的依赖对象。观察者模式属于行为型模式
// 意图：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
// 主要解决：一个对象状态改变给其他对象通知的问题，而且要考虑到易用和低耦合，保证高度的协作。

package observer

import "fmt"

type Subject struct {
	observers []IObserver
	state int
}

func (s *Subject) GetState() int{
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()

}

func (s *Subject) NotifyAllObservers(){
	for _, observer := range s.observers {
		observer.Update()
	}
}

type IObserver interface {
	Update()
}

type BinaryObserver struct {
	subject *Subject
}

func NewBinaryObserver(subject *Subject) *BinaryObserver{
	observer := &BinaryObserver{
		subject: subject,
	}
	subject.observers = append(subject.observers, observer)
	return observer
}

func (b *BinaryObserver) Update(){
	fmt.Printf("binary string:%d \n", b.subject.GetState())
}

type OctalObserver struct {
	subject *Subject
}

func NewOctalObserver(subject *Subject) *OctalObserver{
	observer := &OctalObserver{
		subject: subject,
	}
	subject.observers = append(subject.observers, observer)
	return observer
}

func (b *OctalObserver) Update(){
	fmt.Printf("Octal string:%d \n", b.subject.GetState())
}

type HexaObserver struct {
	subject *Subject
}

func NewHexaObserver(subject *Subject) *HexaObserver{
	observer := &HexaObserver{
		subject: subject,
	}
	subject.observers = append(subject.observers, observer)
	return observer
}

func (b *HexaObserver) Update(){
	fmt.Printf("Hexa string:%d \n", b.subject.GetState())
}

