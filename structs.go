package main

import "fmt"

type example struct {
	Flag    bool
	counter int16
	pi      float32
}

func main() {
	var e1 example
	fmt.Printf("%+v\n", e1)

	e2 := example{
		Flag:    true,
		counter: 37,
	}
	fmt.Printf("%+v\n", e2)

	e3 := example{false, 33, 1} // Should fill all the fields
	fmt.Printf("%+v\n", e3)
}
