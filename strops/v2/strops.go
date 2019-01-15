package strops

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"
)

/*
	'strops.go' is located in source code repository:

			https://github.com/MikeAustin71/stringopsgo.git


*/

// StrOps - encapsulates a collection of methods used to manage string
// operations.
//
// Most of the utility offered by this type is provided through its
// associated methods. However, given that two data elements, 'StrIn'
// and 'StrOut' are provided, the structure may be used as a data
// transport object (dto) containing two strings.
//
// Be advised that this type, 'StrOps', implements the io.Reader and io.Writer
// interfaces. All io.Reader and io.Writer operations utilize the string data
// element, 'StrOps.StrOut'.
//
type StrOps struct {
	StrIn           string
	StrOut          string
	stringData      string
	stringDataMutex sync.Mutex
	cntBytesRead    uint64
	cntBytesWritten uint64
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

// CopyIn - Copies string information from another StrOps
// instance passed as an input parameter to the current
// StrOps instance.
//
func (sops *StrOps) CopyIn(strops2 *StrOps) {

	sops.StrIn = strops2.StrIn
	sops.StrOut = strops2.StrIn

}

// CopyOut - Creates a 'deep' copy of the current
// StrOps instance and returns a pointer to a
// new instance containing that copied information.
func (sops *StrOps) CopyOut() *StrOps {

	strops2 := StrOps{}
	strops2.StrIn = sops.StrIn
	strops2.StrOut = sops.StrOut

	return &strops2
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

// GetCountBytesRead - Returns private member variable
// StrOps.cntBytesRead which holds the number of bytes
// accumulated by the last Read operation executed through
// method StroOps.Read().
//
func (sops *StrOps) GetCountBytesRead() uint64 {
	return sops.cntBytesRead
}

// GetReader() Returns an io.Reader which will read the internal
// data element StrOps.StrOut.
//
func (sops *StrOps) GetReader() io.Reader {

	return strings.NewReader(sops.StrOut)
}

// GetSoftwareVersion - Returns the major software version for StrOps.
func (sops StrOps) GetSoftwareVersion() string {
	return "2.0.0"
}

// GetStringData - Returns the current value of internal
// member string, StrOps.stringData
//
func (sops *StrOps) GetStringData() string {
	var output string

	sops.stringDataMutex.Lock()
	output = sops.stringData
	sops.stringDataMutex.Unlock()

	return output
}

// GetValidBytes - Receives an array of 'targetBytes' which will be examined to determine
// the validity of individual bytes or characters. Each character (byte) in input array
// 'targetBytes' will be compared to input parameter 'validBytes', another array of bytes.
// If a character in 'targetBytes' also exists in 'validBytes' it will be considered valid
// and included in the returned array of bytes.
//
// Input Parameters
// ================
// targetBytes		[] byte		- An array of characters (bytes) which will be examined
//                            for valid characters. The list of valid characters is
//                            found in input parameter 'validBytes'. Valid characters
//                            in targetBytes will be returned by this method as an
//                            array of bytes. Invalid characters will be discarded.
//
//	validBytes		[] byte		- An array of bytes containing valid characters. If a character
//                            (byte) in 'targetBytes' is also present in 'validBytes' it will
//                            be classified as 'valid' and included in the returned array of
//                            bytes. Invalid characters will be discarded.
//
// Return Values
// =============
//
//	[] byte					- An array of bytes which contains bytes that are present in both 'targetBytes'
//										and 'validBytes'. Note: If all characters in 'targetBytes' are classified as
//                    'invalid', the returned array of bytes will be a zero length array.
//
//	error						- If the method completes successfully this value is 'nil'. If an error is
//                    encountered this value will contain the error message. Examples of possible
//                    errors include a zero length 'targetBytes array or 'validBytes' array.
//
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
// Input Parameters
// ================
// targetRunes		[] rune		- An array of characters (runes) which will be examined
//                            for valid characters. The list of valid characters is
//                            found in input parameter 'validRunes'. Valid characters
//                            in targetRunes will be returned by this method as an
//                            array of runes. Invalid characters will be discarded.
//
//	validRunes		[] rune		- An array of runes containing valid characters. If a character
//                            (rune) in targetRunes is also present in 'validRunes' it will
//                            be classified as 'valid' and included in the returned array of
//                            runes. Invalid characters will be discarded.
//
// Return Values
// =============
//
//	[] rune					- An array of runes which contains runes that are present in 'targetRunes' and
//                    'validRunes'. Note: If all characters in 'targetRunes' are classified as
//                    'invalid', the returned array of runes will be a zero length array.
//
// error						- If the method completes successfully this value is 'nil'. If an error is
//                    encountered this value will contain the error message. Examples of possible
//                    errors include a zero length 'targetRunes array or 'validRunes' array.
//
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
// Input Parameter
// ===============
//
// targetStr	string	-	The string which will be screened for valid characters.
//
// validRunes	[] rune	- An array of type rune containing valid characters. Characters
//                      which exist in both 'targetStr' and 'validRunes' will be
//                      returned as a new string. Invalid characters are discarded.
//
// Return Values
// =============
//
// string					- This string will be returned containing valid characters extracted
//                  from 'targetStr'. A character is considered valid if it exists in
//                  both 'targetStr' and 'validRunes'. Invalid characters are discarded.
//                  This means that if no valid characters are identified, a zero length
//                  string will be returned.
//
// error					- If the method completes successfully this value is 'nil'. If an error is
//                  encountered this value will contain the error message. Examples of possible
//                  errors include a zero length 'targetStr' (string) or a zero length
//                  'validRunes' array.
//
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

	if n == 0 {
		err = errors.New(ePrefix + "Error: Input byte array 'p' is zero length!")
		return 0, err
	}

	sops.stringDataMutex.Lock()

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

// ReplaceBytes	- Replaces characters in a target array of bytes ([]bytes) with those specified in
// a two dimensional slice of bytes.
//
// Input Parameters
// ================
//
// targetBytes	[]byte				- The byte array which will be examined. If characters ('bytes') eligible
// 															for replacement are identified by replacementBytes[i][0] they will be
// 															replaced by the character specified in replacementBytes[i][1].
//
// replacementBytes	[][]byte 	- A two dimensional slice of type byte. Element [i][0] contains the target
//                             	character to locate in 'targetBytes'. Element[i][1] contains the replacement
//                             	character which will replace the target character in 'targetBytes'. If
//                             	the replacement character element [i][1] is a zero value, the target character
//                             	will not be replaced. Instead, it will be eliminated or removed from
//                             	the returned byte array ([]byte).
//
// Return Values
// =============
// []byte						- The returned byte array containing the characters and replaced characters
//                    from the original 'targetBytes' array.
//
// error						- If the method completes successfully this value is 'nil'. If an error is
//                    encountered this value will contain the error message. Examples of possible
//                    errors include a zero length targetBytes[] array or replacementBytes[][] array.
// 										In addition, if any of the replacementBytes[][x] 2nd dimension elements have
//                    a length less than two, an error will be returned.
//
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

// ReplaceRunes - Replaces characters in a target array of runes ([]rune) with those specified in
// a two dimensional slice of runes, 'replacementRunes[][]'.
//
// Input Parameters
// ================
//
// targetRunes			[]rune		- The rune array which will be examined. If target characters ('runes')
// 															eligible for replacement are identified by replacementRunes[i][0], they
// 															will be replaced by the character specified in replacementRunes[i][1].
//
// replacementRunes	[][]rune 	- A two dimensional slice of type 'rune'. Element [i][0] contains the target
//                             	character to locate in 'targetRunes'. Element[i][1] contains the
//                             	replacement character which will replace the target character in
//                             	'targetRunes'. If the replacement character element [i][1] is a zero value,
//                             	the	target character will not be replaced. Instead, it will be eliminated
//                             	or removed from the returned rune array ([]rune).
//
// Return Values
// =============
// []rune						- The returned rune array containing the characters and replaced characters
//                    from the original 'targetRunes' array.
//
// error						- If the method completes successfully this value is 'nil'. If an error is
//                    encountered this value will contain the error message. Examples of possible
//                    errors include a zero length 'targetRunes' array or 'replacementRunes' array.
//                    In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//                    a length less than two, an error will be returned.
//
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
// Input Parameters
// ================
//
// targetStr				string		- The string which will be examined. If target string characters
// 															eligible for replacement are identified by replacementRunes[i][0], they
// 															will be replaced by the character specified in replacementRunes[i][1].
//
// replacementRunes	[][]rune 	- A two dimensional slice of type 'rune'. Element [i][0] contains the target
//                             	character to locate in 'targetStr'. Element[i][1] contains the
//                             	replacement character which will replace the target character in
//                             	'targetStr'. If the replacement character element [i][1] is a zero value,
//                             	the	target character will not be replaced. Instead, it will be eliminated
//                             	or removed from the returned string.
//
// Return Values
// =============
// string						- The returned string containing the characters and replaced characters
//                    from the original target string, ('targetStr').
//
// error						- If the method completes successfully this value is 'nil'. If an error is
//                    encountered this value will contain the error message. Examples of possible
//                    errors include a zero length 'targetStr' or 'replacementRunes[][]' array.
// 										In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//                    a length less than two, an error will be returned.
//
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
// It use restricted primarily to debugging operations.
//
// This method sets StrOps.cntBytesRead equal to zero.
//
func (sops *StrOps) ResetBytesReadCounter() {
	sops.cntBytesRead = 0
}

// SetStringData - Sets the value of internal
// member string variable, StrOps.stringData.
//
func (sops *StrOps) SetStringData(str string) {
	sops.stringDataMutex.Lock()
	sops.stringData = str
	sops.stringDataMutex.Unlock()
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

// Write - Implements the io.Writer interface.
// Write writes len(p) bytes from p to the underlying
// data stream.
//
// Receives a byte array 'p' and writes the contents to
// a string, internal structure data element 'StrOps.stringData'.
//
// 'StrOps.stringData' can be accessed through 'Getter' and
// 'Setter' methods, 'GetStringData()' and 'SetStringData()'.
//
func (sops *StrOps) Write(p []byte) (n int, err error) {

	n = 0
	err = nil

	n = len(p)

	if n == 0 {
		return n, err
	}

	w := strings.Builder{}
	w.Grow(n)
	cnt := 0

	for i := 0; i < n; i++ {

		if p[i] == 0 {
			break
		}

		w.WriteByte(p[i])
		cnt++
	}

	n = cnt

	sops.stringDataMutex.Lock()

	sops.stringData = w.String()

	sops.stringDataMutex.Unlock()

	return n, err
}
