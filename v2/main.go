package main

import (
	"fmt"

	stupidpackage "github.com/iagosaito/stupid-golang-module/v2/stupid-package"
)

func main() {
	fmt.Println("Stupid Golang Module")

	fmt.Println(stupidpackage.SaySomethingStupid())
	stupidpackage.TellSomeoneStupid("Brian")
}
