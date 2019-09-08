package strops

import (
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

func TestStrOps_ExtractDataField_01(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\t"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedDataFieldTrailingDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.DataFieldTrailingDelimiter), true))
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrOps_ExtractDataField_02(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}

	commentDelimiters := []string{"#"}

	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := " "
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrOps_ExtractDataField_03(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}
	targetStr := " America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()

	expectedLeadingKeyWordDelimiterIndex := -1
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := -1

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrOps_ExtractDataField_04(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 6
	leadingKeyWordDelimiter := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := " "
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrOps_ExtractDataField_05(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t #America/Chicago\t Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.Index(targetStr, "#")
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()

	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
	expectedNextTargetIdx := -1
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr, "\n")

	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrOps_ExtractDataField_06(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " #Zone:\t America/Chicago\t Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "#")
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedNextTargetIdx := -1
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr, "\n")
	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")

	datDto,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrOps_ExtractDataField_07(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "\tZone:\tAmerica/Chicago\t\tZone:\tAmerica/New_York\t\tZone:\tAmerica/Los_Angeles\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 0
	expectedStartIdx := 46
	leadingKeyWordDelimiter := "Zone:"
	expectedDataFieldStr := "America/Los_Angeles"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()
	expectedLeadingKeyWordDelimiterIndex := strings.LastIndex(targetStr, leadingKeyWordDelimiter)
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	var datDto DataFieldProfileDto
	var err error

	for i := 0; i < 3; i++ {

		datDto,
			err = StrOps{}.ExtractDataField(
			targetStr,
			leadingKeyWordDelimiter,
			startIdx,
			leadingFieldDelimiters,
			trailingFieldDelimiters,
			commentDelimiters,
			endOfLineDelimiters)

		if err != nil {
			t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
				"Cycle No='%v'\n"+
				"targetStr='%v'\tstartIdx='%v'\n"+
				"Error='%v'\n", i, targetStr, startIdx, err.Error())
			return
		}

		startIdx = datDto.NextTargetStrIndex
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if expectedStartIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			expectedStartIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrOps_ExtractDataField_08(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "\tZone:\tAmerica/Chicago\t\t#Zone:\tAmerica/New_York\t\tZone:\tAmerica/Los_Angeles\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "#")
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 3
	expectedStartIdx := 3
	leadingKeyWordDelimiter := "Zone:"
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr,"\n")
	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")
	expectedNextTargetIdx := -1

	var datDto DataFieldProfileDto
	var err error

	datDto,
		err = StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if expectedStartIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			expectedStartIdx, datDto.TargetStrStartIndex)
	}

	if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedEndOfLineDelimiter), false),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrOps{}.ConvertNonPrintableCharacters([]rune(expectedCommentDelimiter), true),
			StrOps{}.ConvertNonPrintableCharacters([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

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
