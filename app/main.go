package main

/* This Works
import (
  "MikeAustin71/stringopsgo/strops/v2"
  "fmt"
  "strings"

)

Reference format:
 strops.StrOps{}.ExtractNumericDigits(..)

*/

import (
  "MikeAustin71/stringopsgo/strops/v2"
  "fmt"
  "strings"

)

func main() {

  mainTest{}.ExampleExtractDataField01()
}

type mainTest struct {
  input string
}

func (mt mainTest) ExampleExtractDataField01() {

  endOfLineRunes := []rune("\n#")
  leadingRunes := []rune("\t \r\f\n\v")
  trailingRunes := []rune ("\t \r\f\n\v")
  targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
  lenTargetStr := len(targetStr)
  startIdx := 0
  leadingKeyWordDelimiter := "Zone:"
  expectedFieldStr := "America/Chicago"
  expectedFieldIdx := strings.Index(targetStr, expectedFieldStr)
  expectedFieldLength := len(expectedFieldStr)
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
    fmt.Printf("Error returned by strops.StrOps{}.ExtractDataField()\n" +
      "targetStr='%v'\tstartIdx='%v'\n" +
      "Error='%v'\n", targetStr, startIdx, err.Error() )
    return
  }

  isError := false

  if datDto.TargetStr != targetStr {
    fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n" +
      "Instead, datDto.TargetStr='%v'.\n",
      datDto.TargetStr, targetStr)
    isError = true
  }

  if datDto.TargetStrLength != lenTargetStr {
    fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n" +
      "Instead, datDto.TargetStrLength='%v'.\n",
      datDto.TargetStrLength, lenTargetStr)
    isError = true
  }

  if datDto.StartIndex != startIdx {
    fmt.Printf("ERROR: Expected datDto.StartIndex='%v'.\n" +
      "Instead, datDto.StartIndex='%v'.\n",
      datDto.StartIndex, startIdx)
    isError = true
  }

  if leadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
    fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n" +
      "Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
      leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
    isError = true
  }

  if datDto.DataFieldStr != expectedFieldStr {
    fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n" +
      "Instead, datDto.DataFieldStr='%v'.\n",
      datDto.DataFieldStr, expectedFieldStr)
    isError = true
  }

  if datDto.DataFieldLength != expectedFieldLength {
    fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n" +
      "Instead, datDto.DataFieldLength='%v'.\n",
      datDto.DataFieldLength, expectedFieldLength)
    isError = true
  }

  if datDto.DataFieldIndex != expectedFieldIdx {
    fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n" +
      "Instead, datDto.DataFieldIndex='%v'.\n",
      datDto.DataFieldIndex, expectedFieldIdx)
    isError = true
  }

  if datDto.NextTargetStrIndex != expectedNextTargetIdx {
    fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n" +
      "Instead, datDto.NextTargetStrIndex='%v'.\n",
      datDto.NextTargetStrIndex, expectedNextTargetIdx)
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
  fmt.Printf("         TargetStr: %v", targetStr)
  fmt.Println("  TargetStr Length: ", lenTargetStr)
  fmt.Println("       Start Index: ", startIdx)
  fmt.Println("------------------------------------------------")
  fmt.Println("                 Expected Results               ")
  fmt.Println("------------------------------------------------")
  fmt.Println("      Field String: ", expectedFieldStr)
  fmt.Println("  Field Str Length: ", expectedFieldLength)
  fmt.Println("       Field Index: ", expectedFieldIdx)
  fmt.Println(" Next Target Index: ", expectedNextTargetIdx)
  fmt.Println("------------------------------------------------")
  fmt.Println("                  Actual Results                ")
  fmt.Println("------------------------------------------------")
  fmt.Println("      Field String: ", datDto.DataFieldStr)
  fmt.Println("  Field Str Length: ", datDto.DataFieldLength)
  fmt.Println("       Field Index: ", datDto.DataFieldIndex)
  fmt.Println(" Next Target Index: ", datDto.NextTargetStrIndex)
  fmt.Println("     Target String: ", datDto.TargetStr)
  fmt.Println(" Target Str Length: ", datDto.TargetStrLength)
  fmt.Println("   Target StartIdx: ", datDto.StartIndex)
  fmt.Println("  Leading Key Word: ", datDto.LeadingKeyWordDelimiter)

}

func (mt mainTest) ExampleExtractNumStr01() {
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
  err :=  strops.StrOps{}.ExtractNumericDigits(
    targetStr,
    startIndex,
    keepLeadingChars,
    keepInteriorChars,
    keepTrailingChars)

  if err != nil {
    fmt.Printf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  isError := false

  if expectedNumIdx != nStrDto.FirstNumCharIndex {
    fmt.Printf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, nStrDto.FirstNumCharIndex)
    isError = true
  }

  if expectedNumStr != nStrDto.NumStr {
    fmt.Printf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, nStrDto.NumStr)
    isError = true
  }

  if expectedNumStrLen != nStrDto.NumStrLen {
    fmt.Printf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, nStrDto.NumStrLen)
    isError = true
  }

  if expectedLeadingSignChar != nStrDto.LeadingSignChar {
    fmt.Printf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, nStrDto.LeadingSignChar)
    isError = true
  }

  if expectedLeadingSignIndex != nStrDto.LeadingSignIndex {
    fmt.Printf("Expected leading sign index ='%v'\n" +
      "Instead, leading sign index ='%v'\n",
      expectedLeadingSignIndex, nStrDto.LeadingSignIndex)
    isError = true
  }

  if expectedNextTargetStrIdx != nStrDto.NextTargetStrIndex {
    fmt.Printf("Expected Next TargetStr Char Index ='%v'\n" +
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

/*
func (mt mainTest) ExampleStripLeadingChars01() {

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

  actualString, actualStrLen := StrOps{}.StripLeadingChars(testString, badChars)

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

func (mt mainTest) ExampleSortStrLenHighestToLowest01() {
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

  sort.Sort(strOps.SortStrLengthHighestToLowest(badChars))

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

func (mt mainTest) ExampleSortStrLenLowestToHighest01() {

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

  sort.Sort(strOps.SortStrLengthLowestToHighest(badChars))

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

func (mt mainTest) ExampleWrite03() {

  fmt.Println("ExampleWrite03() - Version 2")

  originalStr := "Original base string written to sops1"

  sops1 := strOps.StrOps{}.NewPtr()

  sops1.SetStringData(originalStr)

  sops2 := strOps.StrOps{}.NewPtr()

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

func (mt mainTest) ExampleWrite01() {

  fmt.Println("ExampleWrite01() - Version 2")

  originalStr := "Hello World"

  sops1 := strOps.StrOps{}.NewPtr()

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

func (mt mainTest) ExampleWrite02() {

  fmt.Println("ExampleWrite02() - Version 2")

  originalStr := "Original base string written to sops1"

  sops1 := strOps.StrOps{}.NewPtr()

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

func (mt mainTest) ExampleRead01() {

  fmt.Println("ExampleRead01() - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strOps.StrOps{}.NewPtr()
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

func (mt mainTest) ExampleRead02() {
  fmt.Println("ExampleRead02 - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strOps.StrOps{}.NewPtr()
  sops1.SetStringData(originalStr)

  p := make([]byte, 3, 100)

  n, err := sops1.Read(p)

  if err != nil && err != io.EOF {
    fmt.Printf("Error returned by sops1.Read(p). "+
      "Error='%v' \n", err.Error())
    return
  }

  sops2 := strOps.StrOps{}.NewPtr()
  n, err = sops2.Write(p)

  fmt.Println("        Original Str: ", originalStr)
  fmt.Println(" Original Str Length: ", len(originalStr))
  fmt.Println("        sops1.StrOut: ", sops1.GetStringData())
  fmt.Println(" sops1.StrOut Length: ", len(sops1.GetStringData()))
  fmt.Println("    sops1 Bytes Read: ", n)
  fmt.Println("**********************************************")

}

func (mt mainTest) ExampleIoCopy02() {
  fmt.Println("ExampleIOCopy_02() - Version 2")

  originalStr := "Original sops1 base string"

  sops1 := strOps.StrOps{}.NewPtr()
  sops1.SetStringData(originalStr)
  sops2 := strOps.StrOps{}.NewPtr()

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

*/