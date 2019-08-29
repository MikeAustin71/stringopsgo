/*
package strops (string operations) provides string management
utilities designed to perform a variety of string operations
including string centering, justification, multiple replacements
and implementation of the the io.Reader and io.Writer interfaces.

Source file, 'xStrops.go', is located in source code repository:
	https://github.com/MikeAustin71/stringopsgo.git

Currently, the package consists of one type, 'StrOps' which is
documented below.
*/
package strops

import (
  "errors"
  "fmt"
  "io"
  "regexp"
  "sort"
  "strings"
  "sync"
  "unicode/utf8"
)

// SortStrLengthHighestToLowest - Uses to perform two level sort
// on string arrays. The strings are first sorted by string length
// (greatest length to Lowest length) and then by alphabetic sort.
//
// This type is designed to be used in conjunction with 'sort.Sort()'
// Reference the Go Sort Package:
//      https://golang.org/pkg/sort/#Sort
//
// Example Usage:
//   badChars := []string {
//    "aaaaa",
//    "bbbbb",
//    "cccccccccc",
//    "z",
//    "fffffffffff",
//    "xx",
//    "ddddddddd",
//    "eeeeeeeeeee" }
//
//     SortStrLengthLowestToHighest(badChars)
//
//     Output:
//
//        ================================
//        Sort by Length Highest To Lowest
//        Ordered List
//        ================================
//
//        1. fffffffffff
//        2. eeeeeeeeeee
//        3. cccccccccc
//        4. ddddddddd
//        5. bbbbb
//        6. aaaaa
//        7. xx
//        8. z
//
type SortStrLengthHighestToLowest []string

// Len - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Len() int {
  return len(sortStrLenHigh)
}

// Swap - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Swap(i, j int) {
  sortStrLenHigh[i], sortStrLenHigh[j] = sortStrLenHigh[j], sortStrLenHigh[i]
}

// Less - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenHigh SortStrLengthHighestToLowest) Less(i, j int) bool {

  lenI := len(sortStrLenHigh[i])
  lenJ := len(sortStrLenHigh[j])
  if lenI == lenJ {
    return sortStrLenHigh[i] > sortStrLenHigh[j]
  }

  return lenI > lenJ
}

// SortStrLengthLowestToHighest - Uses to perform two level sort
// on string arrays. The strings are first sorted by string length
// (smallest length to greatest length) and then by alphabetic sort.
//
// This type is designed to be used in conjunction with 'sort.Sort()'
// Reference the Go Sort Package:
//      https://golang.org/pkg/sort/#Sort
//
// Example Usage:
//      badChars := []string {
//    "aaaaa",
//    "bbbbb",
//    "cccccccccc",
//    "z",
//    "fffffffffff",
//    "xx",
//    "ddddddddd",
//    "eeeeeeeeeee" }
//
//     SortStrLengthLowestToHighest(badChars)
//
//     Output:
//
//       ================================
//       Sort by Length Lowest To Highest
//       Ordered List
//       ================================
//       1. z
//       2. xx
//       3. aaaaa
//       4. bbbbb
//       5. ddddddddd
//       6. cccccccccc
//       7. eeeeeeeeeee
//       8. fffffffffff
//
type SortStrLengthLowestToHighest []string

// Len - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Len() int {
  return len(sortStrLenLow)
}

// Swap - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Swap(i, j int) {
  sortStrLenLow[i], sortStrLenLow[j] = sortStrLenLow[j], sortStrLenLow[i]
}

// Less - This is part of the sort.Interface. Reference the 'sort' package:
//   https://golang.org/pkg/sort/#Interface
//   https://golang.org/pkg/sort/#Sort
//
func (sortStrLenLow SortStrLengthLowestToHighest) Less(i, j int) bool {

  lenI := len(sortStrLenLow[i])
  lenJ := len(sortStrLenLow[j])
  if lenI == lenJ {
    return sortStrLenLow[i] < sortStrLenLow[j]
  }

  return lenI < lenJ
}

// NumStrProfileDto - This type is used to encapsulate information
// on strings of numeric digits which are extracted from larger
// strings.
//
type NumStrProfileDto struct {
  TargetStr  string // The original target string which is scanned for a number string
  StartIndex int    // The starting index in 'TargetStr' from which the number string
  //    search was initiated.
  LeadingSignIndex int //  The string index of a leading sign in 'NumStr' below. If a
  //    leading sign character is NOT present in 'NumStr' this value
  //    is set to -1
  LeadingSignChar string //  If a leading sign character (plus '+' or minus '-') exists in
  //    data field 'NumStr' (below), it is stored in this string.
  FirstNumCharIndex int //  The index in 'TargetStr' (above) where the first character
  //    of the extracted number string is located.
  NextTargetStrIndex int //  The index of the next character in 'TargetStr' immediately
  //    following the extracted number string.
  NumStrLen int    //  The length of the extracted number string.
  NumStr    string //  The number string extracted from 'TargetStr'.
}

// New - Creates and returns a new instance of NumStrProfileDto
// which is properly initialized.
func (exNumDto NumStrProfileDto) New() NumStrProfileDto {

  newDto := NumStrProfileDto{}
  newDto.TargetStr = ""
  newDto.StartIndex = -1
  newDto.LeadingSignIndex = -1
  newDto.LeadingSignChar = ""
  newDto.FirstNumCharIndex = -1
  newDto.NextTargetStrIndex = -1
  newDto.NumStrLen = 0
  newDto.NumStr = ""
  return newDto
}

// DataFieldProfileDto - This type is used to encapsulate information
// related to an extracted data field string.
//
type DataFieldProfileDto struct {
  TargetStr               string    //  The string from which the data field
                                    //    is extracted.
  TargetStrLength         int       //  Length of 'TargetStr'
  StartIndex              int       //  The index with in 'TargetStr' from which
                                    //    the search for a data field was initiated.
  LeadingKeyWordDelimiter string    //  The Leading Key Word Delimiter which is used
                                    //    identify the beginning of the field search
  DataFieldStr            string    //  The extracted data field string
  DataFieldIndex          int       //  The index in 'TargetStr' where the data field
                                    //    begins.
  DataFieldLength         int       //  The length of the extracted data field string.
  NextTargetStrIndex      int       //  The index in 'TargetStr' immediately following
                                    //    the extracted data field.
}

func (dfProfile DataFieldProfileDto) New() DataFieldProfileDto {
  newDataDto := DataFieldProfileDto{}
  newDataDto.TargetStr = ""
  newDataDto.StartIndex = -1
  newDataDto.LeadingKeyWordDelimiter = ""
  newDataDto.DataFieldStr = ""
  newDataDto.DataFieldIndex = -1
  newDataDto.DataFieldLength = 0
  newDataDto.NextTargetStrIndex = -1
  return newDataDto
}

// StrOps - encapsulates a collection of methods used to manage string
// operations.
//
// Most of the utility offered by this type is provided through its
// associated methods. However, given that two data elements, 'StrIn'
// and 'StrOut' are provided, the structure may be used as a data
// transport object (dto) containing two strings.
//
// Version 2.0.0 and all later versions of this type support Go Modules.
// For Version 2+ implementations, use the following import statement:
//
//		import "github.com/MikeAustin71/stringopsgo/strops/v2"
//
// Earlier Version 1.0.0 implementations require a different import statement:
//
//    import "github.com/MikeAustin71/stringopsgo/strops"
//
// Be advised that this type, 'StrOps', implements the io.Reader and io.Writer
// interfaces. All io.Reader and io.Writer operations utilize the private string
// data element, 'StrOps.stringData'.
type StrOps struct {
  StrIn      string // public string variable available at user's discretion
  StrOut     string // public string variable available at user's discretion
  stringData string // private string variable accessed by StrOps.Read and
  //	StrOps.Write. Accessed through methods
  //	StrOps.GetStringData() and StrOps.SetStringData()
  stringDataMutex sync.Mutex // Used internally to ensure thread safe operations
  cntBytesRead    uint64     // Used internally to track Bytes Read by StrOps.Read()
  cntBytesWritten uint64     // Used internally to track Bytes Written by StrOps.Write()
}

// BreakTextAtLineLength - Breaks string text into lines. Takes a string and inserts a
// line delimiter character (a.k.a 'rune') at the specified line length ('lineLength').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//	targetStr	string	- The string which will be parsed into text lines.
//
//	lineLength	int	- The maximum length of each line.
//
// Note: If the caller specifies a line length of 50, the line delimiter character may be placed in the
// 51st character position depending upon the word breaks.
func (sops StrOps) BreakTextAtLineLength(targetStr string, lineLength int, lineDelimiter rune) (string, error) {

  ePrefix := "StrOps.BreakTextAtLineLength() "

  targetLen := len(targetStr)

  if targetLen == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'targetStr' is a ZERO LENGTH STRING!")
  }

  if lineLength < 5 {
    return "",
      fmt.Errorf(ePrefix+"Error: Input parameter 'lineLength' is LESS THAN 5-CHARACTERS! "+
        "lineLength='%v' ", lineLength)
  }

  if lineDelimiter == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'lineDelimiter' is ZERO VALUE!")
  }

  if sops.IsEmptyOrWhiteSpace(targetStr) {
    return "\n", nil
  }

  var err error

  var b strings.Builder
  b.Grow(((targetLen / lineLength) * targetLen) + 50)

  begIdx := 0
  endWrdIdx := 0
  isAllOneWord := false
  isAllSpaces := false
  actualLastIdx := 0
  beginWrdIdx := 0

  for begIdx < targetLen && begIdx > -1 {

    // skip spaces at the beginning of the line
    begIdx, err = sops.FindFirstNonSpaceChar(targetStr, begIdx, targetLen-1)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+
          "Error returned by sops.FindFirstNonSpaceChar(targetStr, begIdx, actualLastIdx). "+
          "targetStr='%v' begIdx='%v' actualLastIdx='%v'   Error='%v' ",
          targetStr, begIdx, actualLastIdx, err.Error())
    }

    if begIdx == -1 {

      if b.Len() == 0 {
        b.WriteRune(lineDelimiter)
      }

      break // Exit loop
    }

    if begIdx == targetLen-1 {
      b.WriteByte(targetStr[begIdx])
      b.WriteRune(lineDelimiter)
      begIdx = -1 // Exit loop
      break
    }

    // begIdx < targetLen - 1

    actualLastIdx = begIdx + lineLength - 1

    if actualLastIdx >= targetLen {
      actualLastIdx = targetLen - 1
    }

    // Find the last complete word in this string segment
    beginWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err =
      sops.FindLastWord(targetStr, begIdx, actualLastIdx)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+
          "Error returned by sops.FindLastWord(targetStr,begIdx, actualLastIdx). "+
          "targetStr='%v' begIdx='%v' actualLastIdx='%v'   Error='%v' ",
          targetStr, begIdx, actualLastIdx, err.Error())
    }

    if isAllSpaces {
      // This string segment is all spaces
      // write a line delimiter and continue
      begIdx = actualLastIdx + 1

    } else if isAllOneWord {
      // This string segment is all one word
      // and contains NO spaces.

      if actualLastIdx+1 >= targetLen {
        // If this is end of the main string,
        // just write the remaining segment and
        // exit.
        //
        b.WriteString(targetStr[begIdx:])
        b.WriteRune(lineDelimiter)
        break

      } else if actualLastIdx-begIdx+1 <= lineLength {
        // If this string segment is less than the specified
        // line length, just write the entire line segment.
        // Be careful, we may be at the end of the main
        // string.

        if actualLastIdx+1 >= targetLen {
          // This is the end of the main string,
          // just exit.
          b.WriteString(targetStr[begIdx:])
          b.WriteRune(lineDelimiter)
          break

        } else {

          b.WriteString(targetStr[begIdx : actualLastIdx+1])
          begIdx = actualLastIdx + 1
        }

      } else {
        // Out of options. Nothing left to do but hyphenate
        // the word.
        b.WriteString(targetStr[begIdx : actualLastIdx-1])
        b.WriteRune('-')
        begIdx = actualLastIdx

      }

    } else {
      // The segment is NOT All spaces nor is it all one word.

      if endWrdIdx+1 >= targetLen {
        // Are we at the end of targetStr
        b.WriteString(targetStr[begIdx:])
        b.WriteRune(lineDelimiter)
        break

      } else if targetStr[endWrdIdx+1] != ' ' {
        // This word crosses a line break boundary. Try not to split the word.

        // Find  the end of the last word.
        idx, err := sops.FindLastNonSpaceChar(targetStr, begIdx, beginWrdIdx-1)

        if err != nil {
          return "",
            fmt.Errorf(ePrefix+
              "Error returned by sops.FindLastNonSpaceChar(targetStr,begIdx, beginWrdIdx-1). "+
              "targetStr='%v' begIdx='%v' actualLastIdx='%v'   Error='%v' ",
              targetStr, begIdx, actualLastIdx, err.Error())
        }

        if idx == -1 {
          begIdx = beginWrdIdx
          // Do not write end of line delimiter
          // Set bigIdx to beginning of word and
          // loop again
          continue

        } else {
          // Success we found the end of the last word.
          b.WriteString(targetStr[begIdx : idx+1])
          begIdx = idx + 1
        }

      } else {
        // The word does not cross a line break boundary.
        // The next character after the last word is a
        // space.

        b.WriteString(targetStr[begIdx : endWrdIdx+1])
        begIdx = endWrdIdx + 1
      }
    }

    b.WriteRune(lineDelimiter)

  }

  return b.String(), nil
}

// CopyIn - Copies string information from another StrOps
// instance passed as an input parameter to the current
// StrOps instance.
func (sops *StrOps) CopyIn(strops2 *StrOps) {

  sops.StrIn = strops2.StrIn
  sops.StrOut = strops2.StrOut
  sops.stringDataMutex.Lock()
  strops2.stringDataMutex.Lock()
  sops.cntBytesWritten = 0
  sops.cntBytesRead = 0
  sops.stringData = strops2.stringData
  strops2.stringDataMutex.Unlock()
  sops.stringDataMutex.Unlock()

}

// CopyOut - Creates a 'deep' copy of the current
// StrOps instance and returns a pointer to a
// new instance containing that copied information.
func (sops *StrOps) CopyOut() *StrOps {

  strops2 := StrOps{}
  strops2.StrIn = sops.StrIn
  strops2.StrOut = sops.StrOut
  sops.stringDataMutex.Lock()
  strops2.stringData = sops.stringData
  sops.stringDataMutex.Unlock()

  return &strops2
}

// DoesLastCharExist - returns true if the last character (rune) of
// input string 'testStr' is equal to input parameter 'lastChar' which
// is of type 'rune'.
func (sops StrOps) DoesLastCharExist(testStr string, lastChar rune) bool {

  testStrLen := len(testStr)

  if testStrLen == 0 {
    return false
  }

  strLastChar := rune(testStr[testStrLen-1])

  if strLastChar == lastChar {
    return true
  }

  return false
}

// ExtractDataField - Extracts a data field string from a larger target string ('TargetStr').
//
func (sops StrOps) ExtractDataField(
  targetStr string,
  leadingKeyWordDelimiter string,
  startIdx int,
  leadingFieldSeparators []rune,
  trailingFieldSeparators []rune,
  endOfStringDelimiters []rune) (DataFieldProfileDto, error) {

  ePrefix := "StrOps.ExtractDataField() "
  newDataDto := DataFieldProfileDto{}.New()

  lenTargetStr := len(targetStr)

  if lenTargetStr == 0 {

    return newDataDto,
      errors.New(ePrefix +
        "ERROR: Input Parameter 'targetStr' is an EMPTY string!\n")
  }

  if startIdx >= lenTargetStr {
    return newDataDto,
      fmt.Errorf("ERROR: Input Parameter 'startIdx' is out-of-bounds!\n" +
        "startIdx='%v'\t\tLast TargetStr Index='%v'\n" +
        "Length Of TargetStr='%v'\n",
        startIdx, lenTargetStr -1, lenTargetStr)
  }

  newDataDto.TargetStr = targetStr
  newDataDto.StartIndex = startIdx
  newDataDto.TargetStrLength = lenTargetStr
  newDataDto.LeadingKeyWordDelimiter = leadingKeyWordDelimiter

  targetStrRunes := []rune(targetStr)
  lenTargetStr = len(targetStrRunes)
  endTargetStrIdx := lenTargetStr - 1

  lenOfEndOfStringDelimiters := len(endOfStringDelimiters)

  if lenOfEndOfStringDelimiters > 0 {

    for a:=startIdx; a < lenTargetStr; a++ {

      for b:=0; b < lenOfEndOfStringDelimiters; b++ {
        if endOfStringDelimiters[b] == targetStrRunes[a] {
          endTargetStrIdx = a
          goto endOfStringDelimiter
        }
      }
    }
  }

 endOfStringDelimiter:

  if startIdx == endTargetStrIdx {
    return newDataDto, nil
  }

  lenLeadingKeyWordDelimiter := len(leadingKeyWordDelimiter)

  if lenLeadingKeyWordDelimiter > 0 {

    keyWordIdx := strings.Index(targetStr[startIdx:], leadingKeyWordDelimiter)

    if keyWordIdx== -1 ||
      keyWordIdx >= endTargetStrIdx {
      return newDataDto, nil
    }

    startIdx =  lenLeadingKeyWordDelimiter + keyWordIdx
  }

  lenTrailingFieldSepartors := len(trailingFieldSeparators)
  fieldDataRunes := make([]rune, 0, 20)
  firstDataFieldIdx := -1

  for i:= startIdx; i <= endTargetStrIdx; i++ {

    for m:=0; m < lenOfEndOfStringDelimiters; m++ {
      if targetStrRunes[i] == endOfStringDelimiters[m] {
        goto exitMainTargetLoop
      }
    }

    if firstDataFieldIdx == -1 {

      for j:= 0; j < len(leadingFieldSeparators); j++ {

        if targetStrRunes[i] == leadingFieldSeparators[j]  {
          goto mainTargetLoop
        }

      }
    }

    for k:=0; k < lenTrailingFieldSepartors; k++ {
      if targetStrRunes[i] == trailingFieldSeparators[k] {
        goto exitMainTargetLoop
      }
    }

    if firstDataFieldIdx == -1 {
      firstDataFieldIdx = i
    }

    fieldDataRunes = append(fieldDataRunes, targetStrRunes[i])

  mainTargetLoop:
  }

  exitMainTargetLoop:

  if len(fieldDataRunes) == 0 {
    return newDataDto, nil
  }

  newDataDto.DataFieldStr = string(fieldDataRunes)
  newDataDto.DataFieldLength = len(newDataDto.DataFieldStr)
  newDataDto.DataFieldIndex = firstDataFieldIdx
  nextIdx := newDataDto.DataFieldIndex + newDataDto.DataFieldLength

  if nextIdx >= endTargetStrIdx {
    newDataDto.NextTargetStrIndex = -1
  } else {
    newDataDto.NextTargetStrIndex = nextIdx
  }

  return newDataDto, nil
}

// ExtractNumericDigits - Examines an input parameter 'targetStr' to identify and extract the
// first instance of a number string. The number string will be comprised of one or more
// consecutive numeric digits (0-9) and may include leading, trailing or interior non-numeric
// characters as specified by input parameters.
//
// The search for this number string will be started at the index specified by input parameter
// 'startIdx'. Beginning at 'startIdx' the 'targetStr' will be searched to identify and extract
// the first instance of a number string.
//
// A number string is usually defined a string of consecutive numeric digits. However, this
// method allows the caller to include additional non-numeric characters as identified by
// input parameters	'keepLeadingChars', 'keepInteriorChars' and 'keepTrailingChars'.
//
// 'keepLeadingChars' is a string of characters which will be prefixed to the number string
// if those characters exist in 'targetStr' and immediately precede the number string.
//
// 'keepInteriorChars' is a string of characters which, if they exist within the number string,
// will be retained and presented in the final extracted number string.
//
// 'keepTrailingChars' is a string of characters which will be suffixed to the end of the
// final extracted number string.  To qualify, the designated 'keepTrailingChars' must immediately
// follow the number string contained in 'targetStr'.
//
// If successfully located within 'targetStr' the first instance of a number string along with
// characteristics describing that number string are returned in a Type 'NumStrProfileDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string    - The target string to be searched for the first instance of
//                                  a number string. A number string is usually defined as a
//                                  string comprised of one or more consecutive numeric digits.
//                                  Additional parameters provided by this method will allow
//                                  the caller to insert specified non-numeric characters at
//                                  the beginning, end or interior of a number string.
//
//  startIdx               int    - The starting index in input parameter 'targetStr'
//                                  from which the search for a number string will be
//                                  initiated. This useful in extracting multiple number
//                                  strings form a single 'targetStr'.
//
//  keepLeadingChars    string    - This string contains non-numeric characters which will be
//                                  retained as a prefix to the final number string extracted
//                                  from the 'targetStr' parameter. To be included, these characters
//                                  must exist in 'targetStr' and must immediately precede the
//                                  first instance of a number string.
//
//                                  For example, if the target string is "Hello $123789 world" and
//                                  parameter 'keepLeadingChars' includes the USA currency character,
//                                  '$', the returned number string would be '$123789'.  If no currency
//                                  character was included in 'keepLeadingChars', the returned number
//                                  string would be '123789'. It is worth noting that if the target
//                                  string was '$ 123789' and a currency symbol, '$', was included
//                                  in 'keepLeadingChars', the returned number string would still be
//                                  '123789' because 'keepLeadingChars' characters must immediately
//                                  precede the string of numeric digits in 'targetStr'.
//
//  keepInteriorChars   string    - This string contains non-numeric characters which will be retained
//                                  as valid characters within the final extracted number string. The
//                                  characters must exist withing the first instance of a number string
//                                  located in 'targetStr'. Such interior characters might include
//                                  thousands separators (commas) or decimal points (periods).
//
//                                  For example, if a comma and a period are included in 'keepInteriorChars'
//                                  and the target string is "Hello word 123,456,789.25 !", the returned
//                                  number string would be "123,456,789.25".  If the comma character was
//                                  NOT included in the 'keepInteriorChars' string, the returned number
//                                  string would be '123', since the number string extraction parser
//                                  would break on the comma, a non-numeric digit.
//
//                                  'keepInteriorChars' will NOT allow multiple non-numeric characters
//                                  to exist within the interior of the final extracted number string.
//                                  Only single non-numeric characters are allowed within a number string.
//
//  keepTrailingChars   string    - This string contains non-numeric characters which should be retained
//                                  at the end of the final number string. By default, a non-numeric
//                                  character will mark the end of a number string. However, if the caller
//                                  elects to use parameter 'keepTrailingChars' to retain non-numeric
//                                  characters such as a trailing right-parenthesis, then those non-numeric
//                                  characters will be retained in the final extracted number string.
//
//                                  It should be emphasized that 'keepTrailingChars' must immediately
//                                  follow the first instance of a number string in parameter, 'targetStr'.
//
//                                  Example #1:
//                                    Target String = "Hello world, (1234). Today is new day."
//                                    keepLeadingChars = "("
//                                    keepInteriorChars = ""
//                                    keepTrailingChars= ")"
//                                    Extracted Number String = "(1234)"
//
//                                  Example #2:
//                                    Target String = "Hello world, USA GDP growth is projected at 1.8%."
//                                    keepLeadingChars = ""
//                                    keepInteriorChars = "."
//                                    keepTrailingChars= "%"
//                                    Extracted Number String = "1.8%"
// ------------------------------------------------------------------------
//
// Return Value
//
//  NumStrProfileDto    - If successful, this method will return a type 'NumStrProfileDto'
//                        populated with the extracted number string and additional profile
//                        information related to the extracted number string.
//
//					type NumStrProfileDto struct {
//								TargetStr        string   // The original target string which is scanned for a number string
//								StartIndex          int   // The starting index in 'TargetStr' from which the number string
//																			    //    search was initiated.
//								LeadingSignIndex    int   //  The string index of a leading sign in 'NumStr' below. If a
//																			    //    leading sign character is NOT present in 'NumStr' this value
//                                          //    is set to -1
//                LeadingSignChar  string   //  If a leading sign character (plus '+' or minus '-') exists in
//                                          //    data field 'NumStr' (below), it is stored in this string.
//                FirstNumCharIndex   int   //  The index in 'TargetStr' (above) where the first character
//                                          //    of the extracted number string is located.
//                NextTargetStrIndex  int   //  The index of the next character in 'TargetStr' immediately
//                                          //    following the extracted number string.
//                NumStrLen           int   //  The length of the extracted number string.
//                NumStr           string   //  The number string extracted from 'TargetStr'.
//          }
//
//
//  error               - If 'startIndex' is less than zero or if 'startIndex' exceeds the last character
//                        index in 'targetStr', an error will be returned. If no errors are encountered,
//                        this value is set to 'nil'.
//
func (sops StrOps) ExtractNumericDigits(
  targetStr string,
  startIndex int,
  keepLeadingChars string,
  keepInteriorChars string,
  keepTrailingChars string) (NumStrProfileDto,	error) {

  nStrDto := NumStrProfileDto{}.New()
  nStrDto.TargetStr = targetStr
  nStrDto.StartIndex = startIndex

  var err error = nil

  ePrefix := "StrOps.ExtractNumericDigits() "

  lenTargetStr := len(targetStr)

  if lenTargetStr == 0 {
    return nStrDto, err
  }

  if startIndex < 0 {
    err = fmt.Errorf(ePrefix+
      "ERROR: Input parameter 'startIndex' is less than zero!\n"+
      "startIndex='%v'", startIndex)

    return nStrDto, err
  }

  if startIndex >= lenTargetStr {
    err = fmt.Errorf(ePrefix+
      "ERROR: Input parameter 'startIndex' is INVALID!\n"+
      "'startIndex' exceeds the last character index in 'targetStr'\n"+
      "startIndex='%v'\tlast character index='%v'\n"+
      "targetStr='%v'", startIndex, lenTargetStr-1, targetStr)

    return nStrDto, err
  }

  targetRunes := []rune(targetStr)
  lenTargetStr = len(targetRunes)

  keepLeadingRunes := make([]rune, 0)
  lenKeepLeadingRunes := 0

  keepInteriorRunes := make([]rune, 0)
  lenKeepInteriorRunes := 0

  keepTrailingRunes := make([]rune, 0)
  lenKeepTrailingRunes := 0

  if len(keepLeadingChars) > 0 {

    // Remove any numeric characters
    for a := 0; a < len(keepLeadingChars); a++ {

      if keepLeadingChars[a] >= '0' &&
        keepLeadingChars[a] <= '9' {
        continue
      }

      keepLeadingRunes = append(keepLeadingRunes, rune(keepLeadingChars[a]))

    }

    lenKeepLeadingRunes = len(keepLeadingRunes)
  }

  if len(keepInteriorChars) > 0 {

    // Remove any numeric characters
    for a := 0; a < len(keepInteriorChars); a++ {

      if keepInteriorChars[a] >= '0' &&
        keepInteriorChars[a] <= '9' {
        continue
      }

      keepInteriorRunes = append(keepInteriorRunes, rune(keepInteriorChars[a]))

    }

    lenKeepInteriorRunes = len(keepInteriorRunes)
  }

  if len(keepTrailingChars) > 0 {

    // Remove any numeric characters
    for a := 0; a < len(keepTrailingChars); a++ {

      if keepTrailingChars[a] >= '0' &&
        keepTrailingChars[a] <= '9' {
        continue
      }

      keepTrailingRunes = append(keepTrailingRunes, rune(keepTrailingChars[a]))

    }

    lenKeepTrailingRunes = len(keepTrailingRunes)
  }

  numberRunesCaptured := make([]rune, 0, 20)
  lenNumberRunesCaptured := 0

  leadingCharRunesCaptured := make([]rune, 0, 20)
  lenLeadingCharRunesCaptured := 0

  firstNumericDigitIdx := -1

  for e := startIndex; e < lenTargetStr; e++ {

    if targetRunes[e] >= '0' &&
      targetRunes[e] <= '9' &&
      firstNumericDigitIdx == -1 {
      // Target has at least one numeric
      // digit - and we found it.
      firstNumericDigitIdx = e
      break
    }
  }

  if firstNumericDigitIdx == -1 {
    // There are no numeric digits
    // in this target string.
    // EXIT HERE!!!
    return nStrDto, err
  }

  firstNumStrCharIdx := -1
  leadingSignChar := ""

  // Check for leading non-numeric characters that
  // need to be retained at the front of the number
  // string.
  if lenKeepLeadingRunes > 0 &&
    startIndex < firstNumericDigitIdx {

    for f := firstNumericDigitIdx - 1; f >= startIndex; f-- {

      for g := 0; g < lenKeepLeadingRunes; g++ {

        if keepLeadingRunes[g] == targetRunes[f] {

          if keepLeadingRunes[g] == '+' ||
            keepLeadingRunes[g] == '-' {

            // This is a leading sign char
            leadingSignChar = string(targetRunes[f])

            leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
            // Delete Leading Sign character. It will not be repeated in
            // future searches. Only one leading sign char per number string.

            keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
            lenKeepLeadingRunes--

            firstNumStrCharIdx = f

            // Now delete the alternative leading sign character.
            // There are only two - plus or minus
            nextSignChar := '-'

            if leadingSignChar == "-" {
              nextSignChar = '+'
            }

            // Leading sign char has been found. Now delete the
            // alternative lead sign char to avoid duplications
            for m := 0; m < lenKeepLeadingRunes; m++ {
              if keepLeadingRunes[m] == nextSignChar {
                keepLeadingRunes = append(keepLeadingRunes[0:m], keepLeadingRunes[m+1:]...)
                lenKeepLeadingRunes--
              }
            }

            break

          } else {

            // Standard Keep Leading Rune character found
            leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
            // Delete Leading Rune character. It will not be repeated in
            // future searches

            firstNumStrCharIdx = f

            keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
            lenKeepLeadingRunes--
            break
          }
        }
      }

      t := len(leadingCharRunesCaptured)

      if t > lenLeadingCharRunesCaptured {
        lenLeadingCharRunesCaptured = t
        continue
      }

      break
    }
  }

  leadingSignIndex := -1

  if lenLeadingCharRunesCaptured > 0 {

    for h := lenLeadingCharRunesCaptured - 1; h >= 0; h-- {

      if leadingCharRunesCaptured[h] == '+' ||
        leadingCharRunesCaptured[h] == '-' {

        numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
        leadingSignIndex = lenNumberRunesCaptured
        lenNumberRunesCaptured++

      } else {
        numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
        lenNumberRunesCaptured++
      }
    }
  }

  // Main Number String Extraction Loop
  isEndOfNumStr := false

  for i := firstNumericDigitIdx; i < lenTargetStr; i++ {

    if !isEndOfNumStr {

      if targetRunes[i] >= '0' && targetRunes[i] <= '9' {

        numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
        continue
      }

      for j := 0; j < lenKeepInteriorRunes; j++ {

        if targetRunes[i] == keepInteriorRunes[j] {

          if i+1 >= lenTargetStr ||
            (targetRunes[i+1] < '0' || targetRunes[i+1] > '9'){
            // We are either at the end of string or the next char
            // is NOT a numeric character.
            goto trailChar
          }


          numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])

          goto numDigitLoop
        }
      }

    }

    trailChar:
    isEndOfNumStr = true

    for k := 0; k < lenKeepTrailingRunes; k++ {

      if targetRunes[i] == keepTrailingRunes[k] {
        numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
        // Only one instance of a keep trailing rune character is captured.
        // Delete the keep trailing rune character to prevent repeat captures.
        keepTrailingRunes = append(keepLeadingRunes[0:k], keepTrailingRunes[k+1:]...)
        lenKeepTrailingRunes--
        goto numDigitLoop
      }

    }

    // Non-numeric character and Non-Trailing Character: Exit the Loop
    break

  numDigitLoop:
  }

  if len(numberRunesCaptured) > 0 {
    nStrDto.NumStr = string(numberRunesCaptured)

    if firstNumStrCharIdx > -1 {
      nStrDto.FirstNumCharIndex = firstNumStrCharIdx
    } else {
      nStrDto.FirstNumCharIndex = firstNumericDigitIdx
    }

    nStrDto.NumStrLen = len(nStrDto.NumStr)
    nStrDto.LeadingSignChar = leadingSignChar
    nStrDto.LeadingSignIndex = leadingSignIndex
    nStrDto.NextTargetStrIndex =
      nStrDto.FirstNumCharIndex + nStrDto.NumStrLen

    if nStrDto.NextTargetStrIndex >= len(targetStr) {
      nStrDto.NextTargetStrIndex = -1
    }
  }

  err = nil

  return nStrDto, err
}

// FindFirstNonSpaceChar - Returns the string index of the first non-space character in
// a string segment. The string to be searched is input parameter 'targetStr'. The string
// segment which will be searched from left to right in 'targetStr' is defined by the
// starting index ('startIndex') and the ending index ('endIndex').
//
// Searching from left to right, this method identifies the first non-space character
// (any character that is NOT a space ' ') in the target string segment and returns
// the index associated with that non-space character.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	This method returns the index of the first non-space character in the target string
//	segment using a left to right search. If the entire string consists of space characters,
//	this method returns a value of -1.
func (sops StrOps) FindFirstNonSpaceChar(targetStr string, startIndex, endIndex int) (int, error) {

  ePrefix := "StrOps.FindFirstNonSpaceChar() "

  targetStrLen := len(targetStr)

  if targetStrLen == 0 {
    return -1, nil
  }

  if startIndex < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIndex' is LESS THAN ZERO! "+
      "startIndex='%v' ", startIndex)
  }

  if endIndex < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIndex' is LESS THAN ZERO! "+
      "startIndex='%v' ", startIndex)
  }

  if endIndex >= targetStrLen {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIndex' is greater than "+
      "target string length. INDEX OUT OF RANGE! endIndex='%v' target string length='%v' ",
      endIndex, targetStrLen)
  }

  if startIndex > endIndex {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIndex' is GREATER THAN 'endIndex' "+
      "startIndex='%v' endIndex='%v' ", startIndex, endIndex)
  }

  idx := startIndex

  for idx <= endIndex {

    if targetStr[idx] != ' ' {
      return idx, nil
    }

    idx++
  }

  return -1, nil
}

// FindLastNonSpaceChar - Returns the string index of the last non-space character in a
// string segment.  The string to be searched is input parameter, 'targetStr'. The
// string segment is further defined by input parameters 'startIdx' and  'endIdx'. These
// indexes define a segment within 'targetStr' which will be searched to identify the last
// non-space character.
//
// The search is a backwards search, from right to left, conducted within the defined
// 'targetStr' segment. The search therefore starts at 'endIdx' and proceeds towards
// 'startIdx' until the last non-space character in the string segment is identified.
//
//
// If the last non-space character is found, that string index is returned. If the string
// segment consists entirely of space characters, the return value is -1.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
func (sops StrOps) FindLastNonSpaceChar(targetStr string, startIdx, endIdx int) (int, error) {

  ePrefix := "StrOps.FindLastNonSpaceChar() "

  targetStrLen := len(targetStr)

  if targetStrLen == 0 {
    return -1, errors.New(ePrefix +
      "ERROR: Invalid input parameter. 'targetStr' is a ZERO LENGTH STRING! ")
  }

  if startIdx < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIdx' is LESS THAN ZERO! "+
      "startIdx='%v' ", startIdx)
  }

  if endIdx < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIdx' is LESS THAN ZERO! "+
      "startIdx='%v' ", startIdx)
  }

  if endIdx >= targetStrLen {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIdx' is greater than "+
      "target string length. INDEX OUT OF RANGE! endIdx='%v' target string length='%v' ",
      endIdx, targetStrLen)
  }

  if startIdx > endIdx {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIdx' is GREATER THAN 'endIdx' "+
      "startIdx='%v' endIdx='%v' ", startIdx, endIdx)
  }

  floor := startIdx - 1

  for endIdx > floor {

    if targetStr[endIdx] != ' ' {
      return endIdx, nil
    }

    endIdx--
  }

  return -1, nil
}

// FindLastSpace - Returns a string index indicating the last space character (' ') in
// a string segment. The string segment is defined by input parameters, 'startIdx' and
// 'endIdx'.
//
// The string segment search proceeds backwards, from right to left. The search therefore
// starts at 'endIdx' and proceeds towards 'startIdx' until the last space character in
// the string segment is identified.
//
// If a valid index for the last space character is found in the string segment, that
// index value is returned. If a space character is NOT found in the specified string
// segment, a value of -1 is returned.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
func (sops StrOps) FindLastSpace(targetStr string, startIdx, endIdx int) (int, error) {

  ePrefix := "StrOps.FindLastSpace() "

  targetStrLen := len(targetStr)

  if targetStrLen == 0 {
    return -1, errors.New(ePrefix +
      "ERROR: Invalid input parameter. 'targetStr' is a ZERO LENGTH STRING! ")
  }

  if startIdx < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIdx' is LESS THAN ZERO! "+
      "startIdx='%v' ", startIdx)
  }

  if endIdx < 0 {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIdx' is LESS THAN ZERO! "+
      "startIdx='%v' ", startIdx)
  }

  if endIdx >= targetStrLen {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIdx' is greater than "+
      "target string length. INDEX OUT OF RANGE! endIdx='%v' target string length='%v' ",
      endIdx, targetStrLen)
  }

  if startIdx > endIdx {
    return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIdx' is GREATER THAN 'endIdx' "+
      "startIdx='%v' endIdx='%v' ", startIdx, endIdx)
  }

  for endIdx >= startIdx {

    if targetStr[endIdx] == ' ' {
      return endIdx, nil
    }

    endIdx--
  }

  return -1, nil
}

// FindLastWord - Returns the beginning and ending indexes of
// the last word in a target string segment. A 'word' is defined here
// as a contiguous set of non-space characters delimited by spaces or
// the beginning and ending indexes of the target string segment. Note,
// for purposes of this method, a 'word' my consist of a single non-space
// character such as an article 'a' or a punctuation mark '.'
//
// ------------------------------------------------------------------------
//
// Examples
//
//
//	Example (1)
// 		In the text string segment:
//
//			"The cow jumped over the moon."
//
//		The last word would be defined as "moon."
//
//	Example (2)
//		In the text string segment:
//
//			"  somewhere over the rainbow  "
//
//		The last word would be defined as "rainbow"
//
// ------------------------------------------------------------------------
//
// The string to be searched is contained in input parameter, 'targetStr'.
// The string segment within 'targetStr' is defined by input parameters
// 'startIndex' and 'endIndex'.
//
// If the entire string segment is classified as a 'word', meaning that
// there are no space characters in the string segment, the returned
// values for 'beginWrdIdx' and 'endWrdIdx' will be equal to the input
// parameters 'startIndex' and 'endIndex'.
//
// If the string segment is consists entirely of space characters, the
// returned 'beginWrdIdx' and 'endWrdIdx' will be set equal to -1 and
// the returned value, 'isAllSpaces' will be set to 'true'.
//
// If 'targetStr' is an empty string, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//	targetStr	string	- The string containing the string segment which
//                      		will be searched to identify the last word
//                      		in the string segment.
//
//	startIndex	int	- The index marking the beginning of the string
//                      		segment in 'targetStr'.
//
//	endIndex	int	- The index marking the end of the string segment
//                      		in 'targetStr'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	beginWrdIdx	int	-	The index marking the beginning of the last word
// 					in the string segment identified by input parameters
//					'startIndex' and 'endIndex'. If the string segment
//					consists of all spaces or is empty, this value is
//					set to -1.
//
//	endWrdIdx	int	-	The index marking the end of the last word in the
//					string segment identified by input parameters 'startIndex'
//					and 'endIndex'. If the string segment consists of all
//					spaces or is empty, this value is set to -1.
//
//	isAllOneWord	bool	-	If the string segment identified by input parameters
//					'startIndex' and 'endIndex' consists entirely of non-space
//					characters (characters other than ' '), this value is set
//					to 'true'.
//
//	isAllSpaces	bool	-	If the string segment identified by input parameters
//					'startIndex' and 'endIndex' consists entirely of space
//					characters (character = ' '), this value is set to 'true'.
//
//	err		error	-	If targetStr is empty or if startIndex or endIndex is invalid,
//					an error is returned. If the method completes successfully,
//					err = nil.
func (sops StrOps) FindLastWord(
  targetStr string,
  startIndex,
  endIndex int) (beginWrdIdx,
  endWrdIdx int,
  isAllOneWord,
  isAllSpaces bool,
  err error) {

  ePrefix := "StrOps.FindLastWord() "
  beginWrdIdx = -1
  endWrdIdx = -1
  isAllOneWord = false
  isAllSpaces = false

  targetStrLen := len(targetStr)

  if targetStrLen == 0 {

    err = fmt.Errorf(ePrefix + "Error: Input parameter 'targetStr' is an EMPTY STRING!")
    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  if startIndex < 0 {

    err = fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIndex' is LESS THAN ZERO! "+
      "startIndex='%v' ", startIndex)

    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  if endIndex < 0 {
    err = fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIndex' is LESS THAN ZERO! "+
      "startIndex='%v' ", startIndex)

    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  if endIndex >= targetStrLen {

    err = fmt.Errorf(ePrefix+
      "ERROR: Invalid input parameter. 'endIndex' is greater than "+
      "target string length. INDEX OUT OF RANGE! endIndex='%v' "+
      "target string length='%v' ",
      endIndex, targetStrLen)

    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  if startIndex > endIndex {
    err = fmt.Errorf(ePrefix+
      "ERROR: Invalid input parameter. 'startIndex' is GREATER THAN 'endIndex'. "+
      "startIndex='%v' endIndex='%v' ", startIndex, endIndex)

    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  beginWrdIdx = startIndex
  endWrdIdx = endIndex

  idx := endIndex

  endingIdxFound := false
  beginningIdxFound := false

  isAllSpaces = true
  isAllOneWord = true

  if startIndex == endIndex {

    beginWrdIdx = startIndex
    endWrdIdx = startIndex

    if targetStr[startIndex] == ' ' {
      isAllSpaces = true
      isAllOneWord = false
    } else {
      isAllSpaces = false
      isAllOneWord = true
    }

    err = nil

    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  for idx >= startIndex {

    if targetStr[idx] != ' ' {
      isAllSpaces = false
    } else {
      isAllOneWord = false
    }

    if !endingIdxFound &&
      targetStr[idx] != ' ' {

      endWrdIdx = idx
      endingIdxFound = true
      idx--
      continue
    }

    if endingIdxFound &&
      !beginningIdxFound &&
      targetStr[idx] == ' ' {

      beginWrdIdx = idx + 1
      beginningIdxFound = true
      break
    }

    idx--
  }

  if isAllSpaces {
    isAllOneWord = false
    beginWrdIdx = -1
    endWrdIdx = -1
    err = nil
    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  if isAllOneWord {
    beginWrdIdx = startIndex
    endWrdIdx = endIndex
    isAllSpaces = false
    err = nil
    return beginWrdIdx,
      endWrdIdx,
      isAllOneWord,
      isAllSpaces,
      err
  }

  err = nil

  return beginWrdIdx,
    endWrdIdx,
    isAllOneWord,
    isAllSpaces,
    err
}

// FindRegExIndex - returns a two-element slice of integers defining the location
// of the leftmost match in targetStr of the regular expression (regex).
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	The return value is an array of integers. If no match is found the return
//	value is 'nil'.  If regular expression is successfully matched, the match
//	will be located at targetStr[loc[0]:loc[1]]. Again, a return value of 'nil'
//	signals that no match was found.
func (sops StrOps) FindRegExIndex(targetStr string, regex string) []int {

  re := regexp.MustCompile(regex)

  return re.FindStringIndex(targetStr)

}

// GetCountBytesRead - Returns private member variable
// 'StrOps.cntBytesRead' which holds the number of bytes
// accumulated through the last Read operation executed through
// method StrOps.Read().
func (sops *StrOps) GetCountBytesRead() uint64 {

  var bytesRead uint64

  sops.stringDataMutex.Lock()
  bytesRead = sops.cntBytesRead
  sops.stringDataMutex.Unlock()

  return bytesRead
}

// GetCountBytesWritten - Returns private member variable
// 'StrOps.cntBytesWritten' which holds the number of bytes
// accumulated through the last Write operation executed
// through method StrOps.Write().
func (sops *StrOps) GetCountBytesWritten() uint64 {

  var bytesWritten uint64

  sops.stringDataMutex.Lock()
  bytesWritten = sops.cntBytesWritten
  sops.stringDataMutex.Unlock()

  return bytesWritten
}

// GetReader() Returns an io.Reader which will read the private
// member data element StrOps.stringData.
func (sops *StrOps) GetReader() io.Reader {
  var stringData string

  sops.stringDataMutex.Lock()
  stringData = sops.stringData
  sops.stringDataMutex.Unlock()
  return strings.NewReader(stringData)
}

// GetSoftwareVersion - Returns the software version for package 'strops'.
func (sops StrOps) GetSoftwareVersion() string {
  return "2.0.4"
}

// GetStringData - Returns the current value of internal
// member string, StrOps.stringData
func (sops *StrOps) GetStringData() string {
  var output string

  sops.stringDataMutex.Lock()
  output = sops.stringData
  sops.cntBytesWritten = 0
  sops.cntBytesRead = 0
  sops.stringDataMutex.Unlock()

  return output
}

// GetValidBytes - Receives an array of 'targetBytes' which will be examined to determine
// the validity of individual bytes or characters. Each character (byte) in input array
// 'targetBytes' will be compared to input parameter 'validBytes', another array of bytes.
// If a character in 'targetBytes' also exists in 'validBytes' it will be considered valid
// and included in the returned array of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	targetBytes	[] byte	- An array of characters (bytes) which will be examined
//				for valid characters. The list of valid characters is
//				found in input parameter 'validBytes'. Valid characters
//				in targetBytes will be returned by this method as an
//				array of bytes. Invalid characters will be discarded.
//
//	validBytes	[] byte	- An array of bytes containing valid characters. If a character
//				(byte) in 'targetBytes' is also present in 'validBytes' it will
//				be classified as 'valid' and included in the returned array of
//				bytes. Invalid characters will be discarded.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[] byte	- An array of bytes which contains bytes that are present in both 'targetBytes'
//		and 'validBytes'. Note: If all characters in 'targetBytes' are classified as
//		'invalid', the returned array of bytes will be a zero length array.
//
//	error	- If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length 'targetBytes array or 'validBytes' array.
func (sops StrOps) GetValidBytes(targetBytes, validBytes []byte) ([]byte, error) {

  ePrefix := "StrOps.GetValidBytes() "

  lenTargetBytes := len(targetBytes)

  output := make([]byte, 0, lenTargetBytes+10)

  if lenTargetBytes == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'targetBytes' is a ZERO LENGTH ARRAY!")
  }

  lenValidBytes := len(validBytes)

  if lenValidBytes == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'validBytes' is a ZERO LENGTH ARRAY!")
  }

  for i := 0; i < lenTargetBytes; i++ {

    for j := 0; j < lenValidBytes; j++ {
      if targetBytes[i] == validBytes[j] {
        output = append(output, targetBytes[i])
        break
      }
    }

  }

  return output, nil
}

// GetValidRunes - Receives an array of 'targetRunes' which will be examined to determine
// the validity of individual runes or characters. Each character (rune) in input array
// 'targetRunes' will be compared to input parameter 'validRunes', another array of runes.
// If a character in 'targetRunes' also exists in 'validRunes', that character will be considered
// valid and included in the returned array of runes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	targetRunes		[] rune	- An array of characters (runes) which will be examined
//					for valid characters. The list of valid characters is
//					found in input parameter 'validRunes'. Valid characters
//					in targetRunes will be returned by this method as an
//					array of runes. Invalid characters will be discarded.
//
//	validRunes		[] rune	- An array of runes containing valid characters. If a character
//					(rune) in targetRunes is also present in 'validRunes' it will
//					be classified as 'valid' and included in the returned array of
//					runes. Invalid characters will be discarded.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[] rune	- An array of runes which contains runes that are present in 'targetRunes' and
//		'validRunes'. Note: If all characters in 'targetRunes' are classified as
//		'invalid', the returned array of runes will be a zero length array.
//
//	error - If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length 'targetRunes array or 'validRunes' array.
func (sops StrOps) GetValidRunes(targetRunes []rune, validRunes []rune) ([]rune, error) {

  ePrefix := "StrOps.GetValidRunes() "

  lenTargetRunes := len(targetRunes)

  output := make([]rune, 0, lenTargetRunes+10)

  if lenTargetRunes == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'targetRunes' is a ZERO LENGTH ARRAY!")
  }

  lenValidRunes := len(validRunes)

  if lenValidRunes == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'validRunes' is a ZERO LENGTH ARRAY!")
  }

  for i := 0; i < lenTargetRunes; i++ {

    for j := 0; j < lenValidRunes; j++ {
      if targetRunes[i] == validRunes[j] {
        output = append(output, targetRunes[i])
        break
      }
    }

  }

  return output, nil
}

// GetValidString - Validates the individual characters in input parameter string,
// 'targetStr'. To identify valid characters, the characters in 'targetStr' are
// compared against input parameter 'validRunes', an array of type rune. If a character
// exists in both 'targetStr' and 'validRunes' it is deemed valid and returned in
// an output string.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	targetStr	string	-	The string which will be screened for valid characters.
//
//	validRunes	[] rune	-	An array of type rune containing valid characters. Characters
//					which exist in both 'targetStr' and 'validRunes' will be
//					returned as a new string. Invalid characters are discarded.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string	- This string will be returned containing valid characters extracted
//		from 'targetStr'. A character is considered valid if it exists in
//		both 'targetStr' and 'validRunes'. Invalid characters are discarded.
//		This means that if no valid characters are identified, a zero length
//		string will be returned.
//
//	error - If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length 'targetStr' (string) or a zero length
//		'validRunes' array.
func (sops StrOps) GetValidString(targetStr string, validRunes []rune) (string, error) {

  ePrefix := "StrOps.GetValidString() "

  if len(targetStr) == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'targetStr' is a ZERO LENGTH STRING!")
  }

  if len(validRunes) == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'validRunes' is a ZERO LENGTH ARRAY!")
  }

  validRunes, err := sops.GetValidRunes([]rune(targetStr), validRunes)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by sops.GetValidRunes([]rune(targetStr), validRunes). "+
        "Error='%v' ", err.Error())
  }

  return string(validRunes), nil
}

// IsEmptyOrWhiteSpace - If a string is zero length or consists solely of
// white space (contiguous spaces), this method will return 'true'.
//
// Otherwise, a value of false is returned.
func (sops StrOps) IsEmptyOrWhiteSpace(targetStr string) bool {

  targetLen := len(targetStr)

  for i := 0; i < targetLen; i++ {
    if targetStr[i] != ' ' {
      return false
    }
  }

  return true
}

// LowerCaseFirstLetter - Finds the first alphabetic character
// in a string (a-z A-Z) and converts it to lower case.
func (sops StrOps) LowerCaseFirstLetter(str string) string {

  if len(str) == 0 {
    return str
  }

  runeStr := []rune(str)

  for i := 0; i < len(runeStr); i++ {

    if runeStr[i] == ' ' {
      continue
    }

    if runeStr[i] >= 'A' &&

      runeStr[i] <= 'Z' {

      runeStr[i] += 32

      break

    } else if runeStr[i] >= 'a' &&

      runeStr[i] <= 'z' {

      break
    }

  }

  return string(runeStr)
}

// MakeSingleCharString - Creates a string of length 'strLen' consisting of
// a single character passed through input parameter, 'charRune' as type
// 'rune'.
func (sops StrOps) MakeSingleCharString(charRune rune, strLen int) (string, error) {

  ePrefix := "StrOps.MakeSingleCharString() "

  if strLen < 1 {
    return "",
      fmt.Errorf(ePrefix+"Error: Input parameter 'strLen' MUST BE GREATER THAN '1'. "+
        "strLen='%v' ", strLen)
  }

  if charRune == 0 {
    return "",
      fmt.Errorf(ePrefix+"Error: Input parameter 'charRune' IS INVALID! "+
        "charRune='%v' ", charRune)
  }

  var b strings.Builder
  b.Grow(strLen + 1)

  for i := 0; i < strLen; i++ {

    _, err := b.WriteRune(charRune)

    if err != nil {
      return "",
        fmt.Errorf(ePrefix+
          "Error returned by  b.WriteRune(charRune). "+
          "charRune='%v' Error='%v' ", charRune, err.Error())
    }
  }

  return b.String(), nil
}

// NewPtr - Returns a pointer to a new instance of
// StrOps. Useful for cases requiring io.Reader
// and io.Writer.
func (sops StrOps) NewPtr() *StrOps {

  sopsNew := StrOps{}

  return &sopsNew
}

// Read - Implements io.Reader interface. Read reads up to len(p)
// bytes into 'p'. This method supports buffered 'read' operations.
//
// The internal member string variable, 'StrOps.stringData' is written
// into 'p'.
//
// 'StrOps.stringData' can be accessed through Getter an Setter methods,
// GetStringData() and SetStringData()
//
func (sops *StrOps) Read(p []byte) (n int, err error) {

  ePrefix := "StrOps.Read() "

  n = len(p)
  err = io.EOF

  sops.stringDataMutex.Lock()

  if n == 0 {
    sops.cntBytesRead = 0
    sops.stringDataMutex.Unlock()
    err = errors.New(ePrefix + "Error: Input byte array 'p' is zero length!")
    return 0, err
  }

  strData := sops.stringData

  w := []byte(strData)

  lenW := uint64(len(w))

  cntBytesRead := sops.cntBytesRead

  if lenW == 0 ||
    cntBytesRead >= lenW {
    sops.cntBytesRead = 0
    n = 0
    sops.stringDataMutex.Unlock()
    return n, err
  }

  sops.stringDataMutex.Unlock()

  startReadIdx := cntBytesRead

  remainingBytesToRead := lenW - cntBytesRead

  if uint64(n) < remainingBytesToRead {
    remainingBytesToRead = startReadIdx + uint64(n)
    err = nil
  } else {
    remainingBytesToRead += startReadIdx
    err = io.EOF
  }

  n = 0
  i := uint64(0)
  for i = startReadIdx; i < remainingBytesToRead; i++ {
    p[n] = w[i]
    n++
  }

  cntBytesRead += uint64(n)

  sops.stringDataMutex.Lock()

  if cntBytesRead >= lenW {
    sops.cntBytesRead = 0
  } else {
    sops.cntBytesRead = cntBytesRead
  }

  sops.stringDataMutex.Unlock()

  return n, err
}

// ReadStringFromBytes - Receives a byte array and retrieves a string. The beginning of
// the string is designated by input parameter 'startIdx'. The end of the string is determined
// when a carriage return ('\r'), vertical tab ('\v') or a new line character ('\n') is encountered.
//
// The parsed string is returned to the caller along with 'nextStartIdx', which is the byte
// array index of the beginning of the next string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	bytes          []byte - An array of bytes from which a string will be extracted
//	                        and returned.
//
//	startIdx          int - The starting index in input parameter 'bytes' where the string
//	                        extraction will begin. The string extraction will cease when
//	                        a carriage return ('\r'), a vertical tab ('\v') or a new line
//	                        character ('\n') is encountered.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	extractedStr   string - The string extracted from input parameter 'bytes' beginning
//	                        at the index in 'bytes' indicated by input parameter 'startIdx'.
//
//	nextStartIdx      int - The index of the beginning of the next string in the byte array
//	                        'bytes' after 'extractedString'. If no more strings exist in the
//	                        the byte array, 'nextStartIdx' will be set to -1.
//
func (sops StrOps) ReadStringFromBytes(
  bytes []byte,
  startIdx int) (extractedStr string, nextStartIdx int) {

  extractedStr = ""
  nextStartIdx = -1

  bLen := len(bytes)

  if bLen == 0 {
    return extractedStr, nextStartIdx
  }

  if startIdx >= bLen || startIdx < 0 {
    return extractedStr, nextStartIdx
  }

  nextStartIdx = -1

  runeAry := make([]rune, 0, bLen+5)

  for i := startIdx; i < bLen; i++ {

    if bytes[i] == '\r' ||
      bytes[i] == '\n' ||
      bytes[i] == '\v' {

      if i+1 < bLen {

        j := 0

        for j = i + 1; j < bLen; j++ {
          if bytes[j] == '\r' ||
            bytes[j] == '\v' ||
            bytes[j] == '\n' {
            continue
          } else {
            break
          }
        }

        if j >= bLen {
          nextStartIdx = -1
        } else {
          nextStartIdx = j
        }

        break

      } else {
        nextStartIdx = -1
      }

      break
    }

    runeAry = append(runeAry, rune(bytes[i]))
  }

  extractedStr = string(runeAry)

  return extractedStr, nextStartIdx
}

// ReplaceBytes	- Replaces characters in a target array of bytes ([]bytes) with those specified in
// a two dimensional slice of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//	targetBytes	[]byte	- The byte array which will be examined. If characters ('bytes') eligible
//				for replacement are identified by replacementBytes[i][0] they will be
//				replaced by the character specified in replacementBytes[i][1].
//
//	replacementBytes [][]byte - A two dimensional slice of type byte. Element [i][0] contains the target
//				character to locate in 'targetBytes'. Element[i][1] contains the replacement
//				character which will replace the target character in 'targetBytes'. If
//				the replacement character element [i][1] is a zero value, the target character
//				will not be replaced. Instead, it will be eliminated or removed from
//				the returned byte array ([]byte).
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]byte	- The returned byte array containing the characters and replaced characters
//		from the original 'targetBytes' array.
//
//	error	- If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length targetBytes[] array or replacementBytes[][] array.
//		In addition, if any of the replacementBytes[][x] 2nd dimension elements have
//		a length less than two, an error will be returned.
func (sops StrOps) ReplaceBytes(targetBytes []byte, replacementBytes [][]byte) ([]byte, error) {

  ePrefix := "StrOps.ReplaceBytes() "

  output := make([]byte, 0, 100)

  targetLen := len(targetBytes)

  if targetLen == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'targetBytes' is a zero length array!")
  }

  baseReplaceLen := len(replacementBytes)

  if baseReplaceLen == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'replacementBytes' is a zero length array!")
  }

  for h := 0; h < baseReplaceLen; h++ {

    if len(replacementBytes[h]) < 2 {
      return output,
        fmt.Errorf(ePrefix+
          "Error: Invalid Replacement Array Element. replacementBytes[%v] has "+
          "a length less than two. ", h)
    }

  }

  foundReplacement := false

  for i := 0; i < targetLen; i++ {

    foundReplacement = false

    for k := 0; k < baseReplaceLen; k++ {

      if targetBytes[i] == replacementBytes[k][0] {

        if replacementBytes[k][1] != 0 {
          output = append(output, replacementBytes[k][1])
        }

        foundReplacement = true
        break
      }
    }

    if !foundReplacement {
      output = append(output, targetBytes[i])
    }

  }

  return output, nil
}

// ReplaceMultipleStrs - Replaces all instances of string replaceArray[i][0] with
// replacement string from replaceArray[i][1] in 'targetStr'.
//
// Note: The original 'targetStr' is NOT altered.
//
// Input parameter 'replaceArray' should be passed as a two-dimensional slice.
// If the length of the 'replaceArray' second dimension is less than '2', an
// error will be returned.
func (sops StrOps) ReplaceMultipleStrs(targetStr string, replaceArray [][]string) (string, error) {

  ePrefix := "StrOps.ReplaceMultipleStrs() "

  if targetStr == "" {
    return targetStr,
      errors.New(ePrefix + "Input parameter 'targetStr' is an EMPTY STRING.")
  }

  if len(replaceArray) == 0 {
    return "",
      errors.New(ePrefix +
        "Length of first dimension [X][] in two dimensional array 'replaceArray' is ZERO!")
  }

  newString := targetStr

  for aIdx, aVal := range replaceArray {

    if len(aVal) < 2 {
      return "",
        fmt.Errorf(ePrefix+
          "Length of second dimension [][X] in two dimensional array 'replaceArray' is Less Than 2! "+
          "replaceArray[%v][]", aIdx)
    }

    newString = strings.Replace(newString, replaceArray[aIdx][0], replaceArray[aIdx][1], -1)

  }

  return newString, nil
}

// RemoveNewLines - Replaces New Line characters from string. If the specified
// replacement string is empty, the New Line characters are simply removed
// from the input parameter, 'targetStr'.
func (sops StrOps) ReplaceNewLines(targetStr string, replacement string) string {

  return strings.Replace(targetStr, "\n", replacement, -1)
}

// ReplaceRunes - Replaces characters in a target array of runes ([]rune) with those specified in
// a two-dimensional slice of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//	targetRunes	[]rune	- The rune array which will be examined. If target characters ('runes')
//				eligible for replacement are identified by replacementRunes[i][0], they
//				will be replaced by the character specified in replacementRunes[i][1].
//
//	replacementRunes [][]rune - A two dimensional slice of type 'rune'. Element [i][0] contains the target
//				 character to locate in 'targetRunes'. Element[i][1] contains the
//				 replacement character which will replace the target character in
//				 'targetRunes'. If the replacement character element [i][1] is a zero value,
//				 the	target character will not be replaced. Instead, it will be eliminated
//				 or removed from the returned rune array ([]rune).
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]rune	- The returned rune array containing the characters and replaced characters
//		from the original 'targetRunes' array.
//
//	error	- If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length 'targetRunes' array or 'replacementRunes' array.
//		In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//		a length less than two, an error will be returned.
func (sops StrOps) ReplaceRunes(targetRunes []rune, replacementRunes [][]rune) ([]rune, error) {

  ePrefix := "StrOps.ReplaceRunes() "

  output := make([]rune, 0, 100)

  targetLen := len(targetRunes)

  if targetLen == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'targetRunes' is a zero length array!")
  }

  baseReplaceLen := len(replacementRunes)

  if baseReplaceLen == 0 {
    return output,
      errors.New(ePrefix + "Error: Input parameter 'replacementRunes' is a zero length array!")
  }

  for h := 0; h < baseReplaceLen; h++ {

    if len(replacementRunes[h]) < 2 {
      return output,
        fmt.Errorf(ePrefix+
          "Error: Invalid Replacement Array Element. replacementRunes[%v] has "+
          "a length less than two. ", h)
    }

  }

  foundReplacement := false

  for i := 0; i < targetLen; i++ {

    foundReplacement = false

    for k := 0; k < baseReplaceLen; k++ {

      if targetRunes[i] == replacementRunes[k][0] {

        if replacementRunes[k][1] != 0 {
          output = append(output, replacementRunes[k][1])
        }

        foundReplacement = true
        break
      }
    }

    if !foundReplacement {
      output = append(output, targetRunes[i])
    }

  }

  return output, nil
}

// ReplaceStringChars - Replaces string characters in a target string ('targetStr') with those
// specified in a two dimensional slice of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//	targetStr	string		- The string which will be examined. If target string characters
//					eligible for replacement are identified by replacementRunes[i][0], they
//					will be replaced by the character specified in replacementRunes[i][1].
//
//	replacementRunes [][]rune	- A two dimensional slice of type 'rune'. Element [i][0] contains the target
//				 	character to locate in 'targetStr'. Element[i][1] contains the
//					replacement character which will replace the target character in
//					'targetStr'. If the replacement character element [i][1] is a zero value,
//					the target character will not be replaced. Instead, it will be eliminated
//					or removed from the returned string.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string	- The returned string containing the characters and replaced characters
//		from the original target string, ('targetStr').
//
//	error	- If the method completes successfully this value is 'nil'. If an error is
//		encountered this value will contain the error message. Examples of possible
//		errors include a zero length 'targetStr' or 'replacementRunes[][]' array.
//		In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//		a length less than two, an error will be returned.
func (sops StrOps) ReplaceStringChars(
  targetStr string,
  replacementRunes [][]rune) (string, error) {

  ePrefix := "StrOps.ReplaceStringChars() "

  if len(targetStr) == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'targetStr' is an EMPTY STRING!")
  }

  if len(replacementRunes) == 0 {
    return "",
      errors.New(ePrefix + "Error: Input parameter 'replacementRunes' is an EMPTY STRING!")
  }

  outputStr, err := sops.ReplaceRunes([]rune(targetStr), replacementRunes)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+"Error returned by ReplaceRunes([]rune(targetStr), replacementRunes). "+
        "Error='%v' ", err.Error())
  }

  return string(outputStr), nil
}

// ResetBytesReadCounter - Resets the internal 'Bytes Read' counter
// to zero. As practical matter, this method is rarely if ever used.
// Its use is restricted primarily to special circumstances or debugging
// operations.
//
// This method sets StrOps.cntBytesRead equal to zero.
func (sops *StrOps) ResetBytesReadCounter() {
  sops.stringDataMutex.Lock()
  sops.cntBytesRead = 0
  sops.stringDataMutex.Unlock()
}

// ResetBytesWrittenCounter - Resets the internal 'Bytes Written' counter
// to zero. As practical matter, this method is rarely if ever used.
// Its use is restricted primarily to special circumstances or debugging
// operations.
//
// This method sets StrOps.cntBytesWritten equal to zero.
func (sops *StrOps) ResetBytesWrittenCounter() {
  sops.stringDataMutex.Lock()
  sops.cntBytesWritten = 0
  sops.stringDataMutex.Unlock()
}

// SetStringData - Sets the value of internal
// string data element, StrOps.stringData.
func (sops *StrOps) SetStringData(str string) {
  sops.stringDataMutex.Lock()
  sops.stringData = str
  sops.cntBytesWritten = 0
  sops.cntBytesRead = 0
  sops.stringDataMutex.Unlock()
}

// StrCenterInStrLeft - returns a string which includes
// a left pad blank string plus the original string. It
// does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively
// center the original string is a field of specified length.
func (sops StrOps) StrCenterInStrLeft(strToCenter string, fieldLen int) (string, error) {

  ePrefix := "StrOps.StrCenterInStrLeft() "

  if sops.IsEmptyOrWhiteSpace(strToCenter) {
    return strToCenter,
      errors.New(ePrefix +
        "Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!")
  }

  pad, err := sops.StrPadLeftToCenter(strToCenter, fieldLen)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by sops.StrPadLeftToCenter(strToCenter, fieldLen). "+
        "Error='%v'", err.Error())
  }

  return pad + strToCenter, nil

}

// StrCenterInStr - returns a string which includes a left pad blank string plus
// the original string ('strToCenter'), plus a right pad blank string.
//
// The complete string will effectively center the original string is a field of
// specified length ('fieldLen').
func (sops StrOps) StrCenterInStr(strToCenter string, fieldLen int) (string, error) {

  ePrefix := "StrOps.StrCenterInStr() "

  if sops.IsEmptyOrWhiteSpace(strToCenter) {
    return strToCenter,
      errors.New(ePrefix +
        "Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!")
  }

  sLen := len(strToCenter)

  if sLen > fieldLen {
    return strToCenter,
      fmt.Errorf(ePrefix+
        "Error: 'fieldLen' = '%v' strToCenter Length= '%v'. "+
        "'fieldLen is shorter than strToCenter Length!", fieldLen, sLen)
  }

  if sLen == fieldLen {
    return strToCenter, nil
  }

  leftPadCnt := (fieldLen - sLen) / 2

  leftPadStr := strings.Repeat(" ", leftPadCnt)

  rightPadCnt := fieldLen - sLen - leftPadCnt

  rightPadStr := ""

  if rightPadCnt > 0 {
    rightPadStr = strings.Repeat(" ", rightPadCnt)
  }

  return leftPadStr + strToCenter + rightPadStr, nil
}

// StrGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (sops StrOps) StrGetRuneCnt(targetStr string) int {
  return utf8.RuneCountInString(targetStr)
}

// StrGetCharCnt - Uses the 'len' method to
// return the number of rune characters in a
// string.
func (sops StrOps) StrGetCharCnt(targetStr string) int {
  return len([]rune(targetStr))
}

// StripBadChars - Removes/deletes specified characters from a string.
// The characters to remove are contained in a string array passed as
// input parameter, 'badChars'.
//
// All instances of a 'badChars' character are deleted from the target
// string. The target string is passed through input parameter, 'targetStr'.
//
func (sops StrOps) StripBadChars(
  targetStr string,
  badChars []string) (cleanStr string, strLen int) {

  cleanStr = targetStr
  strLen = len(cleanStr)

  if strLen == 0 {
    return cleanStr, strLen
  }

  lenBadChars := len(badChars)

  if lenBadChars == 0 {
    return cleanStr, strLen
  }

  sort.Sort(SortStrLengthHighestToLowest(badChars))

  cycleWhereStringRemoved := 0
  k := -1
  badCharIdx := -1

  for {

    k++

    for i := 0; i < lenBadChars; i++ {

      for {

        badCharIdx = strings.Index(cleanStr, badChars[i])

        if badCharIdx == -1 {
          break
        }

        lastCleanStrIdx := strLen - 1
        lChar := len(badChars[i])
        nextIdx := badCharIdx + lChar

        if nextIdx > lastCleanStrIdx {
          cleanStr = cleanStr[0:badCharIdx]

        } else {

          cleanStr = cleanStr[0:badCharIdx] + cleanStr[nextIdx:]
        }

        cycleWhereStringRemoved = k
      }

      strLen = len(cleanStr)

      if strLen == 0 {
        goto Done
      }
    }

    if k-cycleWhereStringRemoved > 3 || k > 1000000 {
      goto Done
    }
  }

Done:

  return cleanStr, strLen
}

// StripLeadingChars - Strips or deletes characters specified by
// input parameter 'badChars' from the front of 'targetStr'.
//
// The method then returns a string which does not contain leading
// 'bad characters'. In addition, the length of the 'clean string'
// is also returned.
//
func (sops StrOps) StripLeadingChars(
  targetStr string,
  badChars []string) (cleanStr string, strLen int) {

  cleanStr = targetStr
  strLen = len(cleanStr)

  lenBadChars := len(badChars)

  if lenBadChars == 0 {
    return cleanStr, strLen
  }

  if strLen == 0 {
    return cleanStr, strLen
  }

  sort.Sort(SortStrLengthHighestToLowest(badChars))

  cycleWhereStringRemoved := 0
  k := -1

  for {

    k++

    for i := 0; i < lenBadChars; i++ {

      for {

        if !strings.HasPrefix(cleanStr, badChars[i]) {
          break
        }

        cleanStr = cleanStr[len(badChars[i]):]

        cycleWhereStringRemoved = k
      }

      strLen = len(cleanStr)

      if strLen == 0 {
        goto Done
      }
    }

    if k-cycleWhereStringRemoved > 3 || k > 1000000 {
      goto Done
    }
  }

Done:

  return cleanStr, strLen
}

// StripTrailingChars - Strips or deletes bad characters from the
// end of input parameter 'targetStr'. Bad characters are identified
// by the input parameter, 'badChars'.
//
// The method then returns the cleaned string and the length of the
// cleaned string to the caller.  The cleaned string is equivalent
// to input parameter 'targetStr' minus the trailing bad characters
// at the end of 'targetStr'.
//
func (sops StrOps) StripTrailingChars(
  targetStr string,
  badChars []string) (cleanStr string, strLen int) {

  cleanStr = targetStr
  strLen = len(cleanStr)

  lenBadChars := len(badChars)

  if lenBadChars == 0 {
    return cleanStr, strLen
  }

  if strLen == 0 {
    return cleanStr, strLen
  }

  k := -1
  cycleWhereStringRemoved := 0

  for {

    k++

    for i := 0; i < lenBadChars; i++ {

      for {

        if !strings.HasSuffix(cleanStr, badChars[i]) {
          break
        }

        strLen = len(cleanStr) - len(badChars[i])

        cleanStr = cleanStr[0:strLen]

        cycleWhereStringRemoved = k

      }

      if len(cleanStr) == 0 {
        goto Done
      }
    }

    strLen = len(cleanStr)

    if strLen == 0 {
      goto Done
    }

    if k-cycleWhereStringRemoved > 3 || k > 1000000 {
      goto Done
    }

  }

Done:

  return cleanStr, strLen
}

// StrLeftJustify - Creates a new string with a total length equal to input parameter
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the output string and
// spaces are padded to the right in order to create a string with total length of
// 'fieldLen'.
//
// Example:
//
//	fieldLen = 15
//	strToJustify 	= "Hello World"
//	Returned String = "Hello World    "
//	String Index    =  012345648901234
func (sops StrOps) StrLeftJustify(strToJustify string, fieldLen int) (string, error) {

  ePrefix := "StrOps.StrLeftJustify() "

  if sops.IsEmptyOrWhiteSpace(strToJustify) {
    return strToJustify,
      errors.New(ePrefix +
        "Error: Input parameter 'strToJustify' is All White Space or an EMPTY String!")
  }

  strLen := len(strToJustify)

  if fieldLen == strLen {
    return strToJustify, nil
  }

  if fieldLen < strLen {
    return strToJustify,
      fmt.Errorf(ePrefix+
        "Error: Length of string to left justify is '%v'. "+
        "'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
  }

  rightPadLen := fieldLen - strLen

  rightPadStr := strings.Repeat(" ", rightPadLen)

  return strToJustify + rightPadStr, nil

}

// StrPadLeftToCenter - Returns a blank string which allows centering of the target
// string in a fixed length field.
//
// Example:
//
//	Assume that total field length ('fieldlen') is 70. Assume that the string to Center
//	('strToCenter') is 10-characters. In order to center a 10-character string in a
//	70-character field, 30-space characters would need to be positioned on each side
//	of the string to center. This method only returns the left segment, or a string
//	consisting of 30-spaces.
func (sops StrOps) StrPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

  ePrefix := "StrOps.StrPadLeftToCenter() "

  if sops.IsEmptyOrWhiteSpace(strToCenter) {
    return strToCenter,
      errors.New(ePrefix +
        "Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!")
  }

  sLen := sops.StrGetRuneCnt(strToCenter)

  if sLen > fieldLen {
    return "",
      errors.New(ePrefix +
        "Error: Input Parameter String To Center ('strToCenter') " +
        "is longer than Field Length")
  }

  if sLen == fieldLen {
    return "", nil
  }

  margin := (fieldLen - sLen) / 2

  return strings.Repeat(" ", margin), nil
}

// StrRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
//
// Example:
//
//	If the total field length ('fieldLen') is specified as 50-characters and the
//	length of string to justify ('strToJustify') is 20-characters, then this method
//	would return a string consisting of 30-space characters plus the 'strToJustify'.
func (sops StrOps) StrRightJustify(strToJustify string, fieldLen int) (string, error) {

  ePrefix := "StrOps.StrRightJustify() "

  if sops.IsEmptyOrWhiteSpace(strToJustify) {
    return strToJustify,
      errors.New(ePrefix +
        "Error: Input parameter 'strToJustify' is All White Space or an EMPTY String!")
  }

  strLen := len(strToJustify)

  if fieldLen == strLen {
    return strToJustify, nil
  }

  if fieldLen < strLen {
    return strToJustify,
      fmt.Errorf(ePrefix+
        "Error: Length of string to "+
        "right justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'",
        strLen, fieldLen)
  }

  // fieldLen must be greater than strLen
  lefPadCnt := fieldLen - strLen

  leftPadStr := strings.Repeat(" ", lefPadCnt)

  return leftPadStr + strToJustify, nil
}

// SwapRune - Swaps all instances of 'oldRune' character with 'newRune'
// character in input parameter target string ('targetStr').
func (sops StrOps) SwapRune(targetStr string, oldRune rune, newRune rune) (string, error) {

  if targetStr == "" {
    return targetStr, nil
  }

  rStr := []rune(targetStr)

  lrStr := len(rStr)

  for i := 0; i < lrStr; i++ {
    if rStr[i] == oldRune {
      rStr[i] = newRune
    }
  }

  return string(rStr), nil
}

// TrimMultipleChars- Performs the following operations on strings:
//
// 	1. Trims Right and Left ends of 'targetStr' for all instances of 'trimChar'
// 	2. Within the interior of a string, multiple instances of 'trimChar' are reduced
//	   to a single instance.
//
// Example:
//
//	targetStr = "       Hello          World        "
//	trimChar  = ' ' (One Space)
//	returned string (rStr) = "Hello World"
func (sops StrOps) TrimMultipleChars(
  targetStr string,
  trimChar rune) (rStr string, err error) {

  ePrefix := "StrOps.TrimMultipleChars() "

  rStr = ""
  err = nil

  if targetStr == "" {
    err = errors.New(ePrefix + "Error: Input parameter 'targetStr' is an EMPTY STRING!")
    return rStr, err
  }

  if trimChar == 0 {
    err = errors.New(ePrefix + "Error: Input parameter 'trimChar' is ZERO!")
    return rStr, err
  }

  fStr := []rune(targetStr)
  lenTargetStr := len(fStr)
  outputStr := make([]rune, lenTargetStr)
  lenTargetStr--
  idx := lenTargetStr
  foundFirstChar := false

  for i := lenTargetStr; i >= 0; i-- {

    if !foundFirstChar && fStr[i] == trimChar {
      continue
    }

    if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
      continue
    }

    if i == 0 && fStr[i] == trimChar {
      continue
    }

    foundFirstChar = true
    outputStr[idx] = fStr[i]
    idx--
  }

  if idx != lenTargetStr {
    idx++
  }

  if outputStr[idx] == trimChar {
    idx++
  }

  rStr = string(outputStr[idx:])
  err = nil

  return rStr, err
}

// TrimStringEnds - Removes all instances of input
// parameter 'trimChar' from the beginning and end
// of input parameter string 'targetStr'.
func (sops StrOps) TrimStringEnds(
  targetStr string,
  trimChar rune) (rStr string, err error) {

  ePrefix := "StrOps.TrimStringEnds() "
  rStr = ""
  err = nil

  targetStrLen := len(targetStr)

  if targetStrLen == 0 {
    err = errors.New(ePrefix + "Error: Input parameter 'targetStr' is an EMPTY STRING!")
    return rStr, err
  }

  if trimChar == 0 {
    err = errors.New(ePrefix + "Error: Input parameter 'trimChar' is ZERO!")
    return rStr, err
  }

  foundGoodChar := false
  firstIdx := 0

  for !foundGoodChar {

    if rune(targetStr[firstIdx]) == trimChar {
      firstIdx++
    } else {
      foundGoodChar = true
    }

    if firstIdx >= targetStrLen {
      rStr = ""
      return rStr, err
    }
  }

  if firstIdx >= targetStrLen {

    rStr = targetStr

    return rStr, err
  }

  foundGoodChar = false
  lastIdx := targetStrLen - 1

  for !foundGoodChar {

    if rune(targetStr[lastIdx]) == trimChar {
      lastIdx--
    } else {
      foundGoodChar = true
    }
  }

  if lastIdx < 0 {
    rStr = targetStr[firstIdx:]
    return rStr, err
  }

  lastIdx++
  rStr = targetStr[firstIdx:lastIdx]

  return rStr, err
}

// UpperCaseFirstLetter - Finds the first alphabetic character in a string
// (a-z A-Z) and converts it to upper case.
func (sops StrOps) UpperCaseFirstLetter(str string) string {

  if len(str) == 0 {
    return str
  }

  runesStr := []rune(str)

  for i := 0; i < len(runesStr); i++ {

    // Skip leading spaces
    if runesStr[i] == ' ' {
      continue
    }

    // Find the first alphabetic character and
    // convert to upper case.
    if runesStr[i] >= 'a' && runesStr[i] <= 'z' {

      runesStr[i] -= 32
      break

    } else if runesStr[i] >= 'A' && runesStr[i] <= 'Z' {
      break
    }

  }

  return string(runesStr)
}

// Write - Implements the io.Writer interface.
// Write writes len(p) bytes from p to the underlying
// data stream. In this case the underlying data stream
// is private member variable string, 'StrOps.stringData'.
//
// Receives a byte array 'p' and writes the contents to
// a string, private structure data element 'StrOps.stringData'.
//
// 'StrOps.stringData' can be accessed through 'Getter' and
// 'Setter' methods, 'GetStringData()' and 'SetStringData()'.
func (sops *StrOps) Write(p []byte) (n int, err error) {

  ePrefix := "StrOps.Write() "
  n = 0
  err = nil

  sops.stringDataMutex.Lock()

  if sops.cntBytesWritten == 0 {
    sops.stringData = ""
  }

  n = len(p)

  if n == 0 {

    sops.cntBytesWritten = 0

    sops.stringDataMutex.Unlock()

    err = fmt.Errorf(ePrefix + "Error: Input byte array 'p' is ZERO LENGHT!")

    return n, err
  }

  sops.stringDataMutex.Unlock()

  w := strings.Builder{}
  w.Grow(n + 5)
  cnt := 0

  endOfString := false

  for i := 0; i < n; i++ {

    if p[i] == 0 {
      endOfString = true
      break
    }

    w.WriteByte(p[i])
    cnt++
  }

  n = cnt

  sops.stringDataMutex.Lock()

  sops.stringData += w.String()

  if endOfString {
    sops.cntBytesWritten = 0
  } else {
    sops.cntBytesWritten += uint64(n)
  }

  sops.stringDataMutex.Unlock()

  return n, err
}
