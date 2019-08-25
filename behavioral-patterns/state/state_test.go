package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T){
	context := &Context{}

	startState := StartState{}
	startState.DoAction(context)
	fmt.Printf("state:%s \n", context.GetState().ToString())

	StopState := StopState{}
	StopState.DoAction(context)
	fmt.Printf("state:%s \n", context.GetState().ToString())
}
