// 备忘录模式（Memento Pattern）保存一个对象的某个状态，以便在适当的时候恢复对象。备忘录模式属于行为型模式。
// 主要解决：所谓备忘录模式就是在不破坏封装的前提下，捕获一个对象的内部状态，
// 并在该对象之外保存这个状态，这样可以在以后将对象恢复到原先保存的状态。

package memento

type Memento struct {
	state string
}

func NewMemento(state string) *Memento{
	return &Memento{
		state: state,
	}
}

func (m *Memento)GetState() string{
	return m.state
}

type Originator struct {
	state string
}

func NewOriginator() *Originator{
	return &Originator{}
}

func (o *Originator)SetState(state string){
	o.state = state
}

func (o *Originator)GetState()string{
	return o.state
}

func (o *Originator)SaveStateToMemento() *Memento{
	return &Memento{state:o.state}
}

func (o *Originator)GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
}

type CareTaker struct {
	mementoList []*Memento
}

func NewCareTaker() *CareTaker{
	return &CareTaker{
		mementoList: make([]*Memento, 0),
	}
}

func (c *CareTaker) Add(state *Memento){
	c.mementoList = append(c.mementoList, state)
}

func (c *CareTaker) Get(index int) *Memento{
	return c.mementoList[index]
}
