package strops

import (
	"fmt"
	"strings"
	"testing"
)

/*

	'xt_strops_test.go' is located in source code repository:

			https://github.com/MikeAustin71/stringopsgo.git

	It contains tests related to type, 'StrOps' located
	in source file, 'strops.go'.

*/

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

func TestStrOps_TrimEndMultiple_001(t *testing.T) {
	tStr := " 16:26:32   CST "
	expected := "16:26:32 CST"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_002(t *testing.T) {
	tStr := "       Hello          World        "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_003(t *testing.T) {
	tStr := "Hello          World        "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == %v, instead received result== %v", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_004(t *testing.T) {
	tStr := " Hello          World"
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_005(t *testing.T) {
	tStr := "Hello World"
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_006(t *testing.T) {
	tStr := "Hello World "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
	}

}

func TestStrOps_TrimEndMultiple_007(t *testing.T) {
	tStr := " Hello World "
	expected := "Hello World"
	su := StrOps{}

	result, err := su.TrimEndMultiple(tStr, ' ')

	if err != nil {
		t.Error("Error Return from TrimMultipleChars: ", err.Error())
	}

	if result != expected {
		t.Errorf("Expected result == '%v' instead received result== '%v'", expected, result)
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
