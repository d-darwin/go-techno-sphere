package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	fmt.Println("Initialization...")
}

func main() {
	myTimer := getTimer()
	defer catchPanic()

	for i := 0; i < 10; i++ {
		// defer stack
		defer fmt.Println(i)
	}

	showMeTheMoney()
	fmt.Println(sum([]int{1, 2, 3}...))
	myTimer()

	(func() {
		// IIFE
		myTimer := "Just a string"
		fmt.Println(myTimer)
	})()

	ReadFile("some file")

	time.Sleep((1 * time.Second))

	makePanic()
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
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}

func ReadFile(f string) (*os.File, error) {
	file, err := os.OpenFile(f, 0, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

func makePanic() {
	panic("It's ok, it's ok")
}

func catchPanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic deferred")
	}
}
