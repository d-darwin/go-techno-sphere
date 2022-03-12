package main

import "fmt"

func main() {

}

type MyObject struct {
	Num  int
	Name string
}

type MyInt int

func (m MyInt) showYourself() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m += i
}