package strops

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"
)

func TestSortStrLengthHighestToLowest_Len_01(t *testing.T) {
	badChars := []string{
		"aaaaa",
		"bbbbb",
		"cccccccccc",
		"ddddddddd",
		"eeeeeeeeeee",
		"fffffffffff"}

	sort.Sort(SortStrLengthHighestToLowest(badChars))

	goodChars := []string{
		"fffffffffff",
		"eeeeeeeeeee",
		"cccccccccc",
		"ddddddddd",
		"bbbbb",
		"aaaaa"}

	for i := 0; i < len(badChars); i++ {
		if goodChars[i] != badChars[i] {
			errStr := "badChars mismatch!\nbadCharsArray=\n"
			for j := 0; j < len(badChars); j++ {
				errStr += fmt.Sprintf("%v\n", badChars[j])
			}

			t.Errorf("%v", errStr)
		}
	}

}

func TestSortStrLengthLowestToHighest01(t *testing.T) {
	badChars := []string{
		"aaaaa",
		"bbbbb",
		"cccccccccc",
		"ddddddddd",
		"eeeeeeeeeee",
		"fffffffffff",
		"x",
		"z"}

	sort.Sort(SortStrLengthLowestToHighest(badChars))

	goodChars := []string{
		"x",
		"z",
		"aaaaa",
		"bbbbb",
		"ddddddddd",
		"cccccccccc",
		"eeeeeeeeeee",
		"fffffffffff"}

	for i := 0; i < len(badChars); i++ {
		if goodChars[i] != badChars[i] {
			errStr := "badChars mismatch!\nbadCharsArray=\n"
			for j := 0; j < len(badChars); j++ {
				errStr += fmt.Sprintf("%v\n", badChars[j])
			}

			t.Errorf("%v", errStr)
		}
	}

}

func TestStrOps_StripBadChars_001(t *testing.T) {
	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}
	expectedStr := "SomeString"
	expectedStrLen := len(expectedStr)
	testString := "..........      ./../.\\.\\..\\////   " + expectedStr +
		"..........      ./../.\\.\\..\\////   "

	actualString, actualStrLen := StrOps{}.StripBadChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
		return
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
		return
	}
}

func TestStrOps_StripBadChars_002(t *testing.T) {

	badChars := make([]string, 0)

	expectedStr := "SomeString"

	testString := "..........      ./../.\\.\\..\\////   " + expectedStr +
		"..........      ./../.\\.\\..\\////   "

	expectedStr = testString
	expectedStrLen := len(expectedStr)

	actualString, actualStrLen := StrOps{}.StripBadChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
		return
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
		return
	}
}

func TestStrOps_StripBadChars_003(t *testing.T) {
	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}

	expectedStr := ""
	expectedStrLen := len(expectedStr)
	testString := expectedStr

	actualString, actualStrLen := StrOps{}.StripBadChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
		return
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
		return
	}
}

func TestStrOps_StripBadChars_004(t *testing.T) {
	badChars := []string{
		"  "}

	expectedStr := "Some String"
	expectedStrLen := len(expectedStr)
	testString := "  Some         Stri  ng  "

	actualString, actualStrLen := StrOps{}.StripBadChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
		return
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
		return
	}
}

func TestStrOps_StripLeadingChars_001(t *testing.T) {

	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}

	expectedStr := "SomeString"
	expectedStrLen := len(expectedStr)
	testString := "..........      ./../.\\.\\..\\////   " + expectedStr

	actualString, actualStrLen := StrOps{}.StripLeadingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StripLeadingChars_002(t *testing.T) {

	badChars := make([]string, 0)

	expectedStr := "SomeString"

	testString := "..........      ./../.\\.\\..\\////   " + expectedStr

	expectedStr = testString
	expectedStrLen := len(expectedStr)

	actualString, actualStrLen := StrOps{}.StripLeadingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StripLeadingChars_003(t *testing.T) {

	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}

	expectedStr := ""
	expectedStrLen := len(expectedStr)
	testString := expectedStr

	actualString, actualStrLen := StrOps{}.StripLeadingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StripTrailingChars_001(t *testing.T) {

	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}

	expectedStr := "SomeString"
	expectedStrLen := len(expectedStr)
	testString := expectedStr + "..........      ./../.\\.\\..\\////   "

	actualString, actualStrLen := StrOps{}.StripTrailingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StripTrailingChars_002(t *testing.T) {

	badChars := make([]string, 0)

	expectedStr := "SomeString"

	testString := expectedStr + "..........      ./../.\\.\\..\\////   "

	expectedStr = testString
	expectedStrLen := len(expectedStr)

	actualString, actualStrLen := StrOps{}.StripTrailingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StripTrailingChars_003(t *testing.T) {

	badChars := []string{
		" ",
		"/",
		"//",
		"\\\\",
		"\\",
		".\\",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"//",
		"///",
		"////",
		"..."}

	expectedStr := ""
	expectedStrLen := len(expectedStr)
	testString := expectedStr

	actualString, actualStrLen := StrOps{}.StripTrailingChars(testString, badChars)

	if expectedStr != actualString {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
	}

	if expectedStrLen != actualStrLen {
		t.Errorf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}
}

func TestStrOps_StrCenterInStr_01(t *testing.T) {
	strToCenter := "1234567"
	fieldLen := 79
	exLeftPadLen := 36
	exRightPadLen := 36
	exTotalLen := 79

	leftPad := strings.Repeat(" ", exLeftPadLen)
	rightPad := strings.Repeat(" ", exRightPadLen)
	exStr := leftPad + strToCenter + rightPad

	su := StrOps{}
	str, err := su.StrCenterInStr(strToCenter, fieldLen)
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

func TestStrOps_StrCenterInStr_02(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 15
	exLeftPadLen := 5
	exRightPadLen := 5
	exTotalLen := 15

	leftPad := strings.Repeat(" ", exLeftPadLen)
	rightPad := strings.Repeat(" ", exRightPadLen)
	exStr := leftPad + strToCenter + rightPad

	su := StrOps{}
	str, err := su.StrCenterInStr(strToCenter, fieldLen)
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

func TestStrOps_StrCenterInStr_03(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 5
	exTotalLen := 5

	exStr := strToCenter

	su := StrOps{}
	str, err := su.StrCenterInStr(strToCenter, fieldLen)
	if err != nil {
		t.Errorf("StrCenterInStr() generated error: %v", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}
}

func TestStrOps_StrCenterInStr_04(t *testing.T) {
	strToCenter := "Hello World"
	fieldLen := 5

	_, err := StrOps{}.StrCenterInStr(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrOps{}.StrCenterInStr(strToCenter, fieldLen)\n" +
			"because 'fieldLen' is less than the length of 'strToCenter'.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_StrCenterInStr_05(t *testing.T) {
	strToCenter := "     "
	fieldLen := 15

	_, err := StrOps{}.StrCenterInStr(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrOps{}.StrCenterInStr(strToCenter, fieldLen)\n" +
			"because 'strToCenter' consists entirely of white space.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_StrLeftJustify_01(t *testing.T) {
	strToJustify := "1234567"
	fieldLen := 45
	exTotalLen := fieldLen
	exRightPad := strings.Repeat(" ", 38)
	exStr := strToJustify + exRightPad

	su := StrOps{}
	str, err := su.StrLeftJustify(strToJustify, fieldLen)
	if err != nil {
		t.Errorf("StrLeftJustify() generated error:\n"+
			"%v", err.Error())
	}

	l1 := su.StrGetRuneCnt(str)

	if l1 != exTotalLen {
		t.Error(fmt.Sprintf("Expected total str length '%v', got", exTotalLen), l1)
	}

	if str != exStr {
		t.Error(fmt.Sprintf("Strings did not match. Expected string '%v', got ", exStr), str)
	}

}

func TestStrOps_StrLeftJustify_02(t *testing.T) {
	strToJustify := "      "
	fieldLen := 45

	_, err := StrOps{}.StrLeftJustify(strToJustify, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrLeftJustify(strToJustify, fieldLen)\n" +
			"because parameter, 'strToJustify', is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_StrLeftJustify_03(t *testing.T) {
	strToJustify := "Hello"
	fieldLen := len(strToJustify)

	justifiedStr, err := StrOps{}.StrLeftJustify(strToJustify, fieldLen)

	if err != nil {
		t.Errorf("Error returnd by StrLeftJustify(strToJustify, fieldLen)\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if strToJustify != justifiedStr {
		t.Errorf("Error: Expected justified string='%v'.\n"+
			"Instead, justified string='%v'\n",
			strToJustify, justifiedStr)
	}
}

func TestStrOps_StrLeftJustify_04(t *testing.T) {
	strToJustify := "Hello"
	fieldLen := 2

	_, err := StrOps{}.StrLeftJustify(strToJustify, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrLeftJustify(strToJustify, fieldLen)\n" +
			"because parameter, 'fieldLen', is less than the length of 'strToJustify'.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_StrRightJustify_01(t *testing.T) {

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

func TestStrOps_StrRightJustify_02(t *testing.T) {

	strToJustify := "   "
	fieldLen := 45

	_, err := StrOps{}.StrRightJustify(strToJustify, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrRightJustify(strToJustify, fieldLen)\n" +
			"because parameter 'strToJustify' consists entirely of blank spaces.\n" +
			"However, NO ERROR WAS RETURNED!!")
	}
}

func TestStrOps_StrRightJustify_03(t *testing.T) {

	strToJustify := "1234567"
	fieldLen := len(strToJustify)

	actualStr, err := StrOps{}.StrRightJustify(strToJustify, fieldLen)

	if err != nil {
		t.Error("StrRightJustify() generated error: ", err.Error())
	}

	if strToJustify != actualStr {
		t.Errorf("Error: Expected final string='%v'\n"+
			"Instead, final string='%v'\n",
			strToJustify, actualStr)
	}
}

func TestStrOps_StrRightJustify_04(t *testing.T) {

	strToJustify := "1234567"
	fieldLen := 6

	_, err := StrOps{}.StrRightJustify(strToJustify, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrRightJustify(strToJustify, fieldLen)\n" +
			"because parameter, 'fieldLen' is less than the length of 'strToJustify'.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}

}

func TestStrOps_StrCenterInStrLeft_01(t *testing.T) {
	strToCenter := "1234567"
	fieldLen := 79
	exPadLen := 36
	exTotalLen := 43

	exStr := strings.Repeat(" ", exPadLen) + strToCenter

	su := StrOps{}

	str, err := su.StrCenterInStrLeft(strToCenter, fieldLen)

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

func TestStrOps_StrCenterInStrLeft_02(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 15
	exTotalFinalStrLen := 10

	expectedReturnedStr := "     Hello"
	expectedReturnedPrintStr := strings.ReplaceAll(expectedReturnedStr, " ", "@")

	su := StrOps{}

	actualStr, err := su.StrCenterInStrLeft(strToCenter, fieldLen)

	if err != nil {
		t.Errorf("Error returned by StrCenterInStrLeft()\n"+
			"Error='%v'", err.Error())
	}

	lenActualStr := len(actualStr)

	actualPrintStr := strings.ReplaceAll(actualStr, " ", "@")

	if expectedReturnedStr != actualStr {
		t.Errorf("Error: Expected string='%v'\n"+
			"Instead, string='%v'\n"+
			"Note: Spaces have been replaced with '@'.\n",
			expectedReturnedPrintStr, actualPrintStr)
	}

	if exTotalFinalStrLen != lenActualStr {
		t.Errorf("Error: Expected final string length='%v'.\n"+
			"Instead, final string length='%v'.\n",
			exTotalFinalStrLen, lenActualStr)
	}
}

func TestStrOps_StrCenterInStrLeft_03(t *testing.T) {
	strToCenter := "  "
	fieldLen := 15

	_, err := StrOps{}.StrCenterInStrLeft(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrCenterInStrLeft()\n" +
			"because parameter, 'strToCenter', consists entirely of blank spaces.\n" +
			"However, NO ERROR WAS RETURNED!\n\n")
	}

}

func TestStrOps_StrCenterInStrLeft_04(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 2

	_, err := StrOps{}.StrCenterInStrLeft(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected error return from StrCenterInStrLeft(strToCenter, fieldLen) because\n" +
			"input parameter, 'fieldLen' is less than the length of 'strToCenter'.\n" +
			"However, NO ERROR WAS RETURNED!!")
	}
}

func TestStrOps_StrGetRuneCnt_01(t *testing.T) {
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

func TestStrOps_StrPadLeftToCenter_01(t *testing.T) {
	strToCenter := "1234567"
	fieldLen := 79
	exLen := 36
	su := StrOps{}
	padStr, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

	if err != nil {
		t.Error("Error on StrPadLeftToCenter(), got", err.Error())
	}

	l1 := su.StrGetRuneCnt(padStr)

	if l1 != exLen {
		t.Error(fmt.Sprintf("Expected pad length of '%v', got ", exLen), l1)
	}

}

func TestStrOps_StrPadLeftToCenter_02(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 15
	exLen := 5
	su := StrOps{}
	padStr, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

	if err != nil {
		t.Error("Error on StrPadLeftToCenter(), got", err.Error())
	}

	l1 := su.StrGetRuneCnt(padStr)

	if l1 != exLen {
		t.Error(fmt.Sprintf("Expected pad length of '%v', got ", exLen), l1)
	}
}

func TestStrOps_StrPadLeftToCenter_03(t *testing.T) {
	strToCenter := "   "
	fieldLen := 15
	su := StrOps{}
	_, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrPadLeftToCenter(strToCenter, fieldLen)\n" +
			"because 'strToCenter' consists entirely of white space." +
			"However, NO ERROR WAS RETURNED!")
	}
}

func TestStrOps_StrPadLeftToCenter_04(t *testing.T) {
	strToCenter := "Hello World"
	fieldLen := 5
	su := StrOps{}
	_, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

	if err == nil {
		t.Error("Expected an error return from StrPadLeftToCenter(strToCenter, fieldLen)\n" +
			"because 'fieldLen' is less than the length of 'strToCenter'." +
			"However, NO ERROR WAS RETURNED!")
	}
}

func TestStrOps_StrPadLeftToCenter_05(t *testing.T) {
	strToCenter := "Hello"
	fieldLen := 5
	exLen := 0
	su := StrOps{}

	padStr, err := su.StrPadLeftToCenter(strToCenter, fieldLen)

	if err != nil {
		t.Error("Error on StrPadLeftToCenter(), got", err.Error())
	}

	lenPadStr := len(padStr)

	if exLen != lenPadStr {
		t.Errorf("Error: Expected length of Pad String='0'.\n"+
			"Instead, length of Pad String='%v'", lenPadStr)
	}

}

func TestStrOps_SwapRune_01(t *testing.T) {
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

func TestStrOps_SwapRune_02(t *testing.T) {
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

func TestStrOps_SwapRune_03(t *testing.T) {
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

func TestStrOps_SwapRune_04(t *testing.T) {
	su := StrOps{}

	tStr := "xHello World"
	expected := "XHello World"
	result, err := su.SwapRune(tStr, 'x', 'X')

	if err != nil {
		t.Errorf("Error returned from SwapRune:\n%v", err.Error())
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

func TestStrOps_SwapRune_05(t *testing.T) {

	tStr := ""

	newStr, err := StrOps{}.SwapRune(tStr, 'x', 'X')

	if err != nil {
		t.Errorf("Error returned from SwapRune:\n%v", err.Error())
	}

	if newStr != tStr {
		t.Errorf("Error: Expected StrOps{}.SwapRune(tStr, 'x', 'X') would return\n"+
			"an empty string, because 'tStr' is an empty string.\n"+
			"Instead, newStr='%v'", newStr)
	}

}

func TestStrOps_TrimMultipleChars_01(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_02(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_03(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_04(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_05(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_06(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_07(t *testing.T) {
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

func TestStrOps_TrimMultipleChars_08(t *testing.T) {
	tStr := ""

	_, err := StrOps{}.TrimMultipleChars(tStr, ' ')

	if err == nil {
		t.Error("Expected an error return from StrOps{}.TrimMultipleChars(tStr, ' ')\n" +
			"because 'tStr' is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_TrimMultipleChars_09(t *testing.T) {
	tStr := "Hello World"
	replaceRune := rune(0)
	_, err := StrOps{}.TrimMultipleChars(tStr, replaceRune)

	if err == nil {
		t.Error("Expected an error return from StrOps{}.TrimMultipleChars(tStr, replaceRune)\n" +
			"because 'replaceRune' has a zero value.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
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

func TestStrOps_UpperCaseFirstLetter_01(t *testing.T) {

	testStr := "now is the time for all good men to come to the aid of their country."

	expected := "Now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.UpperCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_UpperCaseFirstLetter_02(t *testing.T) {

	testStr := "  now is the time for all good men to come to the aid of their country."

	expected := "  Now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.UpperCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_UpperCaseFirstLetter_03(t *testing.T) {

	testStr := "Now is the time for all good men to come to the aid of their country."

	expected := "Now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.UpperCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_UpperCaseFirstLetter_04(t *testing.T) {

	testStr := "  Now is the time for all good men to come to the aid of their country."

	expected := "  Now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.UpperCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_UpperCaseFirstLetter_05(t *testing.T) {

	testStr := ""

	expected := ""

	actualStr := StrOps{}.UpperCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_Write_01(t *testing.T) {

	originalStr := "Original base string written to sops1"

	sops1 := StrOps{}.NewPtr()

	lenOriginalStr := len(originalStr)

	p := []byte(originalStr)

	n, err := sops1.Write(p)

	if err != nil {
		t.Errorf("Error returned by sops1.Write(p). Error='%v' \n",
			err.Error())
	}

	actualStr := sops1.GetStringData()

	if originalStr != actualStr {
		t.Errorf("Error: Expected string='%v'. Instead, string='%v'. \n",
			originalStr, actualStr)
	}

	if lenOriginalStr != n {
		t.Errorf("Error: Expected Length='%v'. Instead, Bytes Written='%v'. \n",
			lenOriginalStr, n)
	}

}

func TestStrOps_Write_02(t *testing.T) {

	originalStr := "Hello World"

	sops1 := StrOps{}.NewPtr()

	p := make([]byte, 3)

	for i := 0; i < 4; i++ {

		if i == 0 {
			p[0] = 'H'
			p[1] = 'e'
			p[2] = 'l'
		} else if i == 1 {
			p[0] = 'l'
			p[1] = 'o'
			p[2] = ' '
		} else if i == 2 {
			p[0] = 'W'
			p[1] = 'o'
			p[2] = 'r'

		} else if i == 3 {
			p[0] = 'l'
			p[1] = 'd'
			p[2] = byte(0)

		}

		_, err := sops1.Write(p)

		if err != nil {
			t.Errorf("Error returned by sops1.Write(p). Error='%v' ", err.Error())
			return
		}
	}

	actualStr := sops1.GetStringData()

	if originalStr != actualStr {
		t.Errorf("Error: Expected final string='%v'. Instead, string='%v'. ",
			originalStr, actualStr)
	}

	if 11 != len(actualStr) {
		t.Errorf("Error: Expected Length='11'. Instead, Length='%v'. ",
			len(actualStr))
	}

}

func TestStrOps_Write_03(t *testing.T) {

	originalStr := "Original base string written to sops1"

	lenOriginalStr := len(originalStr)

	sops1 := StrOps{}.NewPtr()

	sops1.SetStringData(originalStr)

	sops2 := StrOps{}.NewPtr()

	n, err := io.Copy(sops2, sops1)

	if err != nil {
		t.Errorf("Error returned by io.Copy(sops2, sops1). Error='%v' \n", err.Error())
		return
	}

	if int64(lenOriginalStr) != n {
		t.Errorf("Error: Expected bytes copied='%v'. Instead, bytes copied='%v'. ",
			lenOriginalStr, n)
	}

	actualStr := sops2.GetStringData()

	if originalStr != actualStr {
		t.Errorf("Error: Expected string='%v'. Instead, string='%v'. ",
			originalStr, actualStr)
	}
}

func TestStrOps_Write_04(t *testing.T) {

	originalStr := "Original base string written to sops1"

	sops1 := StrOps{}.NewPtr()
	sops1.SetStringData(originalStr)

	p := make([]byte, 0)

	_, err := sops1.Write(p)

	if err == nil {
		t.Error("Error: Expected Error Return. NO ERROR WAS RETURNED!")
	}
}
