package basis

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	switch number := 1000; {
	case number < 0:
		fmt.Println("Negative number")
	case 0 <= number && number < 60:
		fmt.Println("No pass")
	case 60 <= number && number < 80:
		fmt.Println("Passed")
	case 80 <= number && number <= 100:
		fmt.Println("Excellent")
	default:
		fmt.Println("Perfect !")
	}
}
