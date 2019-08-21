package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	pattern := &InterpreterPattern{}
	isMale := pattern.GetMaleExpression()
	isMarriedWoman := pattern.GetMarriedWomanExpression()

	t.Logf("John is male? %v \n" , isMale.Interpret("John"))
	t.Logf("Julie is a married women? %v \n" , isMarriedWoman.Interpret("Married Julie"))

}