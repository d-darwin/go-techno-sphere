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
}
