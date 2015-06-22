package basis

import (
	"errors"
	"fmt"
	"testing"
)

func genPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Error recovered: %s\n", e)
		}
	}()
	panic(errors.New("Error occuring..."))
	fmt.Println("In genPanic function")
}

func TestRecover(t *testing.T) {
	genPanic()
	fmt.Println("In main function")
}
