package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	context := NewContext(OperationAdd{})
	t.Logf("10+5=%d", context.ExecuteStrategy(10, 5))

	context = NewContext(OperationSubstract{})
	t.Logf("10-5=%d", context.ExecuteStrategy(10, 5))

	context = NewContext(OperationMultiply{})
	t.Logf("10*5=%d", context.ExecuteStrategy(10, 5))
}