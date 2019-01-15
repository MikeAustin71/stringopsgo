package strops

import (
	"fmt"
	"io"
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

func TestStrOps_GetValidBytes_01(t *testing.T) {

	validBytes := []byte{'v', 'a', 'l', 'i', 'd'}

	testBytes := []byte{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	actualBytes, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_02(t *testing.T) {

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualBytes, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_03(t *testing.T) {

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualBytes, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_04(t *testing.T) {

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualBytes, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_05(t *testing.T) {

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualBytes, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_06(t *testing.T) {

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := make([]byte, 0, 5)

	_, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidBytes_07(t *testing.T) {

	validBytes := make([]byte, 0, 5)

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrOps{}.GetValidBytes(testBytes, validBytes)

	if err == nil {
		t.Error("Expected Error return due to empty 'validBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidRunes_01(t *testing.T) {

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testRunes := []rune{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	actualRunes, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_02(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualRunes, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_03(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualRunes, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_04(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualRunes, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_05(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualRunes, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_06(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := make([]rune, 0, 5)

	_, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidRunes_07(t *testing.T) {

	validRunes := make([]rune, 0, 5)

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrOps{}.GetValidRunes(testRunes, validRunes)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidString_01(t *testing.T) {

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testStr := "xjvmRaJlZiFdS"

	expected := "valid"

	actualStr, err := StrOps{}.GetValidString(testStr, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_GetValidString_02(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "x13m5aJ7ZiFd5"

	expected := "1355"

	actualStr, err := StrOps{}.GetValidString(testStr, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_03(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xz3m5aJ7ZiFd5"

	expected := "355"

	actualStr, err := StrOps{}.GetValidString(testStr, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_04(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzJm!aJ%ZiFd^"

	expected := ""

	actualStr, err := StrOps{}.GetValidString(testStr, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_GetValidString_05(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzUmMaJ9ZiFd&"

	expected := ""

	actualStr, err := StrOps{}.GetValidString(testStr, validRunes)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_06(t *testing.T) {

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := ""

	_, err := StrOps{}.GetValidString(testStr, validRunes)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testStr'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidString_07(t *testing.T) {

	validRunes := make([]rune, 0, 5)

	testStr := "xzUmMaJ9ZiFd&"

	_, err := StrOps{}.GetValidString(testStr, validRunes)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
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

func TestStrOps_Read_01(t *testing.T) {

	expected := "original base string"
	lenExpected := len(expected)

	p := make([]byte, 100)

	s1 := StrOps{}.NewPtr()
	s1.SetStringData(expected)

	n, err := s1.Read(p)

	if err != nil && err != io.EOF {
		t.Errorf("Error returned by s1.Read(p). Error='%v' ", err.Error())
	}

	actualStr := string(p[:n])

	if expected != actualStr {
		t.Errorf("Error: Expected StrOut='%v'. Instead, StrOut='%v' ",
			expected, actualStr)
	}

	if lenExpected != n {
		t.Errorf("Error: Expected bytes read n='%v'. Instead, n='%v' ",
			lenExpected, n)
	}
}

func TestStrOps_Read_02(t *testing.T) {

	expected := "Original sops1 base string"
	lenExpected := len(expected)

	p := make([]byte, 5, 15)

	s1 := StrOps{}.NewPtr()
	s1.SetStringData(expected)
	n := 0
	var err error
	err = nil
	b := strings.Builder{}
	b.Grow(len(expected) + 150)

	for err != io.EOF {

		n, err = s1.Read(p)

		if err != nil && err != io.EOF {
			fmt.Printf("Error returned by s1.Read(p). "+
				"Error='%v' \n", err.Error())
			return
		}

		b.Write(p[:n])

		for i := 0; i < len(p); i++ {
			p[i] = byte(0)
		}

	}

	actualStr := b.String()

	if expected != actualStr {
		t.Errorf("Error: Expected StrOut='%v'. Instead, StrOut='%v' ",
			expected, actualStr)
	}

	lenActual := len(actualStr)

	if lenExpected != lenActual {
		t.Errorf("Error: Expected bytes read ='%v'. Instead, bytes read='%v' ",
			lenExpected, lenActual)
	}
}

func TestStrOps_Read_03(t *testing.T) {

	expected := "Original sops1 base string"
	lenExpected := int64(len(expected))

	s1 := StrOps{}.NewPtr()
	s1.SetStringData(expected)

	s2 := StrOps{}.NewPtr()

	n, err := io.Copy(s2, s1)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(sops2, sops1). "+
			"Error='%v' \n", err.Error())
		return
	}

	actualData := s2.GetStringData()

	if expected != actualData {
		t.Errorf("Error: Expected StrOut='%v'. Instead, String Data='%v' ",
			expected, actualData)
	}

	if lenExpected != n {
		t.Errorf("Error: Expected bytes read ='%v'. Instead, bytes read='%v' ",
			lenExpected, n)
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
