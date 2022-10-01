package main

import (
	"fmt"

	stupidpackage "github.com/iagosaito/stupid-golang-module/stupid-package"
)

func main() {
	fmt.Println("Stupid Golang Module")

	fmt.Println(stupidpackage.SaySomethingStupid())
	stupidpackage.TellSomeoneStupid("Brian")
}
