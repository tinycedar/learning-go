package function

import (
	"fmt"
	"testing"
)

// 1, 可变长参数的使用
// 2, 多返回值的使用
// 3, 返回值的声明和使用
func hello(prefix string, names ...string) (result string, err error) {
	result = prefix
	length := len(names)
	for index, name := range names {
		result += name
		if index < length-1 {
			result += ", "
		} else {
			result += " !"
		}
	}
	return
	// return result + " !", nil
}

func TestFunc(t *testing.T) {
	result, _ := hello("Say hello to ", "Daniel", "Elim")
	fmt.Println(result)
}
