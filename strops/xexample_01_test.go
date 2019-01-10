package strops

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// SampleExpressions_Example_01 -
func SampleExpressionsExample01() {
	//AMpm Match \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
	//PM Match V1 "\\d{1}\\s?(?i)p[.\\s]*(?i)m[.]*"
	//PM Match V2 "\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*"
	//AM Match V1 "\\d{1}\\s?(?i)a[.\\s]*(?i)m[.]*"
	//AM Match V2 "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"
	regexAMpm := "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"

	samples := []string{
		"12:15 AM",
		"03:25AM",
		"11:19 A M",
		"11:19 A. M.",
		"11:19 A.M.",
		"11:19A.M.",
		"11:19  A.M.",
		"11:19  AM",
		"11:19  AM",
		"11:19  A  M",
		"12:15 am",
		"03:25am",
		"11:19 a m",
		"11:19a m",
		"11:19 a. m.",
		"11:19 a.m.",
		"11:19a.m.",
		"11:19  A  M",
		"11:19  A. M.",
		"11:19  a  m",
		"11:19  a. m.",
		"11:19 m",
		"11:19 a",
		"10:25 PM",
		"02:15PM",
		"10:18 P M",
		"01:19 P. M.",
		"12:19 P.M.",
		"10:19P.M.",
		"10:15 pm",
		"04:25pm",
		"10:19 p m",
		"10:19p m",
		"10:19 p. m.",
		"10:19p.m.",
		"15:35:03",
		"10:19:16 p.m.",
		"10:15 pm -0600 MST",
		"10:15 pm-0600 MST",
		"10:15 pm PST",
		"10:15  pm -0600 MST",
		"10:15 p.m -0600 MST",
		"10:15 pm. -0600 MST",
		"10:15 m -0600 MST",
		"10:15 p -0600 MST",
		"11:19  P.M.",
		"11:19  PM",
		"11:19  PM",
		"11:19  P  M",
		"11:19  P. M.",
		"11:19  p  m",
		"11:19  p. m.",
	}

	lArray := len(samples)
	for i := 0; i < lArray; i++ {
		match, err := FindExpressionExample01(samples[i], regexAMpm)

		if err != nil {
			if err.Error() == "No Match" {
				fmt.Printf("No Match - testStr == %v  regex == %v\n", samples[i], regexAMpm)
				continue
			} else {
				panic(err)
			}
		}

		fmt.Printf("Match! - testStr == %v  regex == %v  match string: %v \n", samples[i], regexAMpm, match)

	}

}

// FindExpression_Example_01 - Example function.
func FindExpressionExample01(targetStr string, regex string) (string, error) {

	if len(targetStr) < 1 {
		return "", fmt.Errorf("ExampleFindExpression_01() Invalid Target String: %v", targetStr)
	}

	// \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
	r, err := regexp.Compile(regex)

	if err != nil {
		return "", fmt.Errorf("Regex failed to Compile. regex== %v. Error: %v", regex, err.Error())
	}

	bTargetStr := []byte(targetStr)

	loc := r.FindIndex(bTargetStr)

	if loc == nil {
		return "", errors.New("No Match")
	}

	return string(bTargetStr[loc[0]:loc[1]]), nil

}

func TrimMultipleStringsExample01(tStr string, trimChar rune) {

	su := StrOps{}

	r, err := su.TrimEndMultiple(tStr, trimChar)

	if err != nil {
		fmt.Println("Error Return from TrimMultipleChars: ", err.Error())
		return
	}

	fmt.Println("Original String: ", tStr)
	fmt.Println(" Trimmed String: ", r)
	fmt.Println("Original String Length: ", len(tStr))
	fmt.Println(" Trimmed String Length: ", len(r))
	tStr2 := strings.Replace(tStr, " ", "!", -1)
	fmt.Println("Original String TrimChar Locations: ", tStr2)
	r2 := strings.Replace(r, " ", "!", -1)
	fmt.Println(" Trimmed String TrimChar Locations: ", r2)

}

// RegExFindSingleTimeDigits_Example01
func RegExFindSingleTimeDigitsExample01() {
	regex := "\\d:\\d:\\d"
	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"

	fmt.Println("targetStr = ", targetStr)
	su := StrOps{}

	idx := su.FindRegExIndex(targetStr, regex)

	if idx == nil {
		panic(fmt.Errorf("Did not locate Regular Expression,'%v', in 'targetStr', '%v'.", regex, targetStr))
	}

	fmt.Println("Success - Found Regular Expression in targetStr!")
	fmt.Println("idx = ", idx)

	extract := make([]byte, 0)
	s := []byte(targetStr)

	extract = s[idx[0]:idx[1]]

	sExtract := string(extract)

	fmt.Println("Extracted String: ", sExtract)

	result := strings.Split(sExtract, ":")

	if len(result) == 0 {
		panic(fmt.Errorf("Split returned array of zero length"))
	}

	fmt.Println("Printing result array:")
	for j := 0; j < len(result); j++ {
		fmt.Println(result[j])
	}

	hrs, _ := strconv.Atoi(result[0])
	min, _ := strconv.Atoi(result[1])
	sec, _ := strconv.Atoi(result[2])

	fmt.Println("Printing Formatted Time String")
	fmt.Printf("%02d:%02d:%02d\n", hrs, min, sec)

	fmt.Println("Reprint with 2-digit seconds")
	fmt.Printf("%02d:%02d:%02d\n", hrs, min, 14)

}

// PrintFmtExample01
func PrintFmtExample01() {

	s1 := fmt.Sprintf("No1: %d  No2: %d", 1, 2)

	fmt.Println(s1)
}
