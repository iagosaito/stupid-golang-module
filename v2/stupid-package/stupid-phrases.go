// Stupid package that provide just stupid things
package stupidpackage

import "errors"

var stupidPhrases = []string{
	"Sometimes when I close my eyes, I can’t see.",
	"I can’t brain today. I have the dumb.",
	"A wise man once told me to always listen carefully because…um…I can’t remember.",
	"It is not my fault that I never learned to accept responsibility!",
	"You have no reason to fear zombies, do you?",
	"Don’t believe everything you think.",
	"What if there were no hypothetical questions?",
	"To err is human. To arr is pirate.",
	"I’m not crazy. My imaginary friends can prove it.",
	"I can’t talk to you today. I talked to two people yesterday.",
}

// Provide a stupid sentence based on number value. For exemple, if you pass value '1' as a parameter, the method will retunr:
//
// - "Sometimes when I close my eyes, I can’t see."
//
// That's stupid, isn't it?
//
// Max argument number has to be between 1 or 10 (inclusive). If you pass not range number the program will return the follow error:
//
// - "number must be lower than 11 and greather than 0"
func GetSomethingStupid(number int) (string, error) {

	if number < 1 || number > 10 {
		return "", errors.New("error: number must be lower than 11 and greather than 0")
	}

	return stupidPhrases[number-1], nil
}
