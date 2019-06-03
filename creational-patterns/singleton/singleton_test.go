package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestAbstractFactory(t *testing.T ) {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			single1 := GetInstance1(i)
			single1.ShowMessage()
		}(i)
	}
	wg.Wait()

	fmt.Println("**************************")

	var wg1 sync.WaitGroup
	for i := 0; i < 10; i++{
		go func(i int) {
			wg1.Add(1)
			defer wg1.Done()
			single1 := GetInstance2(i)
			single1.ShowMessage()
		}(i)
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++{
		go func(i int) {
			wg2.Add(1)
			defer wg2.Done()
			single1 := GetInstance4(i)
			single1.ShowMessage()
		}(i)
	}
	wg2.Wait()
}