package main

import "fmt"

func main() {
	killer := SecretAgent{
		Person: Person{Name: "Sebastian Shopeen"},
	}

	fmt.Println("secret inn", killer.GetName())
}

// like inheritance

type Person struct {
	Name string
	inn  string
}

func (p Person) GetName() string {
	return p.Name
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person
	Stuff
	inn           float32
	LicenseToKill bool
}
