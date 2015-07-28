package basis

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	// declare: [n]T
	var array [2]string

	// array[0] = 1
	// array  = [2]int{2,3}

	// const len int = 3
	// array := [2 * len]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)
}
