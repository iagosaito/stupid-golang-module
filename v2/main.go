package main

import (
	"fmt"

	stupidpackage "github.com/iagosaito/stupid-golang-module/v2/stupid-package"
)

func main() {
	fmt.Println("Stupid Golang Module v2")

	fmt.Println(stupidpackage.GetSomethingStupid(11))
}
