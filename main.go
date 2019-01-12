package main

import (
	strOps "./strops/v2"
	"fmt"
)

func main() {

	fmt.Println("main() - Version 2")

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testRunes := []rune{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	actualRunes, err := strOps.StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		fmt.Printf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		fmt.Printf("Error: Expected result='%v'. Instead, result='%v'. \n",
			expected, actualStr)
		fmt.Println()
	}

	fmt.Println("  target Runes: ", string(testRunes))
	fmt.Println("  expected Str: ", expected)
	fmt.Println("  actual Runes: ", actualStr)

}
