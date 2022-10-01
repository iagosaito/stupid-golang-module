// Stupid package that provide just stupid things
package stupidpackage

import "fmt"

// Provide the stupid sentence:
//
// - "Sometimes when I close my eyes, I can’t see."
//
// That's stupid, isn't it?
func SaySomethingStupid() string {
	return "Sometimes when I close my eyes, I can’t see."
}

// TellSomeoneStupid takes in one parameter 'arg', the name of the person that you wanna call stupid.
// i.e ”You're completely stupid, Brian" if 'Brian was provided'
func TellSomeoneStupid(arg string) {
	fmt.Printf("You're completely stupid, %s\n", arg)
}
