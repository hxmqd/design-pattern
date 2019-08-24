package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestObserver(t *testing.T){
	subject := &Subject{}

	_ = NewBinaryObserver(subject)
	_ = NewOctalObserver(subject)
	_ = NewHexaObserver(subject)

	subject.SetState(15)
	time.Sleep(1000)
	fmt.Print("\nsubject change state\n\n")
	subject.SetState(10)
}
