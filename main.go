package main

import (
	"fmt"

	"github.com/vn-ki/go-exp/stringutil"
)

func main() {
	fmt.Println("go go gophers...")

	// stringutil.ReverseNew, stringutil.Reverse
	//
	// fmt.Println(stringutil.ReverseNew("abcde"))
	// fmt.Println(stringutil.Reverse("abcde"))

	// stringutil.Capitalize
	// a := stringutil.Capitalize("sdfasdSSD2948::;lv")
	// fmt.Println(string(a))

	fmt.Println(stringutil.CapRev("asd"))

}
