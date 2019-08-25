package template_methodbn

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {

	game := NewCricket()
	game.Play()
	fmt.Println("")
	game = NewFootball()
	game.Play()
}
