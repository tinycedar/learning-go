package basis

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	begin := func(funcName string) string {
		fmt.Printf("Begin function: %s\n", funcName)
		return funcName
	}
	end := func(funcName string) string {
		fmt.Printf("End function: %s\n", funcName)
		return funcName
	}
	defer end(begin("test"))

	fmt.Println("In main function")
}

func TestLoop(t *testing.T) {
	defer func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Number: %d\n", i)
		}
	}()

	fmt.Println("In main function")
}

func TestAnotherLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("Number: %d\n", i)
		}()
	}

	fmt.Println("In main function")
}

func appendNumbers(ints []int) (result []int) {
	result = append(ints, 1)
	defer func() {
		result = append(result, 2)
	}()
	result = append(result, 3)
	defer func() {
		result = append(result, 4)
	}()
	result = append(result, 5)
	defer func() {
		result = append(result, 6)
	}()
	return
}

func TestDeferOrder(t *testing.T) {
	fmt.Println(appendNumbers([]int{0}))
}
