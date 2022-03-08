package main

import "fmt"

func main() {
	showMeTheMoney()
	fmt.Println(sum([]int{1, 2, 3}...))
}

func showMeTheMoney() {
	fmt.Println("$$$")
}

func sum(argList ...int) (res int) {
	for _, arg := range argList {
		res += arg
	}
	return
}
