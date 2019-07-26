package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"woof woof"}, true}
	fifi := cat{animal{"meow meow"}, true}
	shadow := dog{animal{"woof"}, true}
	crittersArray := []interface{}{fido, fifi, shadow}
	fmt.Println(crittersArray)
}
