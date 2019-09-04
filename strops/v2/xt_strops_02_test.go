package strops

import (
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
/*
func TestStrOps_ExtractDataField_01(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := "America/Chicago"
  expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := '\t'
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
  
  expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)

  expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

  if expectedNextTargetIdx > expectedLastGoodIdx {
    expectedNextTargetIdx = -1
  }

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}

func TestStrOps_ExtractDataField_02(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " Zone:\t America/Chicago Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := "America/Chicago"
  expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
  expectedDataFieldTrailingDelimiter := ' '
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()

  expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength


  if expectedNextTargetIdx > expectedLastGoodIdx {
    expectedNextTargetIdx = -1
  }

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}

func TestStrOps_ExtractDataField_03(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " America/Chicago Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--

  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := ""
  expectedDataFieldIdx := -1
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := '\n'
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()

  expectedLeadingKeyWordDelimiterIndex := -1
  expectedNextTargetIdx := -1

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}

func TestStrOps_ExtractDataField_04(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " Zone:\t America/Chicago Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--
  startIdx := 6
  leadingKeyWordDelimiter := ""
  expectedLeadingKeyWordDelimiterIndex := -1
  expectedDataFieldStr := "America/Chicago"
  expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := ' '
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()

  expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

  if expectedNextTargetIdx > expectedLastGoodIdx {
    expectedNextTargetIdx = -1
  }

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}


func TestStrOps_ExtractDataField_05(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " Zone:\t #America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.Index(targetStr, "#")
  expectedLastGoodIdx--

  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := ""
  expectedDataFieldIdx := -1
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := '#'
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()

  expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
  expectedNextTargetIdx := -1

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}

func TestStrOps_ExtractDataField_06(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " #Zone:\t America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := ""
  expectedDataFieldIdx := -1
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := '#'
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
  expectedLeadingKeyWordDelimiterIndex := -1
  expectedNextTargetIdx := -1

  datDto,
  err := StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    t.Errorf("Error returned by StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}

func TestStrOps_ExtractDataField_07(t *testing.T) {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := "\tZone:\tAmerica/Chicago\t\tZone:\tAmerica/New_York\t\tZone:\tAmerica/Los_Angeles\n"
  expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
  expectedLastGoodIdx--
  lenTargetStr := len(targetStr)
  startIdx := 0
  expectedStartIdx := 46
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := "America/Los_Angeles"
  expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedDataFieldLength := len(expectedDataFieldStr)
  expectedDataFieldTrailingDelimiter := '\n'
  expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()
  expectedLeadingKeyWordDelimiterIndex := strings.LastIndex(targetStr, leadingKeyWordDelimiter)
  expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

  if expectedNextTargetIdx > expectedLastGoodIdx {
    expectedNextTargetIdx = -1
  }

  var datDto DataFieldProfileDto
  var err error

  for i:=0; i < 3; i++ {

    datDto,
      err = StrOps{}.ExtractDataField(
      targetStr,
      leadingKeyWordDelimiter,
      startIdx,
      leadingRunes,
      trailingRunes,
      endOfLineRunes)

    if err != nil {
      t.Errorf("Error returned by StrOps{}.ExtractDataField()\n" +
        "Cycle No='%v'\n"+
        "targetStr='%v'\tstartIdx='%v'\n"+
        "Error='%v'\n", i, targetStr, startIdx, err.Error())
      return
    }

    startIdx = datDto.NextTargetStrIndex
  }

  if targetStr  != datDto.TargetStr {
    t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
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
      expectedDataFieldStr, datDto.DataFieldStr )
  }

  if  expectedDataFieldLength != datDto.DataFieldLength {
    t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedDataFieldLength, datDto.DataFieldLength )
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

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
  }
}
*/

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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
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
    t.Errorf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    t.Errorf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
  }

  if expectedNumStr != nStrDto.NumStr {
    t.Errorf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    t.Errorf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    t.Errorf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
  }

  if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
    t.Errorf("Expected leading sign char index ='%v'\n" +
      "Instead, leading sign char index ='%v'\n",
      expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
  }

  if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
    t.Errorf("Expected next target index after number string ='%v'\n" +
      "Instead, next target string index ='%v'\n",
      expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
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
    return
  }

  sExtract := targetStr[idx[0]:idx[1]]

  if expected != sExtract {
    t.Errorf("Error: Expected regular expression match on string='%v'. "+
      "Instead, matched string='%v'. ", expected, sExtract)
  }
}

func TestStrOps_GetReader_01(t *testing.T) {
  originalStr := "Now is the time for all good men to come to the aid of their country."
  s1 := StrOps{}.NewPtr()
  s1.SetStringData(originalStr)
  s2 := StrOps{}.NewPtr()
  rdr := s1.GetReader()
  n, err := io.Copy(s2, rdr)

  if err != nil {
    t.Errorf("Error returned by io.Copy(s2, s1.GetReader()). "+
      "Error='%v' ", err.Error())
  }

  actualStr := s2.GetStringData()

  if originalStr != actualStr {
    t.Errorf("Error: Expected actualStr='%v'. Instead, actualStr='%v'",
      originalStr, actualStr)
  }

  if int64(len(originalStr)) != n {
    t.Errorf("Error: Expected characters read='%v'. Instead, "+
      "characters read='%v' ",
      len(originalStr), n)
  }

}

func TestStrOps_GetReader_02(t *testing.T) {
  originalStr := "xx"
  s1 := StrOps{}.NewPtr()
  s1.SetStringData(originalStr)
  s2 := StrOps{}.NewPtr()
  rdr := s1.GetReader()
  n, err := io.Copy(s2, rdr)

  if err != nil {
    t.Errorf("Error returned by io.Copy(s2, s1.GetReader()). "+
      "Error='%v' ", err.Error())
  }

  actualStr := s2.GetStringData()

  if originalStr != actualStr {
    t.Errorf("Error: Expected actualStr='%v'. Instead, actualStr='%v'",
      originalStr, actualStr)
  }

  if int64(len(originalStr)) != n {
    t.Errorf("Error: Expected characters read='%v'. Instead, "+
      "characters read='%v' ",
      len(originalStr), n)
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

func TestStrOps_LowerCaseFirstLetter_01(t *testing.T) {

  testStr := "Now is the time for all good men to come to the aid of their country."

  expected := "now is the time for all good men to come to the aid of their country."

  actualStr := StrOps{}.LowerCaseFirstLetter(testStr)

  if expected != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expected, actualStr)
  }

}

func TestStrOps_LowerCaseFirstLetter_02(t *testing.T) {

  testStr := "  Now is the time for all good men to come to the aid of their country."

  expected := "  now is the time for all good men to come to the aid of their country."

  actualStr := StrOps{}.LowerCaseFirstLetter(testStr)

  if expected != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expected, actualStr)
  }

}

func TestStrOps_LowerCaseFirstLetter_03(t *testing.T) {

  testStr := "now is the time for all good men to come to the aid of their country."

  expected := "now is the time for all good men to come to the aid of their country."

  actualStr := StrOps{}.LowerCaseFirstLetter(testStr)

  if expected != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expected, actualStr)
  }

}

func TestStrOps_LowerCaseFirstLetter_04(t *testing.T) {

  testStr := "  now is the time for all good men to come to the aid of their country."

  expected := "  now is the time for all good men to come to the aid of their country."

  actualStr := StrOps{}.LowerCaseFirstLetter(testStr)

  if expected != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expected, actualStr)
  }

}

func TestStrOps_LowerCaseFirstLetter_05(t *testing.T) {

  testStr := ""

  expected := ""

  actualStr := StrOps{}.LowerCaseFirstLetter(testStr)

  if expected != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expected, actualStr)
  }

}
