package strops

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Package strops - located in source code repository:
//
//			https://github.com/MikeAustin71/stringopsgo.git
//

// StrOps - encapsulates a collection of
// methods used to manipulate strings
type StrOps struct {
	StrIn  string
	StrOut string
}

//FindRegExIndex - returns a two-element slice of integers defining the location
// of the leftmost match in targetStr of the regular expression (regex).
//
// Return Value
// ============
// The return value is an array of integers. If no match is found the return
// value is 'nil'.  If regular expression is successfully matched, the match
// will be located at targetStr[loc[0]:loc[1]]. Again, a return value of 'nil'
// signals that no match was found.
//
func (sops StrOps) FindRegExIndex(targetStr string, regex string) []int {

	re := regexp.MustCompile(regex)

	return re.FindStringIndex(targetStr)

}

// GetSoftwareVersion - Returns the software version for StrOps.
func (sops StrOps) GetSoftwareVersion() string {
	return "1.0.0"
}

// ReplaceMultipleStrs - Replaces all instances of string replaceMap[i][0][0] with
// replacement string replaceMap[i][0][1] in 'targetStr'
func (sops StrOps) ReplaceMultipleStrs(targetStr string, replaceMap [][][]string) string {

	max := len(replaceMap)

	for i := 0; i < max; i++ {
		if strings.Contains(targetStr, replaceMap[i][0][0]) {
			targetStr = strings.Replace(targetStr, replaceMap[i][0][0], replaceMap[i][0][1], 1)
		}

	}

	return targetStr
}

// StrCenterInStrLeft - returns a string which includes
// a left pad blank string plus the original string. It
// does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively
// center the original string is a field of specified length.
func (sops StrOps) StrCenterInStrLeft(strToCenter string, fieldLen int) (string, error) {

	pad, err := sops.StrPadLeftToCenter(strToCenter, fieldLen)

	if err != nil {
		return "", errors.New("StrOps:StrCenterInStrLeft() - " + err.Error())
	}

	return pad + strToCenter, nil

}

// StrCenterInStr - returns a string which includes
// a left pad blank string plus the original string,
// plus a right pad blank string.
//
// The complete string will effectively center the
// original string is a field of specified length.
func (sops StrOps) StrCenterInStr(strToCenter string, fieldLen int) (string, error) {

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter, fmt.Errorf("'fieldLen' = '%v' strToCenter Length= '%v'. "+
			"'fieldLen' is shorter than strToCenter Length!", fieldLen, sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen - sLen) / 2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}

	return leftPadStr + strToCenter + rightPadStr, nil

}

// StrLeftJustify - Creates a new string with a total length equal to input parameter
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the output string and
// spaces are padded to the right in order to create a string with total length of
// 'fieldLen'.
//
// Example:
//
//  fieldLen        = 15
//  strToJustify    = "Hello World"
//  Returned String = "Hello World    "
//  String Index    =  012345648901234
//
func (sops StrOps) StrLeftJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify, fmt.Errorf("StrLeftJustify() Error: Length of string to left justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	rightPadLen := fieldLen - strLen

	rightPadStr := strings.Repeat(" ", rightPadLen)

	return strToJustify + rightPadStr, nil

}

// StrRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
func (sops StrOps) StrRightJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify, fmt.Errorf("StrRightJustify() Error: Length of string to right justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)

	return leftPadStr + strToJustify, nil
}

// StrPadLeftToCenter - Returns a blank string
// which allows centering of the target string
// in a fixed length field.
func (sops StrOps) StrPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

	sLen := sops.StrGetRuneCnt(strToCenter)

	if sLen > fieldLen {
		return "", errors.New("StrOps:StrPadLeftToCenter() - String To Center is longer than Field Length")
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// StrGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (sops StrOps) StrGetRuneCnt(targetStr string) int {
	return utf8.RuneCountInString(targetStr)
}

// StrGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (sops StrOps) StrGetCharCnt(targetStr string) int {
	return len([]rune(targetStr))
}

// TrimMultipleChars - Performs the following operations on strings:
// 1. Trims Right and Left for all instances of 'trimChar'
// 2. Within the interior of a string, multiple instances
// 		of 'trimChar' are reduce to a single instance.
//
// Example:
//
// targetStr = "       Hello          World        "
// trimChar  = ' ' (One Space)
// returned string (rStr) = "Hello World"
//
func (sops StrOps) TrimEndMultiple(targetStr string, trimChar rune) (rStr string, err error) {

	if targetStr == "" {
		err = errors.New("Empty targetStr")
		return
	}

	fStr := []rune(targetStr)
	lenTargetStr := len(fStr)
	outputStr := make([]rune, lenTargetStr)
	lenTargetStr--
	idx := lenTargetStr
	foundFirstChar := false

	for i := lenTargetStr; i >= 0; i-- {

		if !foundFirstChar && fStr[i] == trimChar {
			continue
		}

		if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
			continue
		}

		if i == 0 && fStr[i] == trimChar {
			continue
		}

		foundFirstChar = true
		outputStr[idx] = fStr[i]
		idx--
	}

	if idx != lenTargetStr {
		idx++
	}

	if outputStr[idx] == trimChar {
		idx++
	}

	result := string(outputStr[idx:])

	return result, nil

}

// SwapRune - Swaps all instances of 'oldRune' character with 'newRune'
// character.
func (sops StrOps) SwapRune(currentStr string, oldRune rune, newRune rune) (string, error) {

	if currentStr == "" {
		return currentStr, nil
	}

	rStr := []rune(currentStr)

	lrStr := len(rStr)

	for i := 0; i < lrStr; i++ {
		if rStr[i] == oldRune {
			rStr[i] = newRune
		}
	}

	return string(rStr), nil
}
