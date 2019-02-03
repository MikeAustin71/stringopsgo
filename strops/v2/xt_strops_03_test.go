package strops

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

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

func TestStrOps_Read_04(t *testing.T) {
	originalStr := "Hello World"

	sops1 := StrOps{}.NewPtr()
	sops1.SetStringData(originalStr)
	p := make([]byte, 0)

	_, err := sops1.Read(p)

	if err == nil {
		t.Error("Error: Expected error return. NO ERROR WAS RETURNED!")
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
