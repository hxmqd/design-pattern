package chain_of_responsibility

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	loggerChain := GetChainOfLoggers()
	loggerChain.LogMessage(INFO, "this is an information.")

	fmt.Println("")

	loggerChain.LogMessage(DEBUG, "this is a debug level information.")

	fmt.Println("")

	loggerChain.LogMessage(ERROR, "this is an error information")
}
