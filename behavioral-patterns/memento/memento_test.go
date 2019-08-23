package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := NewOriginator()
	careTaker := NewCareTaker()
	originator.SetState("state #1")
	originator.SetState("state #2")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("state #3")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("state #4")


	fmt.Println("current state :" + originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("first saved state:"+ originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("second saved state:" + originator.GetState())
}
