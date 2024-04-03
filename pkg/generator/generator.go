package generator

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
)

var(
	length uint
    isNumbersIncluded bool
    isSymbolsIncluded bool
    isUppercaseIncluded bool
	passwordType string
)

type PasswordParams struct {
	Length uint
	PasswordType string
	IsNumbersIncluded bool
	IsSymbolsIncluded bool
	IsUppercaseIncluded bool
}

func Start() {
	flag.UintVar(&length, "length", 0, "password length, no negative values")
	flag.BoolVar(&isNumbersIncluded, "includeNumbers", false, "set to true if password should include numbers")
	flag.BoolVar(&isSymbolsIncluded, "includeSymbols", false, "set to true if password should include symbols")
	flag.BoolVar(&isUppercaseIncluded, "includeUppercase", false, "set to true if password should include uppercase")
	flag.StringVar(&passwordType, "type", "random", "password type, valid values are random,alphanumeric,pin")
	flag.Parse()

	passwordParams := PasswordParams{
		Length: length,
		PasswordType: passwordType,
		IsNumbersIncluded: isNumbersIncluded,
		IsSymbolsIncluded: isSymbolsIncluded,
		IsUppercaseIncluded: isUppercaseIncluded,
	}
	password := GeneratePassword(passwordParams)
	fmt.Println("Generated password: ", password)

}

func GeneratePassword(options ...PasswordParams) string {
	var generatedPassword string

	if passwordType == "pin" {
		generatedPassword = generatePin()
	} else {
		generatedPassword = generateSecurePassword()
	}

	return generatedPassword
}

func generateSecurePassword() string {
	passwordLength:= uint(12)

	if length >= 1 {
		passwordLength = length
	}
	charset := "abcdefghijklmnopqrstuvwxyz"
	
	if passwordType == "alphanumeric" {
			charset += "0123456789"
	} else {
		if isSymbolsIncluded {
			charset += "!@#$%^&*()-_=+"
		}
		if isNumbersIncluded{
			charset += "0123456789"
		}
	}

	if isUppercaseIncluded {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	charsetLen := len(charset)

	randomBytes := make([]byte, passwordLength)

	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}

	password := make([]byte, passwordLength)
	for i := uint(0); i < passwordLength; i++ {
		// Convert random byte to index in the character set
		randomIndex := int(randomBytes[i]) % charsetLen
		password[i] = charset[randomIndex]
	}

	return string(password) 
}

func generatePin() string {
	// default pin to 6 digits
	pinLength := uint(6)

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
func pow(x, y uint) uint {
	if y == 0 {
		return 1
	}
	result := x
	for i := uint(1); i < y; i++ {
		result *= x
	}
	return result
}