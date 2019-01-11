package strops

import (
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

func TestStrOps_ReplaceBytes_01(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1A2B3C4D5E6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 'A'

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 'B'

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 'C'

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 'D'

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 'E'

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceBytes_02(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1A23C45E6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 'A'

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 'C'

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 'E'

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_ReplaceBytes_03(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1a2b3c4d5e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'z'
	replaceBytes[0][1] = 'Z'

	replaceBytes[1][0] = 'y'
	replaceBytes[1][1] = 'Y'

	replaceBytes[2][0] = 'x'
	replaceBytes[2][1] = 'X'

	replaceBytes[3][0] = 'w'
	replaceBytes[3][1] = 'W'

	replaceBytes[4][0] = 'v'
	replaceBytes[4][1] = 'V'

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceBytes_04(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "3a4b5c6d7e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = '1'
	replaceBytes[0][1] = '3'

	replaceBytes[1][0] = '2'
	replaceBytes[1][1] = '4'

	replaceBytes[2][0] = '3'
	replaceBytes[2][1] = '5'

	replaceBytes[3][0] = '4'
	replaceBytes[3][1] = '6'

	replaceBytes[4][0] = '5'
	replaceBytes[4][1] = '7'

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceBytes_05(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1a23c4d5e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'z'
	replaceBytes[0][1] = 'Z'

	replaceBytes[1][0] = 'y'
	replaceBytes[1][1] = 'Y'

	replaceBytes[2][0] = 'x'
	replaceBytes[2][1] = 'X'

	replaceBytes[3][0] = 'w'
	replaceBytes[3][1] = 'W'

	replaceBytes[4][0] = 'b'
	replaceBytes[4][1] = 0

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceBytes_06(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "123456"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 0

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 0

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 0

	actualRunes, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_ReplaceBytes_07(t *testing.T) {

	testBytes := make([]byte, 0, 0)

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 0

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 0

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 0

	_, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR RETURNED!. ")
	}
}

func TestStrOps_ReplaceBytes_08(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	replaceBytes := make([][]byte, 0, 0)

	_, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR WAS RETURNED! ")
	}

}

func TestStrOps_ReplaceBytes_09(t *testing.T) {

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	replaceBytes := make([][]byte, 5, 10)

	_, err := StrOps{}.ReplaceBytes(testBytes, replaceBytes)

	if err == nil {
		t.Errorf("Error: Expected error return. NO ERROR WAS RETURNED!")
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