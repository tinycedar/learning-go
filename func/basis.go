package main

import (
	"fmt"
)

// 1, 如果接收者值未在方法体内使用，可以省略方法接收者标识符
// 2, 可变长参数的使用
// 3, 多返回值的使用
// 4, 返回值的声明和使用
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

func main() {
	result, _ := hello("Say hello to ", "Daniel", "Elim")
	fmt.Println(result)
}
