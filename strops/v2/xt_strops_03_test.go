package strops

import (
	"fmt"
	"strings"
	"testing"
)

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
