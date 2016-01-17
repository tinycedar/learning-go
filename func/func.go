package main

import (
	"fmt"
	"sort"
)

// 值方法
type ArrayStruct struct {
	name [5]int
}

func (self ArrayStruct) Len() int {
	return len(self.name)
}

func (self ArrayStruct) Swap(i, j int) {
	self.name[i], self.name[j] = self.name[j], self.name[i]
}

func (self ArrayStruct) Less(i, j int) bool {
	return self.name[i] < self.name[j]
}

// 指针方法
type MyIntArray [5]int

func (self *MyIntArray) Len() int {
	return len(self)
}

func (self *MyIntArray) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self *MyIntArray) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self *MyIntArray) print() {
	fmt.Println(*self)
}

func main() {
	arrayStruct := ArrayStruct{[5]int{2, 4, 1, 3, 0}}
	fmt.Printf("Origin array: %o\n", arrayStruct.name)
	sort.Sort(arrayStruct)
	fmt.Printf("Sorted array: %o\n\n", arrayStruct.name)

	myIntArray := MyIntArray{2, 4, 1, 3, 0}
	fmt.Printf("Origin array: %o\n", myIntArray)
	sort.Sort(&myIntArray)
	fmt.Printf("Sorted array: %o\n", myIntArray)
	// 值为啥能调用指针方法？
	// Refer to《Go并发编程实战》P71
	myIntArray.print()
}
