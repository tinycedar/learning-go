package main

import (
	"fmt"
	"sort"
)

// 自定义类型，int slice的别名
type MyIntSlice []int

// begin of implements sort.Interface for slice sorting
func (self MyIntSlice) Len() int {
	return len(self)
}

func (self MyIntSlice) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self MyIntSlice) Less(i, j int) bool {
	return self[i] < self[j]
}

// 1, 如果接收者值未在方法体内使用，可以省略方法接收者标识符
// 2, 可变长参数的使用
// 3, 方法返回值的声明和使用
func (MyIntSlice) hello(strs ...string) (result string) {
	result = "Say hello to "
	for index, item := range strs {
		result += item
		if index < len(strs)-1 {
			result += ", "
		}
	}
	return result + " !"
}

// 值方法
func (self MyIntSlice) Sum() int {
	sum := 0
	for _, i := range self {
		sum += i
	}
	return sum
}

// 指针方法
func (self *MyIntSlice) Max() int {
	max := 0
	for _, i := range *self {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	myIntSlice := MyIntSlice{2, 4, 1, 3}
	fmt.Println(myIntSlice)

	fmt.Println(myIntSlice.hello("Daniel", "Elim"))
	fmt.Println(myIntSlice.Sum())
	fmt.Println(myIntSlice.Max())

	sort.Sort(myIntSlice)
	fmt.Println(myIntSlice)
}
