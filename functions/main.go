package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Golang!")
	var printString string = "GoLang but a 2nd time"
	printMe(printString)

	var numerator int = 11
	var denominator int = 11
	var result, remainder, err = intDivision(numerator, denominator)

	// the two snippets between the comment slashes do the same thing functionality wise but obviously one is a if else and one is a switch statement
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result is %v", result)
	} else {
		fmt.Printf("The result is %v with a remainder of %v", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v but in a switch statement", result)
	default:
		fmt.Printf("The result is %v with a remainder of %v but within a switch statement", result, remainder)
	}
	///////////////
}

func printMe(input string) {
	fmt.Println("Hello", input)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division attempt by 0")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
