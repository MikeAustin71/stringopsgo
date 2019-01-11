package strops

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrOps_ReplaceStringChars_01(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "1A2B3C4D5E6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 'A'

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 'B'

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 'C'

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 'D'

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 'E'

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceStringChars_02(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "1A23C45E6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 'A'

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 'C'

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 'E'

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceStringChars_03(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "1a2b3c4d5e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'z'
	replaceRunes[0][1] = 'Z'

	replaceRunes[1][0] = 'y'
	replaceRunes[1][1] = 'Y'

	replaceRunes[2][0] = 'x'
	replaceRunes[2][1] = 'X'

	replaceRunes[3][0] = 'w'
	replaceRunes[3][1] = 'W'

	replaceRunes[4][0] = 'v'
	replaceRunes[4][1] = 'V'

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceStringChars_04(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "3a4b5c6d7e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = '1'
	replaceRunes[0][1] = '3'

	replaceRunes[1][0] = '2'
	replaceRunes[1][1] = '4'

	replaceRunes[2][0] = '3'
	replaceRunes[2][1] = '5'

	replaceRunes[3][0] = '4'
	replaceRunes[3][1] = '6'

	replaceRunes[4][0] = '5'
	replaceRunes[4][1] = '7'

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceStringChars_05(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "1a23c4d5e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'z'
	replaceRunes[0][1] = 'Z'

	replaceRunes[1][0] = 'y'
	replaceRunes[1][1] = 'Y'

	replaceRunes[2][0] = 'x'
	replaceRunes[2][1] = 'X'

	replaceRunes[3][0] = 'w'
	replaceRunes[3][1] = 'W'

	replaceRunes[4][0] = 'b'
	replaceRunes[4][1] = 0

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceStringChars_06(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	expected := "123456"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 0

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 0

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 0

	actualStr, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceStringChars(testStr, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceStringChars_07(t *testing.T) {

	testStr := ""

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 0

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 0

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 0

	_, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR RETURNED!. ")
	}
}

func TestStrOps_ReplaceStringChars_08(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	replaceRunes := make([][]rune, 0, 0)

	_, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR WAS RETURNED! ")
	}

}

func TestStrOps_ReplaceStringChars_09(t *testing.T) {

	testStr := "1a2b3c4d5e6"

	replaceRunes := make([][]rune, 5, 10)

	_, err := StrOps{}.ReplaceStringChars(testStr, replaceRunes)

	if err == nil {
		t.Errorf("Error: Expected error return. NO ERROR WAS RETURNED!")
	}
}

func TestStrOps_ReplaceRunes_01(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "1A2B3C4D5E6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 'A'

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 'B'

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 'C'

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 'D'

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 'E'

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceRunes_02(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "1A23C45E6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 'A'

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 'C'

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 'E'

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceRunes_03(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "1a2b3c4d5e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'z'
	replaceRunes[0][1] = 'Z'

	replaceRunes[1][0] = 'y'
	replaceRunes[1][1] = 'Y'

	replaceRunes[2][0] = 'x'
	replaceRunes[2][1] = 'X'

	replaceRunes[3][0] = 'w'
	replaceRunes[3][1] = 'W'

	replaceRunes[4][0] = 'v'
	replaceRunes[4][1] = 'V'

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceRunes_04(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "3a4b5c6d7e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = '1'
	replaceRunes[0][1] = '3'

	replaceRunes[1][0] = '2'
	replaceRunes[1][1] = '4'

	replaceRunes[2][0] = '3'
	replaceRunes[2][1] = '5'

	replaceRunes[3][0] = '4'
	replaceRunes[3][1] = '6'

	replaceRunes[4][0] = '5'
	replaceRunes[4][1] = '7'

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceRunes_05(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "1a23c4d5e6"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'z'
	replaceRunes[0][1] = 'Z'

	replaceRunes[1][0] = 'y'
	replaceRunes[1][1] = 'Y'

	replaceRunes[2][0] = 'x'
	replaceRunes[2][1] = 'X'

	replaceRunes[3][0] = 'w'
	replaceRunes[3][1] = 'W'

	replaceRunes[4][0] = 'b'
	replaceRunes[4][1] = 0

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceRunes_06(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	expected := "123456"

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 0

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 0

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 0

	actualRunes, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceRunes(testRunes, replaceRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceRunes_07(t *testing.T) {

	testRunes := make([]rune, 0, 0)

	replaceRunes := make([][]rune, 5, 10)

	for i := 0; i < 5; i++ {
		replaceRunes[i] = make([]rune, 2, 5)
	}

	replaceRunes[0][0] = 'a'
	replaceRunes[0][1] = 0

	replaceRunes[1][0] = 'b'
	replaceRunes[1][1] = 0

	replaceRunes[2][0] = 'c'
	replaceRunes[2][1] = 0

	replaceRunes[3][0] = 'd'
	replaceRunes[3][1] = 0

	replaceRunes[4][0] = 'e'
	replaceRunes[4][1] = 0

	_, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR RETURNED!. ")
	}
}

func TestStrOps_ReplaceRunes_08(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	replaceRunes := make([][]rune, 0, 0)

	_, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR WAS RETURNED! ")
	}

}

func TestStrOps_ReplaceRunes_09(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testRunes := []rune(testStr)

	replaceRunes := make([][]rune, 5, 10)

	_, err := StrOps{}.ReplaceRunes(testRunes, replaceRunes)

	if err == nil {
		t.Errorf("Error: Expected error return. NO ERROR WAS RETURNED!")
	}
}

func TestStrOps_StrCenterInStr_001(t *testing.T) {
	strToCntr := "1234567"
	fieldLen := 79
	exLeftPadLen := 36
	exRightPadLen := 36
	exTotalLen := 79

	leftPad := strings.Repeat(" ", exLeftPadLen)
	rightPad := strings.Repeat(" ", exRightPadLen)
	exStr := leftPad + strToCntr + rightPad

	su := StrOps{}
	str, err := su.StrCenterInStr(strToCntr, fieldLen)
	if err != nil {
		t.Error("StrCenterInStr() generated error: ", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}

}

func TestStrOps_StrLeftJustify_001(t *testing.T) {
	strToJustify := "1234567"
	fieldLen := 45
	exTotalLen := fieldLen
	exRightPad := strings.Repeat(" ", 38)
	exStr := strToJustify + exRightPad

	su := StrOps{}
	str, err := su.StrLeftJustify(strToJustify, fieldLen)
	if err != nil {
		t.Error("StrLeftJustify() generated error: ", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}

}

func TestStrOps_StrRightJustify_001(t *testing.T) {

	strToJustify := "1234567"
	fieldLen := 45
	exTotalLen := fieldLen
	exLeftPad := strings.Repeat(" ", 38)
	exStr := exLeftPad + strToJustify

	su := StrOps{}
	str, err := su.StrRightJustify(strToJustify, fieldLen)
	if err != nil {
		t.Error("StrRightJustify() generated error: ", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}

}

func TestStrOps_StrCenterInStrLeft_001(t *testing.T) {
	strToCntr := "1234567"
	fieldLen := 79
	exPadLen := 36
	exTotalLen := 43

	exStr := strings.Repeat(" ", exPadLen) + strToCntr
	su := StrOps{}
	str, err := su.StrCenterInStrLeft(strToCntr, fieldLen)
	if err != nil {
		t.Error("StrCenterInStrLeft() generated error: ", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}

}

func TestStrOps_StrGetRuneCnt(t *testing.T) {
	strToCnt := "1234567"
	exCnt := 7
	su := StrOps{}
	l1 := su.StrGetRuneCnt(strToCnt)

	if l1 != exCnt {
		t.Error(fmt.Sprintf("Expected string character count of '%v', got", exCnt), l1)
	}

}

func TestStrOps_StrGetCharCnt01(t *testing.T) {
	strToCnt := "1234567"
	exCnt := 7

	su := StrOps{}
	l1 := su.StrGetCharCnt(strToCnt)

	if l1 != exCnt {
		t.Error(fmt.Sprintf("Expected string character count of '%v', got", exCnt), l1)
	}
}

func TestStrOps_StrPadLeftToCenter(t *testing.T) {
	strToCntr := "1234567"
	fieldLen := 79
	exLen := 36
	su := StrOps{}
	padStr, err := su.StrPadLeftToCenter(strToCntr, fieldLen)

	if err != nil {
		t.Error("Error on StrPadLeftToCenter(), got", err.Error())
	}

	l1 := su.StrGetRuneCnt(padStr)

	if l1 != exLen {
		t.Error(fmt.Sprintf("Expected pad length of '%v', got ", exLen), l1)
	}

}

func TestStrOps_SwapRune_001(t *testing.T) {
	su := StrOps{}

	tStr := "  Hello   World  "
	expected := "!!Hello!!!World!!"
	result, err := su.SwapRune(tStr, ' ', '!')

	if err != nil {
		t.Error("Error returned from SwapRune: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

	resultLen := len(result)
	expectedLen := len(expected)

	if resultLen != expectedLen {
		t.Errorf("Expected result length == '%v' instead received result length == '%v'", expectedLen, resultLen)
	}

}

func TestStrOps_SwapRune_002(t *testing.T) {
	su := StrOps{}

	tStr := "HelloWorld"
	expected := "HelloWorld"
	result, err := su.SwapRune(tStr, ' ', '!')

	if err != nil {
		t.Error("Error returned from SwapRune: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

	resultLen := len(result)
	expectedLen := len(expected)

	if resultLen != expectedLen {
		t.Errorf("Expected result length == '%v' instead received result length == '%v'", expectedLen, resultLen)
	}

}

func TestStrOps_SwapRune_003(t *testing.T) {
	su := StrOps{}

	tStr := "Hello Worldx"
	expected := "Hello WorldX"
	result, err := su.SwapRune(tStr, 'x', 'X')

	if err != nil {
		t.Error("Error returned from SwapRune: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

	resultLen := len(result)
	expectedLen := len(expected)

	if resultLen != expectedLen {
		t.Errorf("Expected result length == '%v' instead received result length == '%v'", expectedLen, resultLen)
	}

}

func TestStrOps_SwapRune_004(t *testing.T) {
	su := StrOps{}

	tStr := "xHello World"
	expected := "XHello World"
	result, err := su.SwapRune(tStr, 'x', 'X')

	if err != nil {
		t.Error("Error returned from SwapRune: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

	resultLen := len(result)
	expectedLen := len(expected)

	if resultLen != expectedLen {
		t.Errorf("Expected result length == '%v' instead received result length == '%v'", expectedLen, resultLen)
	}

}

func TestStrOps_TrimMultipleChars_001(t *testing.T) {
	tStr := " 16:26:32   CST "
	expected := "16:26:32 CST"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_002(t *testing.T) {
	tStr := "       Hello          World        "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_003(t *testing.T) {
	tStr := "Hello          World        "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_004(t *testing.T) {
	tStr := " Hello          World"
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_005(t *testing.T) {
	tStr := "Hello World"
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_006(t *testing.T) {
	tStr := "Hello World "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimMultipleChars_007(t *testing.T) {
	tStr := " Hello World "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimMultipleChars(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimStringEnds_01(t *testing.T) {

	tStr := "  Hello    World  "
	expected := "Hello    World"
	trimChar := ' '
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_02(t *testing.T) {

	tStr := "Hello X World"
	expected := "Hello X World"
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_03(t *testing.T) {

	tStr := "Hello WorlXd"
	expected := "Hello WorlXd"
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_04(t *testing.T) {

	tStr := "XXXHello WorlXdXXX"
	expected := "Hello WorlXd"
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_05(t *testing.T) {

	tStr := "XXXHello WorlXd"
	expected := "Hello WorlXd"
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_06(t *testing.T) {

	tStr := "Hello WorlXdXXXX"
	expected := "Hello WorlXd"
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_07(t *testing.T) {

	tStr := "X"
	expected := ""
	trimChar := 'X'
	result, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.TrimStringEnds(tStr, trimChar). "+
			"Error='%v' ", err.Error())
	}

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, result)
	}
}

func TestStrOps_TrimStringEnds_08(t *testing.T) {

	tStr := ""
	_, err := StrOps{}.TrimStringEnds(tStr, '!')

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}
}

func TestStrOps_TrimStringEnds_09(t *testing.T) {

	tStr := "Jay Ray"
	trimChar := rune(0)
	_, err := StrOps{}.TrimStringEnds(tStr, trimChar)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}
}
