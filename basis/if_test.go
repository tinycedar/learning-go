package basis

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if i := 1; i > 0 {
		fmt.Println("Bigger than zero")
	}
}
