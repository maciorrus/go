package main

import (
	"fmt"

	"github.com/maciorrus/go/person"
)

func main() {
	p := new(person.Person)
	p.SetName("Maciek")
	fmt.Println(p.GetName())
}
