package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Initialization...")
}

func main() {
	myTimer := getTimer()
	defer myTimer()
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

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Println("Time from start %v", time.Since(start))
	}
}
