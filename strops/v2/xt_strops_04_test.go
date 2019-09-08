package strops

import "testing"

func TestStrOps_ExtractNumericDigits_01(t *testing.T) {

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "12"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := ""
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedLeadingSignCharIndex := -1
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_02(t *testing.T) {

	targetStr := "Etc/GMT+11"
	startIndex := 0
	keepLeadingChars := "+"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+11"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)

	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_03(t *testing.T) {

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := 23
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+0000"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := "+"
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedLeadingSignCharIndex := 0
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_04(t *testing.T) {

	targetStr := "2016 1:6:3pm +0000 UTC"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "2016"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_05(t *testing.T) {

	targetStr := "2016"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "2016"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_06(t *testing.T) {

	targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
	startIndex := 0
	keepLeadingChars := "$("
	keepInteriorChars := ",."
	keepTrailingChars := ")"

	expectedNumStr := "$(1,250,364.33)"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_07(t *testing.T) {

	targetStr := "Hello World! The time zone here is 'Etc/GMT+11'. What do you think?"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+11"
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_08(t *testing.T) {

	targetStr := "Etc/GMT-4"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "-4"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_09(t *testing.T) {

	targetStr := "+$697,621,911.77"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "+$697,621,911.77"
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_10(t *testing.T) {

	targetStr := "Hello World\t+-$697,621,911.77\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "-$697,621,911.77"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_11(t *testing.T) {

	targetStr := "Hello World\t\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := ""
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := -1
	expectedNextTargetStrIndex := -1

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_12(t *testing.T) {

	targetStr := ""
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := ""
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := 0
	expectedNumIdx := -1
	expectedNextTargetStrIndex := -1

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_13(t *testing.T) {

	targetStr := "Hello World7Have a great day!"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_14(t *testing.T) {

	targetStr := "7Hello World Have a great day!"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_15(t *testing.T) {

	targetStr := "Hello World Have a great day!7"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_16(t *testing.T) {

	targetStr := "Hello World -7\t6 Have a great day!"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "-7"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrOps_ExtractNumericDigits_17(t *testing.T) {

	targetStr := "Hello World.\t+$-697,621,911.77.\nHow are you.\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "$-697,621,911.77"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
	err := StrOps{}.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

