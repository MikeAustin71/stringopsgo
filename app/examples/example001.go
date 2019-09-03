package examples

import (
  "MikeAustin71/stringopsgo/strops/v2"
  "errors"
  "fmt"
  "io"
  "os"
  "regexp"
  "sort"
  "strconv"
  "strings"
  "time"
)


type MainTest struct {
  input string
}

func (mt MainTest) ExampleExtractDataField01() {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := ""
  expectedFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedFieldLength := len(expectedDataFieldStr)
  expectedLeadingKeyWordDelimiterIndex := -1

  expectedNextTargetIdx := expectedFieldIdx + expectedFieldLength

  if expectedNextTargetIdx >= len(targetStr) {
    expectedNextTargetIdx = -1
  }


  datDto,
  err := strops.StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    fmt.Printf("Error returned by strops.StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  isError := false

  if targetStr  != datDto.TargetStr {
    fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
    isError = true
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
    isError = true
  }

  if startIdx != datDto.StartIndex {
    fmt.Printf("ERROR: Expected datDto.StartIndex='%v'.\n"+
      "Instead, datDto.StartIndex='%v'.\n",
      startIdx, datDto.StartIndex)
    isError = true
  }

  if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
    fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
      "Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
      leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
    isError = true
  }

  if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
    fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
      "Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
      expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiter)
    isError = true
  }

  if expectedDataFieldStr != datDto.DataFieldStr {
    fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
      "Instead, datDto.DataFieldStr='%v'.\n",
      expectedDataFieldStr, datDto.DataFieldStr )
    isError = true
  }

  if  expectedFieldLength != datDto.DataFieldLength {
    fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedFieldLength, datDto.DataFieldLength )
    isError = true
  }

  if expectedFieldIdx != datDto.DataFieldIndex {
    fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
      "Instead, datDto.DataFieldIndex='%v'.\n",
      expectedFieldIdx, datDto.DataFieldIndex)
    isError = true
  }

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
    isError = true
  }

  fmt.Println("================================================")
  fmt.Println("           ExampleExtractDataField01            ")
  fmt.Println("================================================")
  if isError {
    fmt.Println("              @@@@ FAILURE @@@@                 ")
  } else {
    fmt.Println("                   SUCCESS!                     ")
  }
  fmt.Println("------------------------------------------------")
  fmt.Println("                    Base Data                   ")
  fmt.Println("------------------------------------------------")
  fmt.Printf("             TargetStr: %v", targetStr)
  fmt.Println("      TargetStr Length: ", lenTargetStr)
  fmt.Println("           Start Index: ", startIdx)
  fmt.Println("    Key Word Delimiter: ", leadingKeyWordDelimiter)
  fmt.Println("Key Word Delimiter Idx: ", expectedLeadingKeyWordDelimiterIndex)
  fmt.Println("------------------------------------------------")
  fmt.Println("                 Expected Results               ")
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Field String: ", expectedDataFieldStr)
  fmt.Println("              Field Str Length: ", expectedFieldLength)
  fmt.Println("                   Field Index: ", expectedFieldIdx)
  fmt.Println("             Next Target Index: ", expectedNextTargetIdx)
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Actual Results                ")
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Field String: ", datDto.DataFieldStr)
  fmt.Println("              Field Str Length: ", datDto.DataFieldLength)
  fmt.Println("                   Field Index: ", datDto.DataFieldIndex)
  fmt.Println("             Next Target Index: ", datDto.NextTargetStrIndex)
  fmt.Println("                 Target String: ", datDto.TargetStr)
  fmt.Println("             Target Str Length: ", datDto.TargetStrLength)
  fmt.Println("               Target StartIdx: ", datDto.StartIndex)
  fmt.Println("    Leading Key Delimiter Word: ", datDto.LeadingKeyWordDelimiter)
  fmt.Println("Leading Key Word Delimiter Idx: ", datDto.LeadingKeyWordDelimiterIndex)

}


func (mt MainTest) ExampleExtractDataField02() {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune("\t \r\f\n\v")
  targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedDataFieldStr := "America/Chicago"
  expectedFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
  expectedFieldLength := len(expectedDataFieldStr)
  expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, leadingKeyWordDelimiter)
  expectedNextTargetIdx := expectedFieldIdx + expectedFieldLength

  if expectedNextTargetIdx >= len(targetStr) {
    expectedNextTargetIdx = -1
  }

  datDto,
  err := strops.StrOps{}.ExtractDataField(
    targetStr,
    leadingKeyWordDelimiter,
    startIdx,
    leadingRunes,
    trailingRunes,
    endOfLineRunes)

  if err != nil {
    fmt.Printf("Error returned by strops.StrOps{}.ExtractDataField()\n"+
      "targetStr='%v'\tstartIdx='%v'\n"+
      "Error='%v'\n", targetStr, startIdx, err.Error())
    return
  }

  isError := false

  if targetStr  != datDto.TargetStr {
    fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
      "Instead, datDto.TargetStr='%v'.\n",
      targetStr ,datDto.TargetStr)
    isError = true
  }

  if lenTargetStr !=  datDto.TargetStrLength {
    fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
      "Instead, datDto.TargetStrLength='%v'.\n",
      lenTargetStr ,datDto.TargetStrLength)
    isError = true
  }

  if startIdx != datDto.StartIndex {
    fmt.Printf("ERROR: Expected datDto.StartIndex='%v'.\n"+
      "Instead, datDto.StartIndex='%v'.\n",
      startIdx, datDto.StartIndex)
    isError = true
  }

  if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
    fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
      "Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
      leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
    isError = true
  }

  if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
    fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
      "Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
      expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiter)
    isError = true
  }

  if expectedDataFieldStr != datDto.DataFieldStr {
    fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
      "Instead, datDto.DataFieldStr='%v'.\n",
      expectedDataFieldStr, datDto.DataFieldStr )
    isError = true
  }

  if  expectedFieldLength != datDto.DataFieldLength {
    fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
      "Instead, datDto.DataFieldLength='%v'.\n",
      expectedFieldLength, datDto.DataFieldLength )
    isError = true
  }

  if expectedFieldIdx != datDto.DataFieldIndex {
    fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
      "Instead, datDto.DataFieldIndex='%v'.\n",
      expectedFieldIdx, datDto.DataFieldIndex)
    isError = true
  }

  if expectedNextTargetIdx != datDto.NextTargetStrIndex  {
    fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      expectedNextTargetIdx, datDto.NextTargetStrIndex)
    isError = true
  }

  fmt.Println("================================================")
  fmt.Println("           ExampleExtractDataField02            ")
  fmt.Println("================================================")
  if isError {
    fmt.Println("              @@@@ FAILURE @@@@                 ")
  } else {
    fmt.Println("                   SUCCESS!                     ")
  }
  fmt.Println("------------------------------------------------")
  fmt.Println("                    Base Data                   ")
  fmt.Println("------------------------------------------------")
  fmt.Printf("             TargetStr: %v", targetStr)
  fmt.Println("      TargetStr Length: ", lenTargetStr)
  fmt.Println("           Start Index: ", startIdx)
  fmt.Println("    Key Word Delimiter: ", leadingKeyWordDelimiter)
  fmt.Println("Key Word Delimiter Idx: ", expectedLeadingKeyWordDelimiterIndex)
  fmt.Println("------------------------------------------------")
  fmt.Println("                 Expected Results               ")
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Field String: ", expectedDataFieldStr)
  fmt.Println("              Field Str Length: ", expectedFieldLength)
  fmt.Println("                   Field Index: ", expectedFieldIdx)
  fmt.Println("             Next Target Index: ", expectedNextTargetIdx)
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Actual Results                ")
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Field String: ", datDto.DataFieldStr)
  fmt.Println("              Field Str Length: ", datDto.DataFieldLength)
  fmt.Println("                   Field Index: ", datDto.DataFieldIndex)
  fmt.Println("             Next Target Index: ", datDto.NextTargetStrIndex)
  fmt.Println("                 Target String: ", datDto.TargetStr)
  fmt.Println("             Target Str Length: ", datDto.TargetStrLength)
  fmt.Println("               Target StartIdx: ", datDto.StartIndex)
  fmt.Println("    Leading Key Delimiter Word: ", datDto.LeadingKeyWordDelimiter)
  fmt.Println("Leading Key Word Delimiter Idx: ", datDto.LeadingKeyWordDelimiterIndex)

}

func (mt MainTest) exampleExtractNumStr01() {
  // Etc/GMT-4
  // "Etc/GMT+11"
  // "November 12, 2016 1:6:3pm -(+0000) UTC"
  // "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
  targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"

  expectedNumStr := "$(1,250,364.33)"
  expectedLeadingSignChar := ""
  startIndex := 0
  keepLeadingChars := "$(+-"
  keepInteriorChars := ",."
  keepTrailingChars := ")"

  expectedLeadingSignIndex := -1

  expectedNumStrLen := len(expectedNumStr)
  expectedNumIdx := strings.Index(targetStr, expectedNumStr)
  expectedNextTargetStrIdx := expectedNumIdx + expectedNumStrLen

  if expectedNextTargetStrIdx >= len(targetStr) {
    expectedNextTargetStrIdx = -1
  }

  nStrDto,
  err := strops.StrOps{}.ExtractNumericDigits(
    targetStr,
    startIndex,
    keepLeadingChars,
    keepInteriorChars,
    keepTrailingChars)

  if err != nil {
    fmt.Printf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  isError := false

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    fmt.Printf("Expected starting numeric index='%v'\n"+
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
    isError = true
  }

  if expectedNumStr != nStrDto.NumStr {
    fmt.Printf("Expected number string ='%v'\n"+
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
    isError = true
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    fmt.Printf("Expected number string length ='%v'\n"+
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
    isError = true
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    fmt.Printf("Expected leading sign char ='%v'\n"+
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
    isError = true
  }

  if expectedLeadingSignIndex != nStrDto.LeadingSignIndex {
    fmt.Printf("Expected leading sign index ='%v'\n"+
      "Instead, leading sign index ='%v'\n",
      expectedLeadingSignIndex, nStrDto.LeadingSignIndex)
    isError = true
  }

  if expectedNextTargetStrIdx != nStrDto.NextTargetStrIndex {
    fmt.Printf("Expected Next TargetStr Char Index ='%v'\n"+
      "Instead, Next TargetStr Char Index ='%v'\n",
      expectedNextTargetStrIdx, nStrDto.NextTargetStrIndex)
    isError = true
  }

  fmt.Println("  mainTest.ExampleExtractNumStr01()  ")
  fmt.Println("-------------------------------------")
  if isError {
    fmt.Println("     @@@@@  FAILURE @@@@@@           ")
  } else {
    fmt.Println("          SUCCESS!!!")
  }

  fmt.Println("-------------------------------------")
  fmt.Println("          TargetStr: ", targetStr)
  fmt.Println("           startIdx: ", startIndex)
  fmt.Println("-------------------------------------")
  fmt.Println("           Expected                  ")
  fmt.Println("-------------------------------------")
  fmt.Println("       Number Index: ", expectedNumIdx)
  fmt.Println("         Num Length: ", expectedNumStrLen)
  fmt.Println("  Leading Sign Char: ", expectedLeadingSignChar)
  fmt.Println(" Leading Sign Index: ", expectedLeadingSignIndex)
  fmt.Println("      Number String: ", expectedNumStr)
  fmt.Println(" Next TargetStr Idx: ", expectedNextTargetStrIdx)
  fmt.Println("-------------------------------------")
  fmt.Println("            Results                  ")
  fmt.Println("-------------------------------------")
  fmt.Println("        NumberIndex: ", nStrDto.FirstNumCharIndex)
  fmt.Println("         Num Length: ", nStrDto.NumStrLen)
  fmt.Println("  Leading Sign Char: ", nStrDto.LeadingSignChar)
  fmt.Println(" Leading Sign Index: ", nStrDto.LeadingSignIndex)
  fmt.Println("      Number String: ", nStrDto.NumStr)
  fmt.Println("Target Str Next Idx: ", nStrDto.NextTargetStrIndex)
}


func (mt MainTest) exampleStripLeadingChars01() {

  badChars := []string {
    " ",
    "/",
    "//",
    "../",
    ".",
    "..\\",
    "\\\\\\",
    "..",
    "./",
    "///",
    "..."}
  expectedStr := "SomeString"
  expectedStrLen := len(expectedStr)
  testString := "..........      ./../.\\.\\..\\////   " + expectedStr

  actualString, actualStrLen := strops.StrOps{}.StripLeadingChars(testString, badChars)

  if expectedStr != actualString {
    fmt.Printf("ERROR: Expected result string='%v'\n" +
      "Instead, result string='%v'\n",
      expectedStr, actualString)
    return
  }

  if expectedStrLen != actualStrLen {
    fmt.Printf("ERROR: Expected result string length='%v'\n" +
      "Instead, result string length='%v'\n",
      expectedStrLen, actualStrLen)
  }

}

func (mt MainTest) exampleSortStrLenHighestToLowest01() {
  badChars := []string {
    "aaaaa",
    "bbbbb",
    "cccccccccc",
    "z",
    "fffffffffff",
    "xx",
    "ddddddddd",
    "eeeeeeeeeee" }

  fmt.Println("Sort by Length Highest To Lowest")
  fmt.Println("          Unordered List")
  fmt.Println("================================")
  fmt.Println()

  for i:=0; i < len(badChars); i++ {
    fmt.Printf("%3d. %v\n", i+1, badChars[i])
  }

  sort.Sort(strops.SortStrLengthHighestToLowest(badChars))

  fmt.Println()
  fmt.Println("================================")
  fmt.Println("Sort by Length Highest To Lowest")
  fmt.Println("          Ordered List")
  fmt.Println("================================")
  fmt.Println()

  for i:=0; i < len(badChars); i++ {
    fmt.Printf("%3d. %v\n", i+1, badChars[i])
  }

}

func (mt MainTest) exampleSortStrLenLowestToHighest01() {

  badChars := []string {
    "aaaaa",
    "bbbbb",
    "cccccccccc",
    "z",
    "fffffffffff",
    "xx",
    "ddddddddd",
    "eeeeeeeeeee" }

  fmt.Println("Sort by Length Lowest To Highest")
  fmt.Println("          Unordered List")
  fmt.Println("================================")
  fmt.Println()

  for i:=0; i < len(badChars); i++ {
    fmt.Printf("%3d. %v\n", i+1, badChars[i])
  }

  sort.Sort(strops.SortStrLengthLowestToHighest(badChars))

  fmt.Println()
  fmt.Println("================================")
  fmt.Println("Sort by Length Lowest To Highest")
  fmt.Println("          Ordered List")
  fmt.Println("================================")
  fmt.Println()

  for i:=0; i < len(badChars); i++ {
    fmt.Printf("%3d. %v\n", i+1, badChars[i])
  }

}


func (mt MainTest) ExampleExpressions01() {
  //AMpm Match \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
  //PM Match V1 "\\d{1}\\s?(?i)p[.\\s]*(?i)m[.]*"
  //PM Match V2 "\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*"
  //AM Match V1 "\\d{1}\\s?(?i)a[.\\s]*(?i)m[.]*"
  //AM Match V2 "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"
  regexAMpm := "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"

  samples := []string{
    "12:15 AM",
    "03:25AM",
    "11:19 A M",
    "11:19 A. M.",
    "11:19 A.M.",
    "11:19A.M.",
    "11:19  A.M.",
    "11:19  AM",
    "11:19  AM",
    "11:19  A  M",
    "12:15 am",
    "03:25am",
    "11:19 a m",
    "11:19a m",
    "11:19 a. m.",
    "11:19 a.m.",
    "11:19a.m.",
    "11:19  A  M",
    "11:19  A. M.",
    "11:19  a  m",
    "11:19  a. m.",
    "11:19 m",
    "11:19 a",
    "10:25 PM",
    "02:15PM",
    "10:18 P M",
    "01:19 P. M.",
    "12:19 P.M.",
    "10:19P.M.",
    "10:15 pm",
    "04:25pm",
    "10:19 p m",
    "10:19p m",
    "10:19 p. m.",
    "10:19p.m.",
    "15:35:03",
    "10:19:16 p.m.",
    "10:15 pm -0600 MST",
    "10:15 pm-0600 MST",
    "10:15 pm PST",
    "10:15  pm -0600 MST",
    "10:15 p.m -0600 MST",
    "10:15 pm. -0600 MST",
    "10:15 m -0600 MST",
    "10:15 p -0600 MST",
    "11:19  P.M.",
    "11:19  PM",
    "11:19  PM",
    "11:19  P  M",
    "11:19  P. M.",
    "11:19  p  m",
    "11:19  p. m.",
  }

  lArray := len(samples)
  for i := 0; i < lArray; i++ {
    match, err := mt.findExpressionExample01(samples[i], regexAMpm)

    if err != nil {
      if err.Error() == "No Match" {
        fmt.Printf("No Match - testStr == %v  regex == %v\n", samples[i], regexAMpm)
        continue
      } else {
        panic(err)
      }
    }

    fmt.Printf("Match! - testStr == %v  regex == %v  match string: %v \n", samples[i], regexAMpm, match)
  }

}

// FindExpression_Example_01 - Example function.
func (mt MainTest) FindExpressionExample01(targetStr string, regex string) (string, error) {

  if len(targetStr) < 1 {
    return "", fmt.Errorf("ExampleFindExpression_01() Invalid Target String: %v", targetStr)
  }

  // \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
  r, err := regexp.Compile(regex)

  if err != nil {
    return "", fmt.Errorf("Regex failed to Compile. regex== %v. Error: %v", regex, err.Error())
  }

  bTargetStr := []byte(targetStr)

  loc := r.FindIndex(bTargetStr)

  if loc == nil {
    return "", errors.New("No Match")
  }

  return string(bTargetStr[loc[0]:loc[1]]), nil

}

func (mt MainTest) TrimMultipleStringsExample01(tStr string, trimChar rune) {

  su := strops.StrOps{}

  r, err := su.TrimMultipleChars(tStr, trimChar)

  if err != nil {
    fmt.Println("Error Return from TrimMultipleChars: ", err.Error())
    return
  }

  fmt.Println("Original String: ", tStr)
  fmt.Println(" Trimmed String: ", r)
  fmt.Println("Original String Length: ", len(tStr))
  fmt.Println(" Trimmed String Length: ", len(r))
  tStr2 := strings.Replace(tStr, " ", "!", -1)
  fmt.Println("Original String TrimChar Locations: ", tStr2)
  r2 := strings.Replace(r, " ", "!", -1)
  fmt.Println(" Trimmed String TrimChar Locations: ", r2)

}

// RegExFindSingleTimeDigits_Example01
func (mt MainTest) RegExFindSingleTimeDigitsExample01() {
  regex := "\\d:\\d:\\d"
  targetStr := "November 12, 2016 1:6:3pm +0000 UTC"

  fmt.Println("targetStr = ", targetStr)
  su := strops.StrOps{}

  idx := su.FindRegExIndex(targetStr, regex)

  if idx == nil {
    panic(fmt.Errorf("Did not locate Regular Expression,'%v', in 'targetStr', '%v'.", regex, targetStr))
  }

  fmt.Println("Success - Found Regular Expression in targetStr!")
  fmt.Println("idx = ", idx)

  s := []byte(targetStr)

  extract := s[idx[0]:idx[1]]

  sExtract := string(extract)

  fmt.Println("Extracted String: ", sExtract)

  result := strings.Split(sExtract, ":")

  if len(result) == 0 {
    panic(fmt.Errorf("Split returned array of zero length"))
  }

  fmt.Println("Printing result array:")
  for j := 0; j < len(result); j++ {
    fmt.Println(result[j])
  }

  hrs, _ := strconv.Atoi(result[0])
  min, _ := strconv.Atoi(result[1])
  sec, _ := strconv.Atoi(result[2])

  fmt.Println("Printing Formatted Time String")
  fmt.Printf("%02d:%02d:%02d\n", hrs, min, sec)

  fmt.Println("Reprint with 2-digit seconds")
  fmt.Printf("%02d:%02d:%02d\n", hrs, min, 14)

}

func (mt MainTest) PrintFmtExample01() {

  s1 := fmt.Sprintf("No1: %d  No2: %d", 1, 2)

  fmt.Println(s1)
}


func (mt MainTest) ExampleWrite03() {

  fmt.Println("ExampleWrite03() - Version 2")

  originalStr := "Original base string written to sops1"

  sops1 := strops.StrOps{}.NewPtr()

  sops1.SetStringData(originalStr)

  sops2 := strops.StrOps{}.NewPtr()

  n, err := io.Copy(sops2, sops1)

  if err != nil {
    fmt.Printf("Error returned by io.Copy(sops2, sops1). Error='%v' \n", err.Error())
    return
  }

  actualStr := sops2.GetStringData()

  if originalStr != actualStr {
    fmt.Printf("Error: Expected string='%v'. Instead, string='%v'. \n",
      originalStr, actualStr)
  }

  fmt.Println()
  fmt.Println("----------------------------------------------------------------")
  fmt.Println("                       ExampleWrite03()")
  fmt.Println("----------------------------------------------------------------")
  fmt.Println("            Original String: ", originalStr)
  fmt.Println("               sops2 String: ", actualStr)
  fmt.Println("  Length of Original String: ", len(originalStr))
  fmt.Println("     Length of sops2 String: ", len(actualStr))
  fmt.Println("                    n Value: ", n)
  fmt.Println("sops1 Bytes Written Counter: ", sops1.GetCountBytesWritten())

}

func (mt MainTest) ExampleWrite01() {

  fmt.Println("ExampleWrite01() - Version 2")

  originalStr := "Hello World"

  sops1 := strops.StrOps{}.NewPtr()

  lenOriginalStr := len(originalStr)

  nArray := [4]int{}

  bytesWrittenArray := [4]uint64{}

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

    n, err := sops1.Write(p)

    if err != nil {
      fmt.Printf("Error returned by sops1.Write(p). Error='%v' ", err.Error())
      return
    }

    nArray[i] = n
    bytesWrittenArray[i] = sops1.GetCountBytesWritten()
  }

  actualStr := sops1.GetStringData()

  fmt.Println("=========================================================")
  fmt.Println("                  ExampleWrite01()")
  fmt.Println("=========================================================")
  fmt.Println("       Original String: ", originalStr)
  fmt.Println("Original String Length: ", lenOriginalStr)
  fmt.Println("         Actual String: ", actualStr)
  fmt.Println("  Actual String Length: ", len(actualStr))
  fmt.Println("               N Array: ", nArray)
  fmt.Println("   Bytes Written Array: ", bytesWrittenArray)

}

func (mt MainTest) ExampleWrite02() {

  fmt.Println("ExampleWrite02() - Version 2")

  originalStr := "Original base string written to sops1"

  sops1 := strops.StrOps{}.NewPtr()

  lenOriginalStr := len(originalStr)

  p := []byte(originalStr)

  n, err := sops1.Write(p)

  if err != nil {
    fmt.Printf("Error returned by sops1.Write(p). Error='%v' \n", err.Error())
    return
  }

  actualStr := sops1.GetStringData()

  if originalStr != actualStr {
    fmt.Printf("Error: Expected string='%v'. Instead, string='%v'. \n",
      originalStr, actualStr)
  }

  if lenOriginalStr != n {
    fmt.Printf("Error: Expected Length='%v'. Instead, Bytes Written='%v'. \n",
      lenOriginalStr, n)
  }

  fmt.Println()
  fmt.Println("----------------------------------------------------------------")
  fmt.Println("                       ExampleWrite02()")
  fmt.Println("----------------------------------------------------------------")
  fmt.Println("         Original String: ", originalStr)
  fmt.Println("    Actual Output String: ", actualStr)
  fmt.Println("           Length of 'p': ", len(p))
  fmt.Println("                 n Value: ", n)
  fmt.Println("   Bytes Written Counter: ", sops1.GetCountBytesWritten())
}

func (mt MainTest) ExampleRead01() {

  fmt.Println("ExampleRead01() - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strops.StrOps{}.NewPtr()
  sops1.SetStringData(originalStr)

  p := make([]byte, 5, 15)

  n := 0
  var err error
  err = nil
  cntr := uint64(0)

  b := strings.Builder{}
  b.Grow(len(originalStr) + 150)
  cntrArray := make([]uint64, 0, 50)

  for err != io.EOF {

    n, err = sops1.Read(p)

    if err != nil && err != io.EOF {
      fmt.Printf("Error returned by sops1.Read(p). "+
        "Error='%v' \n", err.Error())
      return
    }

    b.Write(p[:n])

    for i := 0; i < len(p); i++ {
      p[i] = byte(0)
    }

    cntrArray = append(cntrArray, sops1.GetCountBytesRead())

    cntr++

  }

  strBuilderStr := b.String()

  fmt.Println("         Original Str: ", originalStr)
  fmt.Println("  Original Str Length: ", len(originalStr))
  fmt.Println("         sops1.StrOut: ", sops1.GetStringData())
  fmt.Println("  sops1.StrOut Length: ", len(sops1.GetStringData()))
  fmt.Println("     sops1 Bytes Read: ", sops1.GetCountBytesRead())
  fmt.Println("              Counter: ", cntr)
  fmt.Println("Counter History Array: ", cntrArray)
  fmt.Println("       String Builder: ", strBuilderStr)
  fmt.Println("String Builder Length: ", len(strBuilderStr))
  fmt.Println("                    n: ", n)
  fmt.Println("                    p: ", p)
  fmt.Println("**********************************************")

}

func (mt MainTest) ExampleRead02() {
  fmt.Println("ExampleRead02 - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strops.StrOps{}.NewPtr()
  sops1.SetStringData(originalStr)

  p := make([]byte, 3, 100)

  _, err := sops1.Read(p)

  if err != nil && err != io.EOF {
    fmt.Printf("Error returned by sops1.Read(p). "+
      "Error='%v' \n", err.Error())
    return
  }

  sops2 := strops.StrOps{}.NewPtr()
  n, err := sops2.Write(p)

  fmt.Println("        Original Str: ", originalStr)
  fmt.Println(" Original Str Length: ", len(originalStr))
  fmt.Println("        sops1.StrOut: ", sops1.GetStringData())
  fmt.Println(" sops1.StrOut Length: ", len(sops1.GetStringData()))
  fmt.Println("    sops1 Bytes Read: ", n)
  fmt.Println("**********************************************")

}

func (mt MainTest) ExampleIoCopy02() {
  fmt.Println("ExampleIOCopy_02() - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strops.StrOps{}.NewPtr()
  sops1.SetStringData(originalStr)
  sops2 := strops.StrOps{}.NewPtr()

  n, err := io.Copy(sops2, sops1)

  if err != nil {
    fmt.Printf("Error returned by io.Copy(sops2, sops1). "+
      "Error='%v' \n", err.Error())
    return
  }

  fmt.Println("           Original Str: ", originalStr)
  fmt.Println("    Original Str Length: ", len(originalStr))
  fmt.Println("       sops1.stringData: ", sops1.GetStringData())
  fmt.Println("       sops2.stringData: ", sops2.GetStringData())
  fmt.Println("sops2.stringData Length: ", len(sops2.GetStringData()))
  fmt.Println("          Bytes Written: ", n)
  fmt.Println("**********************************************")
  fmt.Println("      Copying sops2 To StringData")
  fmt.Println("**********************************************")

  n, err = io.Copy(os.Stdout, sops2)

  if err != nil {
    fmt.Printf("Error returned by io.Copy(os.Stdout, sops2). "+
      "Error='%v' \n", err.Error())
    return
  }

  fmt.Println()
  fmt.Println("New value of n: ", n)
}

func (mt MainTest) timer(starTime, endTime time.Time) string {

  // MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
  // 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
  MicroSecondNanoseconds := int64(time.Microsecond)

  // MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
  //	 A millisecond is 1/1,000 or 1 one-thousandth of a second
  MilliSecondNanoseconds := int64(time.Millisecond)

  // SecondNanoseconds - Number of Nanoseconds in a Second
  SecondNanoseconds := int64(time.Second)

  // MinuteNanoseconds - Number of Nanoseconds in a minute
  MinuteNanoseconds := int64(time.Minute)

  // HourNanoseconds - Number of Nanoseconds in an hour
  HourNanoseconds := int64(time.Hour)

  t2Dur := endTime.Sub(starTime)

  str := ""

  totalNanoseconds := t2Dur.Nanoseconds()
  var numOfHours, numOfMinutes, numOfSeconds, numOfMillisecionds,
  numOfMicroseconds, numOfNanoseconds int64

  if totalNanoseconds >= HourNanoseconds {
    numOfHours = totalNanoseconds / HourNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfHours * HourNanoseconds)
  }

  if totalNanoseconds >= MinuteNanoseconds {
    numOfMinutes = totalNanoseconds / MinuteNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMinutes * MinuteNanoseconds)
  }

  if totalNanoseconds >= SecondNanoseconds {
    numOfSeconds = totalNanoseconds / SecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
  }

  if totalNanoseconds >= SecondNanoseconds {
    numOfSeconds = totalNanoseconds / SecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
  }

  if totalNanoseconds >= MilliSecondNanoseconds {
    numOfMillisecionds = totalNanoseconds / MilliSecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMillisecionds * MilliSecondNanoseconds)
  }

  if totalNanoseconds >= MicroSecondNanoseconds {
    numOfMicroseconds = totalNanoseconds / MicroSecondNanoseconds
    totalNanoseconds = totalNanoseconds - (numOfMicroseconds * MicroSecondNanoseconds)
  }

  numOfNanoseconds = totalNanoseconds

  if numOfHours > 0 {

    str += fmt.Sprintf("%v-Hours ", numOfHours)

  }

  if numOfMinutes > 0 {

    str += fmt.Sprintf("%v-Minutes ", numOfMinutes)

  }

  if numOfSeconds > 0 || str != "" {

    str += fmt.Sprintf("%v-Seconds ", numOfSeconds)

  }

  if numOfMillisecionds > 0 || str != "" {

    str += fmt.Sprintf("%v-Milliseconds ", numOfMillisecionds)

  }

  if numOfMicroseconds > 0 || str != "" {

    str += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

  }

  str += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

  return str
}

