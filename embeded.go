package main

// like inheritance

type Person struct {
	Name string
	inn  string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person
	Stuff
	LicenseToKill bool
}
