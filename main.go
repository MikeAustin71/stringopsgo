package main

// This Works
// strOps "./strops/v2"
// Sometimes GoLand will declare this invalid. If so, just
// turn off warning.
//
import (
  strOps "./strops/v2"
  "fmt"
  "io"
  "os"
  "sort"
  "strings"
  "time"
)

func main() {

  mainTest{}.ExampleStripBadChars01()

}

type mainTest struct {
  input string
}

func (mt mainTest) ExampleStripBadChars01() {
  badChars := []string {
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
  testString :=  "..........      ./../.\\.\\..\\////   " + expectedStr +
    "..........      ./../.\\.\\..\\////   "

  var startTime time.Time
  var endTime time.Time

  startTime = time.Now()
  actualString, actualStrLen := strOps.StrOps{}.StripBadChars(testString, badChars)
  endTime = time.Now()

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
    return
  }

  elapsedTime := mt.timer(startTime, endTime)

  fmt.Println("mainTest.ExampleStripBadChars01()")
  fmt.Println("-------------------------------------")
  fmt.Println("          SUCCESS!!!")
  fmt.Println("-------------------------------------")
  fmt.Println("            Test String: ", testString)
  fmt.Println("     Test String Length: ", len(testString))
  fmt.Println("        Expected String: ", expectedStr)
  fmt.Println(" Expected String Length: ", expectedStrLen)
  fmt.Println("           Clean String: ", actualString)
  fmt.Println("    Clean String Length: ", actualStrLen)
  fmt.Println("Actual Clean Str Length: ", len(actualString))
  fmt.Println("           Elapsed Time: ", elapsedTime)

}

func (mt mainTest) ExampleStripTrailingChars01() {

  badChars := []string {
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

  actualString, actualStrLen := strOps.StrOps{}.StripTrailingChars(testString, badChars)

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
    return
  }

  fmt.Println("mainTest.ExampleStripTrailingChars01()")
  fmt.Println("-------------------------------------")
  fmt.Println("          SUCCESS!!!")
  fmt.Println("-------------------------------------")
  fmt.Println("            Test String: ", testString)
  fmt.Println("     Test String Length: ", len(testString))
  fmt.Println("        Expected String: ", expectedStr)
  fmt.Println(" Expected String Length: ", expectedStrLen)
  fmt.Println("           Clean String: ", actualString)
  fmt.Println("    Clean String Length: ", actualStrLen)
  fmt.Println("Actual Clean Str Length: ", len(actualString))
}

func (mt mainTest) ExampleStripLeadingChars01() {

  badChars := []string {
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

  actualString, actualStrLen := strOps.StrOps{}.StripLeadingChars(testString, badChars)

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
    return
  }

  fmt.Println("mainTest.ExampleStripLeadingChars01()")
  fmt.Println("-------------------------------------")
  fmt.Println("          SUCCESS!!!")
  fmt.Println("-------------------------------------")
  fmt.Println("            Test String: ", testString)
  fmt.Println("     Test String Length: ", len(testString))
  fmt.Println("        Expected String: ", expectedStr)
  fmt.Println(" Expected String Length: ", expectedStrLen)
  fmt.Println("           Clean String: ", actualString)
  fmt.Println("    Clean String Length: ", actualStrLen)
  fmt.Println("Actual Clean Str Length: ", len(actualString))

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


func (mt mainTest) timer(starTime, endTime time.Time) string {

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
  numOfHours := int64(0)
  numOfMinutes := int64(0)
  numOfSeconds := int64(0)
  numOfMillisecionds := int64(0)
  numOfMicroseconds := int64(0)
  numOfNanoseconds := int64(0)

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
