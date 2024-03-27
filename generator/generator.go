package generator

import (
	"flag"
)

var(
	length int
    isNumbersIncluded bool
    isSymbolsIncluded bool
    isUppercaseIncluded bool
	passwordType string
)

func Start() {

	flag.IntVar(&length, "length", 12, "password length")
	flag.BoolVar(&isNumbersIncluded, "includeNumbers", false, "should include numbers")
	flag.BoolVar(&isSymbolsIncluded, "includeSymbols", false, "should include symbols")
	flag.BoolVar(&isUppercaseIncluded, "includeUppercase", false, "should include uppercase")
	flag.StringVar(&passwordType, "type", "random", "password type")
	flag.Parse()

	generatePassword(length,isNumbersIncluded,isSymbolsIncluded,isUppercaseIncluded,passwordType)

}

func generatePassword(length int,isNumbersIncluded bool,isSymbolsIncluded bool,isUppercaseIncluded bool,passwordType string) {
	switch passwordType {
    case "random":
		  // TO DO: Func to create random password
    case "alphanumeric":
		  // TO DO: Func to create alphanumeric password
    case "pin":
		  // TO DO: Func to create pin
    }

}