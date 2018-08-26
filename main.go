package main

import (
	"fmt"
)

type Named interface {
	name() string
}
type Person struct {
	nm string
}

func (p Person) name() string {
	return p.nm
}

func main() {
	p := Person{nm: "Him"}
	named := Named(p)
	//named = p
	fmt.Println(named.name())
}
