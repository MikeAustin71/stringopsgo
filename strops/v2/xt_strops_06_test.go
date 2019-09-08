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

	charRune := '*'

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

	charRune := '='

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

func TestStrOps_ReadStringFromBytes_01(t *testing.T) {

	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "Hello World"
	expectedNextIdx := 13
	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_02(t *testing.T) {

	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "Does your program run?"
	expecteNextIdx := -1

	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_03(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24  25
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		//26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47   48
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "How are you?"
	expecteNextIdx := 26
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_04(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r',
		//
		//12  13  14  15  16  17  18  19  20  21  22  23  24
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		//25  26  27  28  29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46   47
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "How are you?"
	expecteNextIdx := 25
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'.\nInstead, result='%v'\n",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'.\nInstead, nextStartIdx='%v'\n",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_05(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', ',', ' ',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24
		'h', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := "Hello World, how are you?"
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_06(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', ',', ' ',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24
		'h', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := ""
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_07(t *testing.T) {

	var bytes[]byte

	expectedStr := ""
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_08(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24   25   26
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\r', '\n',
		//27  28  29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46   47 48   49   50
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r', '\n'}

	expectedStr := "Does your program run?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'.\n Instead, result='%v'\n",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'.\nInstead, nextStartIdx='%v'\n",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_09(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12   13
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v', '\n',
		//
		//14  15  16  17  18  19  20  21  22  23  24   25   26  27   28
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\r', '\v', '\n',
		//29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46  47  48  49  50   51   52
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r', '\n'}

	expectedStr := "Does your program run?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_10(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12   13
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v', '\n',
		//
		//14  15  16  17  18  19  20  21  22  23  24   25
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := "How are you?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_11(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}

	expectedStr := "Hello World!"
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_12(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\n'}

	expectedStr := "Hello World!"
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_13(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\r'}

	expectedStr := "Hello World!"
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
	}

}

func TestStrOps_ReadStringFromBytes_14(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\v'}

	expectedStr := "Hello World!"
	expecteNextIdx := -1

	result, nextStartIdx := StrOps{}.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expecteNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expecteNextIdx, nextStartIdx)
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
