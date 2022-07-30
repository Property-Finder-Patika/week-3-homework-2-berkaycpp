package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 1023
	max := 9876
	RandomIntegerwithinRange := rand.Intn(max-min) + min
	first := int(RandomIntegerwithinRange / 1000)
	second := int(RandomIntegerwithinRange/100) % 10
	third := int(RandomIntegerwithinRange/10) % 100
	forth := RandomIntegerwithinRange % 1000

	for {
		if first == second {
			second = (second + 1) % 10
			continue
		}
		if first == third {
			third = (third + 1) % 10
			continue
		}
		if first == forth {
			forth = (forth + 1) % 10
			continue
		}
		if second == third {
			third = (third + 1) % 10
			continue
		}
		if second == forth {
			forth = (forth + 1) % 10
			continue
		}
		if third == forth {
			forth = (forth + 1) % 10
			continue
		}
		break
	}

	var guess int
	var guessList []int

	for {
		fmt.Println("Enter guess:")
		fmt.Println("You've tried:", guessList)
		fmt.Scan(&guess)
		guessList = append(guessList, guess)

		firstGuess := int(guess / 1000)
		secondGuess := int(guess/100) % 10
		thirdGuess := int(guess/10) % 100
		forthGuess := guess % 1000

		fullyCorrect, partiallyCorrect := 0, 0

		if firstGuess == first {
			fullyCorrect = fullyCorrect + 1
		}
		if firstGuess == second {
			partiallyCorrect = partiallyCorrect + 1
		}
		if firstGuess == third {
			partiallyCorrect = partiallyCorrect + 1
		}
		if firstGuess == forth {
			partiallyCorrect = partiallyCorrect + 1
		}

		if secondGuess == first {
			partiallyCorrect = partiallyCorrect + 1
		}
		if secondGuess == second {
			fullyCorrect = partiallyCorrect + 1
		}
		if secondGuess == third {
			partiallyCorrect = partiallyCorrect + 1
		}
		if secondGuess == forth {
			partiallyCorrect = partiallyCorrect + 1
		}

		if thirdGuess == first {
			partiallyCorrect = partiallyCorrect + 1
		}
		if thirdGuess == second {
			partiallyCorrect = partiallyCorrect + 1
		}
		if thirdGuess == third {
			fullyCorrect = fullyCorrect + 1
		}
		if thirdGuess == forth {
			partiallyCorrect = partiallyCorrect + 1
		}

		if forthGuess == first {
			partiallyCorrect = partiallyCorrect + 1
		}
		if forthGuess == second {
			partiallyCorrect = partiallyCorrect + 1
		}
		if forthGuess == third {
			partiallyCorrect = partiallyCorrect + 1
		}
		if forthGuess == forth {
			fullyCorrect = fullyCorrect + 1
		}
		if fullyCorrect == 4 {
			fmt.Println("You did it!")
			break
		} else {
			fmt.Println("+", fullyCorrect, "-", partiallyCorrect)
		}
	}
	return
}
