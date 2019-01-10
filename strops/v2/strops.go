package strops

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

/*
	'strops.go' is located in source code repository:

			https://github.com/MikeAustin71/stringopsgo.git


*/

// StrOps - encapsulates a collection of
// methods used to manipulate strings
type StrOps struct {
	StrIn  string
	StrOut string
}

// BreakTextAtLineLength - Breaks string text into lines. Takes a string and inserts a
// line delimiter character (a.k.a 'rune') at the specified line length ('lineLength').
//
// Input Parameters
// ================
//
// targetStr	string	- The string which will be parsed into text lines.
//
// lineLength	int			- The maximum length of each line.
//
// Note: If the caller specifies a line length of 50, new characters may be placed in the
// 51st character position depending upon the word breaks.
//
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

// DoesLastCharExist - returns true if the last character (rune) of
// input string 'testStr' is equal to input parameter 'lastChar' which
// is of type 'rune'.
//
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

// FindFirstNonSpaceChar - Returns the string index of the first non-space character in
// a string segment. The string to be searched is input parameter 'targetStr'. The string
// segment which will be searched from left to right in 'targetStr' is defined by the
// starting index ('startIndex') and the ending index ('endIndex').
//
// Searching from left to right, this method identifies the first non-space character
// (any character that is NOT a space ' ') in the target string segment and returns
// the index associated with that non-space character.
//
// Return Values
// =============
//
// The method returns the index of the first non-space character in the target string
// segment using a left to right search. If the entire string consists of space characters,
// this method returns a value of -1.
//
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
//
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
//
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
// Examples
// ========
//
//	Example (1)
// 				In the text string segment:
//
//       				"The cow jumped over the moon."
//
// 						The last word would be defined as "moon."
//
//	Example (2)
//				In the text string segment:
//							"  somewhere over the rainbow  "
//
//							The last word would be defined as "rainbow"
//
// ------------------------------------------------------------------------
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
// Input Parameters
// ================
//
// targetStr			string	- The string containing the string segment which
//                      		will be searched to identify the last word
//                      		in the string segment.
//
// startIndex			int			- The index marking the beginning of the string
//                      		segment in 'targetStr'.
//
// endIndex				int			- The index marking the end of the string segment
//                      		in 'targetStr'.
//
// Return Values
// =============
//
//	beginWrdIdx		int			-	The index marking the beginning of the last word
//                          in the string segment identified by input parameters
//                          'startIndex' and 'endIndex'. If the string segment
//                          consists of all spaces or is empty, this value is
//                          set to -1.
//
//	endWrdIdx			int			- The index marking the end of the last word in the
//                          string segment identified by input parameters 'startIndex'
//                          and 'endIndex'. If the string segment consists of all
//                          spaces or is empty, this value is set to -1.
//
//	isAllOneWord	bool		- If the string segment identified by input parameters
//                          'startIndex' and 'endIndex' consists entirely of non-space
//                          characters (characters other than ' '), this value is set
//                          to 'true'.
//
//	isAllSpaces		bool		- If the string segment identified by input parameters
//                          'startIndex' and 'endIndex' consists entirely of space
//                          characters (character = ' '), this value is set to 'true'.
//
//  err						error		- If targetStr is empty or if startIndex or endIndex is invalid,
//                          an error is returned. If the method completes successfully,
//                          err = nil.
//
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
// Return Value
// ============
// The return value is an array of integers. If no match is found the return
// value is 'nil'.  If regular expression is successfully matched, the match
// will be located at targetStr[loc[0]:loc[1]]. Again, a return value of 'nil'
// signals that no match was found.
//
func (sops StrOps) FindRegExIndex(targetStr string, regex string) []int {

	re := regexp.MustCompile(regex)

	return re.FindStringIndex(targetStr)

}

// GetSoftwareVersion - Returns the major software version for StrOps.
func (sops StrOps) GetSoftwareVersion() string {
	return "2.0.0"
}

// IsEmptyOrWhiteSpace - If a string is zero length or consists solely of
// white space (contiguous spaces), this method will return 'true'.
//
// Otherwise, a value of false is returned.
//
func (sops StrOps) IsEmptyOrWhiteSpace(targetStr string) bool {

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}

// MakeSingleCharString - Creates a string of length 'strLen' consisting of
// a single character passed through input parameter, 'charRune' as type
// 'rune'.
//
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
	b.Grow(32)

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

// ReplaceMultipleStrs - Replaces all instances of string replaceArray[i][0] with
// replacement string from replaceArray[i][1] in 'targetStr'.
//
// Note: The original 'targetStr' is NOT altered.
//
// Input parameter 'replaceArray' should be passed as multi-dimensional slices.
// If the length of the 'replaceArray' second dimension is less than '2', an
// error will be returned.
//
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
//
func (sops StrOps) ReplaceNewLines(targetStr string, replacement string) string {

	return strings.Replace(targetStr, "\n", replacement, -1)
}

// StrCenterInStrLeft - returns a string which includes
// a left pad blank string plus the original string. It
// does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively
// center the original string is a field of specified length.
func (sops StrOps) StrCenterInStrLeft(strToCenter string, fieldLen int) (string, error) {

	pad, err := sops.StrPadLeftToCenter(strToCenter, fieldLen)

	if err != nil {
		return "", errors.New("StrOps:StrCenterInStrLeft() - " + err.Error())
	}

	return pad + strToCenter, nil

}

// StrCenterInStr - returns a string which includes
// a left pad blank string plus the original string,
// plus a right pad blank string.
//
// The complete string will effectively center the
// original string is a field of specified length.
func (sops StrOps) StrCenterInStr(strToCenter string, fieldLen int) (string, error) {

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter, fmt.Errorf("'fieldLen' = '%v' strToCenter Length= '%v'. "+
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

func (sops StrOps) StrLeftJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("StrLeftJustify() Error: Length of string to left justify is '%v'. "+
				"'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	rightPadLen := fieldLen - strLen

	rightPadStr := strings.Repeat(" ", rightPadLen)

	return strToJustify + rightPadStr, nil

}

// StrRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
func (sops StrOps) StrRightJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("StrRightJustify() Error: Length of string to "+
				"right justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)

	return leftPadStr + strToJustify, nil
}

// StrPadLeftToCenter - Returns a blank string
// which allows centering of the target string
// in a fixed length field.
func (sops StrOps) StrPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

	sLen := sops.StrGetRuneCnt(strToCenter)

	if sLen > fieldLen {
		return "", errors.New("StrOps:StrPadLeftToCenter() - String To Center is longer than Field Length")
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// StrGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (sops StrOps) StrGetRuneCnt(targetStr string) int {
	return utf8.RuneCountInString(targetStr)
}

// StrGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (sops StrOps) StrGetCharCnt(targetStr string) int {
	return len([]rune(targetStr))
}

// TrimMultipleChars- Performs the following operations on strings:
// 1. Trims Right and Left for all instances of 'trimChar'
// 2. Within the interior of a string, multiple instances
// 		of 'trimChar' are reduce to a single instance.
//
// Example:
//
// targetStr = "       Hello          World        "
// trimChar  = ' ' (One Space)
// returned string (rStr) = "Hello World"
//
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
//
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

// SwapRune - Swaps all instances of 'oldRune' character with 'newRune'
// character.
func (sops StrOps) SwapRune(currentStr string, oldRune rune, newRune rune) (string, error) {

	if currentStr == "" {
		return currentStr, nil
	}

	rStr := []rune(currentStr)

	lrStr := len(rStr)

	for i := 0; i < lrStr; i++ {
		if rStr[i] == oldRune {
			rStr[i] = newRune
		}
	}

	return string(rStr), nil
}
