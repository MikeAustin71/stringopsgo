package strops

import (
	"strings"
	"testing"
)

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

func TestStrOps_ExtractDataField_09(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"

	_,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'leadingFieldDelimiters' is an empty string array.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_ExtractDataField_10(t *testing.T) {

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"

	_,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'trailingFieldDelimiters' is an empty string array.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_ExtractDataField_11(t *testing.T) {

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

	targetStr := ""
	startIdx := 0
	leadingKeyWordDelimiter := "Zone:"

	_,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'targetStr' is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_ExtractDataField_12(t *testing.T) {

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
	startIdx := -1
	leadingKeyWordDelimiter := "Zone:"

	_,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'startIdx' is less than zero.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_ExtractDataField_13(t *testing.T) {

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
	startIdx :=  999
	leadingKeyWordDelimiter := "Zone:"

	_,
	err := StrOps{}.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'startIdx' is exceeds the outer boundary of 'targetStr'.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_ExtractDataField_14(t *testing.T) {

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
		"\v"}

	targetStr := "Good morning America!"

	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := lenTargetStr - 1
	expectedEndOfLineDelimiterIdx := -1
	startIdx := 0
	leadingKeyWordDelimiter := ""
	expectedDataFieldStr := "Good morning America!"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := ""
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfString()
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedEndOfLineDelimiter := ""
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

