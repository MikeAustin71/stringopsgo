package main

import (
	strOps "./strops/v2"
	"fmt"
)

func main() {

	fmt.Println("main() - Version 2")

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'x'
	replaceRunes[0][1] = 'X'

	replaceRunes[1][0] = 'y'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'z'
	replaceRunes[2][1] = 'Z'

	replaceRunes[3][0] = 'a'
	replaceRunes[3][1] = 'A'

	replaceRunes[4][0] = 'b'
	replaceRunes[4][1] = 0

	targetStr := "1x2y3z4a5b"
	expected := "1X23Z4A5"

	su := strOps.StrOps{}

	outputRunes, err := su.ReplaceRunes([]rune(targetStr), replaceRunes)

	if err != nil {
		fmt.Printf("Error returned from su.ReplaceRunes( []rune(targetStr), replaceRunes). "+
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("   targetStr: ", targetStr)
	fmt.Println("expected Str: ", expected)
	fmt.Println("  actual Str: ", string(outputRunes))

}
