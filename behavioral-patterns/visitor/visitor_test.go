package visitor

import "testing"

func TestVisitor(t *testing.T) {
	parts := []ComputerPart{new(Keyboard), new(Monitor), new(Mouse)}

	var computer ComputerPart
	computer = NewComputer(parts)
	computer.Accept(&ComputerPartVisitor{})
}
