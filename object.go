package main

import "fmt"

func main() {

}

type MyStruct struct {
	Num  int
	Name string
}

type myStructSlice []MyStruct

func (m myStructSlice) Less(i int) bool
func (m myStructSlice) Len() int
func (m myStructSlice) Swap(i, j int)

type MyInt int

func (m MyInt) showYourself() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m += i
}
