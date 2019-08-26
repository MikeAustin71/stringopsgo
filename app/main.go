package main

// This Works
// strOps "./strops/v2"
// Sometimes GoLand will declare this invalid. If so, just
// turn off warning.
//
import (
 "../strops/v2"
  "fmt"
  "strings"
)

func main() {

  mainTest{}.ExampleExtractNumStr01()
}

type mainTest struct {
  input string
}

func (mt mainTest) ExampleExtractNumStr01() {

  targetStr := "-123.5"

  expectedNumStr := "-123.5"
  expectedLeadingSignChar := "-"
  expectedNumStrLen := len(expectedNumStr)
  expectedNumIdx := strings.Index(targetStr, expectedNumStr)
  startIdx := 0

  numIndex,
  numLen,
  leadingSignChar,
  numStr,
  err := strops.StrOps{}.FindNumericDigitsString(targetStr, startIdx,".","")

  if err != nil {
    fmt.Printf("Error returned by StrOps{}.FindNumericDigitsString(targetStr, 0)\n" +
      "targetStr='%v'\nError='%v'\n", targetStr, err.Error())
    return
  }

  isError := false

  if expectedNumIdx != numIndex {
    fmt.Printf("Expected starting numeric index='%v'\n" +
      "Instead, staring numeric index='%v'\n",
      expectedNumIdx, numIndex)
    isError = true
  }

  if expectedNumStr != numStr {
    fmt.Printf("Expected number string ='%v'\n" +
      "Instead, number string ='%v'\n",
      expectedNumStr, numStr)
    isError = true
  }

  if expectedNumStrLen != numLen {
    fmt.Printf("Expected number string length ='%v'\n" +
      "Instead, number string length ='%v'\n",
      expectedNumStrLen, numLen)
    isError = true
  }

  if expectedLeadingSignChar != leadingSignChar {
    fmt.Printf("Expected leading sign char ='%v'\n" +
      "Instead, leading sign char ='%v'\n",
      expectedLeadingSignChar, leadingSignChar)
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
  fmt.Println("           startIdx: ", startIdx)
  fmt.Println("-------------------------------------")
  fmt.Println("           Expected                  ")
  fmt.Println("-------------------------------------")
  fmt.Println("       Number Index: ", expectedNumIdx)
  fmt.Println("         Num Length: ", expectedNumStrLen)
  fmt.Println("  Leading Sign Char: ", expectedLeadingSignChar)
  fmt.Println("      Number String: ", expectedNumStr)
  fmt.Println("-------------------------------------")
  fmt.Println("            Results                  ")
  fmt.Println("-------------------------------------")
  fmt.Println("        NumberIndex: ", numIndex)
  fmt.Println("         Num Length: ", numLen)
  fmt.Println("  Leading Sign Char: ", leadingSignChar)
  fmt.Println("      Number String: ", numStr)
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