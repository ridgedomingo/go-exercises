package generator

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
)

var(
	length int
    isNumbersIncluded bool
    isSymbolsIncluded bool
    isUppercaseIncluded bool
	passwordType string
)

func Start() {

	flag.IntVar(&length, "length", 0, "password length")
	flag.BoolVar(&isNumbersIncluded, "includeNumbers", false, "should include numbers")
	flag.BoolVar(&isSymbolsIncluded, "includeSymbols", false, "should include symbols")
	flag.BoolVar(&isUppercaseIncluded, "includeUppercase", false, "should include uppercase")
	flag.StringVar(&passwordType, "type", "random", "password type")
	flag.Parse()

	password := generatePassword()
	fmt.Println("Generated password", password)

}

func generatePassword() string {
	var generatedPassword string
	switch passwordType {
    case "random":
		  // TO DO: Func to create random password
    case "alphanumeric":
		  // TO DO: Func to create alphanumeric password
    case "pin":
		  // TO DO: Func to create pin
		generatedPassword = generatePin()
    }

	return generatedPassword
}

func generatePin() string {
	// default pin to 6 digits
	pinLength := 6

	if length >= 1 {
		pinLength = length
	}
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(pow(10, pinLength)-pow(10, pinLength-1))))
	if err != nil {
		log.Fatal(err)
	}
	randomNum = randomNum.Add(randomNum, big.NewInt(int64(pow(10, pinLength-1))))


	return randomNum.String()
}

// Function to calculate power
func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	result := x
	for i := 1; i < y; i++ {
		result *= x
	}
	return result
}