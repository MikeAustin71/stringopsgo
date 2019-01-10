package strops

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrOps_FindLastWord_13(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, 0, 71)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindLastWord_14(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, 6, 5)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindRegExIndex_01(t *testing.T) {

	regex := "\\d:\\d:\\d"
	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	expected := "1:6:3"
	su := StrOps{}

	idx := su.FindRegExIndex(targetStr, regex)

	if idx == nil {
		t.Errorf("Error: Did not locate Regular Expression,'%v', in 'targetStr', '%v'.",
			regex, targetStr)
	}

	sExtract := string(targetStr[idx[0]:idx[1]])

	if expected != sExtract {
		t.Errorf("Error: Expected regular expression match on string='%v'. "+
			"Instead, matched string='%v'. ", expected, sExtract)
	}
}

func TestStrOps_IsEmptyOrWhiteSpace_01(t *testing.T) {

	testStr := "       "

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_02(t *testing.T) {

	testStr := ""

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_03(t *testing.T) {

	testStr := " xyz "

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_04(t *testing.T) {

	testStr := "xyz"

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_05(t *testing.T) {

	testStr := "/t"

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_06(t *testing.T) {

	testStr := "/n           "

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_07(t *testing.T) {

	testStr := "  /n"

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_08(t *testing.T) {

	testStr := "  x"

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_09(t *testing.T) {

	testStr := "x   "

	result := StrOps{}.IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_MakeSingleCharString_01(t *testing.T) {

	sUtil := StrOps{}
	requestedLen := 20

	charRune := rune('*')

	outputStr, err := sUtil.MakeSingleCharString(charRune, requestedLen)

	if err != nil {
		t.Errorf("Error returned by sUtil.MakeSingleCharString(charRune, 10). "+
			"Error='%v' ", err.Error())
		return
	}

	outputStrLen := len(outputStr)

	if requestedLen != outputStrLen {
		t.Errorf("Error: Expected outputStr length='%v'. Instead, string length='%v'",
			requestedLen, outputStrLen)
	}

	for i := 0; i < outputStrLen; i++ {
		if rune(outputStr[i]) != charRune {
			t.Errorf("Error: outputStr rune at index='%v' DOES NOT MATCH "+
				"specified rune '%v'. Actual rune='%v' ", i, charRune, rune(outputStr[i]))
		}

	}

}

func TestStrOps_MakeSingleCharString_02(t *testing.T) {

	sUtil := StrOps{}
	requestedLen := 100

	charRune := rune('=')

	outputStr, err := sUtil.MakeSingleCharString(charRune, requestedLen)

	if err != nil {
		t.Errorf("Error returned by sUtil.MakeSingleCharString(charRune, 10). "+
			"Error='%v' ", err.Error())
		return
	}

	outputStrLen := len(outputStr)

	if requestedLen != outputStrLen {
		t.Errorf("Error: Expected outputStr length='%v'. Instead, string length='%v'",
			requestedLen, outputStrLen)
	}

	for i := 0; i < outputStrLen; i++ {
		if rune(outputStr[i]) != charRune {
			t.Errorf("Error: outputStr rune at index='%v' DOES NOT MATCH "+
				"specified rune '%v'. Actual rune='%v' ", i, charRune, rune(outputStr[i]))
		}

	}

}

func TestStrOps_ReplaceMultipleStrs_01(t *testing.T) {

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = "x"
	rStrs[1][0] = " "
	rStrs[1][1] = "J"
	rStrs[2][0] = "l"
	rStrs[2][1] = "F"

	expectedStr := "HeFFxJWxrFd"

	actualStr, err := StrOps{}.ReplaceMultipleStrs(testStr, rStrs)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrOps_ReplaceMultipleStrs_02(t *testing.T) {

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = ""
	rStrs[1][0] = " "
	rStrs[1][1] = ""
	rStrs[2][0] = "l"
	rStrs[2][1] = ""

	expectedStr := "HeWrd"

	actualStr, err := StrOps{}.ReplaceMultipleStrs(testStr, rStrs)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrOps_ReplaceMultipleStrs_03(t *testing.T) {

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "f"
	rStrs[0][1] = " "
	rStrs[1][0] = "j"
	rStrs[1][1] = "r"
	rStrs[2][0] = "M"
	rStrs[2][1] = "x"

	expectedStr := "Hello World"

	actualStr, err := StrOps{}.ReplaceMultipleStrs(testStr, rStrs)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrOps_ReplaceMultipleStrs_04(t *testing.T) {

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = "x"
	rStrs[1][0] = " "
	rStrs[1][1] = "J"
	rStrs[2][0] = "l"
	rStrs[2][1] = "F"

	expectedStr := "HeFFxJWxrFdJHeFFxJWxrFd"

	actualStr, err := StrOps{}.ReplaceMultipleStrs(testStr, rStrs)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrOps_ReplaceNewLines_01(t *testing.T) {

	testStr := "Hello\nWorld"
	replaceStr := " "
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_02(t *testing.T) {

	testStr := "Hello World"
	replaceStr := " "
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestStrOps_ReplaceNewLines_03(t *testing.T) {

	testStr := "\n\nHello\nWorld\n\n\n"
	replaceStr := ""
	expectedStr := "HelloWorld"
	lenExpectedStr := len(expectedStr)

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_04(t *testing.T) {

	testStr := "\n\nHello World"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_05(t *testing.T) {

	testStr := "Hello World\n\n"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_06(t *testing.T) {

	testStr := "Hello World\n"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_07(t *testing.T) {

	testStr := "\nHello World"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrOps_ReplaceNewLines_08(t *testing.T) {

	testStr := "\tHello World"
	replaceStr := ""
	expectedStr := "\tHello World"

	actualStr := StrOps{}.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
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
