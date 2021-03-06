// String operations provides string management utilities designed to
// perform a variety of string operations including string centering,
// justification, multiple replacements numeric and date field extraction
// as well as implementing the io.Reader and io.Writer interfaces.
//
//
package strops

import (
	"io"
	"strings"
	"sync"
)

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
	stringDataMutex *sync.Mutex // Used internally to ensure thread safe operations
	cntBytesRead    uint64      // Used internally to track Bytes Read by StrOps.Read()
	cntBytesWritten uint64      // Used internally to track Bytes Written by StrOps.Write()
}

// BreakTextAtLineLength - Breaks string text into lines. Takes a
// string and inserts a line delimiter character (a.k.a 'rune') at
// the specified line length ('lineLength').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be parsed into text lines. If
//       'targetStr' is an empty string an error will be returned.
//        If 'targetStr' consists entirely of white space, this
//        method will return a string consisting of a new-line
//        character and an error value of 'nil'.
//
//
//  lineLength          int
//     - The maximum length of each line.
//
//
//  lineDelimiter       rune
//     - The line delimiter character which will be inserted at the
//       end of a line break.
//
//
//  Note: If the caller specifies a line length of 50, the line delimiter
//  character may be placed in the 51st character position depending upon
//  the word breaks.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string
//       parameter will contain the text with line breaks delimited
//       by the input parameter 'lineDelimiter'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
func (sops *StrOps) BreakTextAtLineLength(
	targetStr string,
	lineLength int,
	lineDelimiter rune) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix := "StrOps.BreakTextAtLineLength() "

	sOpsAtom := strOpsAtom{}

	return sOpsAtom.breakTextAtLineLength(
		targetStr,
		lineLength,
		lineDelimiter,
		ePrefix)
}

// ConvertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// Examples of non-printable characters are '\n', '\t' or 0x06
// (Acknowledge). These example characters would be translated into
// printable string characters as: "\\n", "\\t" and "[ACK]".
//
// Space characters are typically translated as " ". However, if
// the input parameter 'convertSpace' is set to 'true' then all
// spaces are converted to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters will be converted to
//       printable characters.
//
//  convertSpace        bool
//     - Space or white space characters (0x20) are by default
//       translated as " ". However, if this parameter is set to
//       'true', space characters will be converted to "[SPACE]".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  printableChars      string
//     - This returned string is identical to input parameter
//       'nonPrintableChars' with the exception that non-printable
//       characters are translated into printable characters.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello world!\n"
//  testRunes := []rune(testStr)
//
//  actualStr :=
//    StrOps{}.NewPtr().
//      ConvertNonPrintableChars(testRunes, true)
//
//  ----------------------------------------------------
//  'actualStr' is now equal to:
//     "Hello[SPACE]world!\\n"
//
func (sops *StrOps) ConvertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool) (
	printableChars string) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix := "StrOps.ConvertNonPrintableChars() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.convertNonPrintableChars(
		nonPrintableChars,
		convertSpace,
		ePrefix)
}

// CopyIn - Copies string information from another StrOps
// instance passed as an input parameter to the current
// StrOps instance.
func (sops *StrOps) CopyIn(strops2 *StrOps) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	if strops2 == nil {
		return
	}

	sOpsAtom := strOpsAtom{}

	_ = sOpsAtom.copyIn(
		sops,
		strops2,
		"")
}

// CopyOut - Creates a 'deep' copy of the current
// StrOps instance and returns a pointer to a
// new instance containing that copied information.
func (sops *StrOps) CopyOut() *StrOps {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsAtom := strOpsAtom{}

	newStrOps,
		_ := sOpsAtom.copyOut(
		sops,
		"")

	return newStrOps
}

// DoesLastCharExist - returns true if the last character (rune) of
// input string 'testStr' is equal to input parameter 'lastChar' which
// is of type 'rune'.
func (sops StrOps) DoesLastCharExist(
	testStr string,
	lastChar rune) bool {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.doesLastCharExist(
		testStr,
		lastChar)
}

// ExtractDataField - Extracts a data field string from a larger target string ('targetStr').
// The target string is searched for a data field. If the 'leadingKeyWordDelimiter' parameter
// is populated, the data field MUST contain this leading key word, otherwise an empty data field
// is returned.
//
// If 'leadingKeyWordDelimiter' is an empty string, the search for the data field will begin at
// 'targetStr' index, 'startIdx'.
//
// The returned data field must occur in 'targetStr' prior to a comment or End-Of-Line character.
//
// The extracted data field MUST be preceded by one of the characters specified in input
// parameter, 'leadingFieldSeparators'. In addition, the data field must be immediately
// followed by one of the characters in input parameter 'trailingFieldSeparators' or a comment
// or an End-Of-Line character.
//
// ------------------------------------------------------------------------
//
// Input Values
//
//  targetStr                  string
//     - The target string from which the data field will be extracted.
//
//
//  leadingKeyWordDelimiters   []string
//     - Data fields are often preceded by field names or field designators.
//       The 'leadingKeyWordDelimiters' parameter is a string array
//       containing 'Key Word Delimiters'. A Key Word Delimiter may be
//       a Key Word string or a character which identifies and immediately
//       precedes the data field. If multiple Key Word Delimiters exist
//       in 'targetStr' the first instance of a key word in targetStr'
//       will be designated as the Key Word Delimiter.
//
//       If this parameter is populated, the search for a data field
//       will begin immediately after the first located Key Word
//       Delimiter string. If none of Key Words in this string array
//       are located in 'targetStr', an empty string will be returned
//       for data field. If this parameter is populated, at least one
//       of the Key Words MUST exist in 'targetStr' before a data field
//       will be extracted and returned.
//
//       If this parameter is an empty string array, the search for a
//       data field will begin at the string index designated by
//       parameter, 'startIdx'.
//
//
//  startIdx                   int
//     - The string index in parameter 'targetStr' from which the search for
//       a data field will begin. Note that the starting index will be adjusted
//       according to the existence of a Key Word Delimiter as explained
//       above.
//
//
//  leadingFieldSeparators     []string
//     - An array of characters or groups of characters which delimit the
//       leading edge of the data field.
//
//
//  trailingFieldSeparators    []string
//     - An array of characters or groups of characters which delimit the
//       end of a data field.
//
//
//  commentDelimiters        []string
//     - Comments effectively terminate the search for a data field. This
//       array stores comment characters or phrases which signal the beginning
//       of a comment.
//
//
//  endOfLineDelimiters        []string
//     - Those characters or groups of characters which mark the end of a line.
//       Generally this includes characters like 'new line' or 'carriage return'.
//       End of line characters will terminate the search for a data field.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DataFieldProfileDto - If successful, this method returns a structure containing
//                        characteristics describing the extracted data field.
//
//    type DataFieldProfileDto struct {
//       TargetStr                      string //  The string from which the data field is extracted.
//       TargetStrLength                int    //  Length of 'TargetStr'
//       TargetStrStartIndex            int    //  The index with in 'TargetStr' from which the search for a data field was initiated.
//       TargetStrLastGoodIndex         int    //  Last valid index in target string which is less than the target string length and is NOT an 'End Of Field' or 'End Of Line' Delimiter.
//       LeadingKeyWordDelimiter        string //  The Leading Key Word Delimiter which is used to identify the beginning of the field search.
//       LeadingKeyWordDelimiterIndex   int    //  Index of the found Leading Key Word Delimiter.
//       DataFieldStr                   string //  The extracted data field string.
//       DataFieldIndex                 int    //  The index in 'TargetStr' where the data field begins.
//       DataFieldLength                int    //  The length of the extracted data field string.
//       DataFieldTrailingDelimiter     string //  The trailing character which marked the end of the data field. A zero value indicates end of string encountered.
//       DataFieldTrailingDelimiterType DataFieldTrailingDelimiterType // A constant or enumeration type used to describe the type of delimiter used to mark the end of a data field.
//       NextTargetStrIndex             int    //  The index in 'TargetStr' immediately following the extracted data field.
//       CommentDelimiter               string //  If a Comment Delimiter is detected it is stored here.
//       CommentDelimiterIndex          int    //  If a Comment Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
//       EndOfLineDelimiter             string //  If an End-Of-Line Delimiter is detected it is captured and stored here.
//       EndOfLineDelimiterIndex        int    //  If an End-Of-Line Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
//     }
//
//   error
//     - If the method completes successfully and no errors are encountered
//       this return value is set to 'nil'. Otherwise, if errors are encountered
//       this return value will contain an appropriate error message.
//
//       The most likely source of errors are invalid input parameters.
//       Input parameters 'targetStr', 'startIdx', 'leadingFieldSeparators',
//       'trailingFieldSeparators' and 'endOfStringDelimiters' are required input
//       parameters and must be populated with valid data.
//
//       If an error message is returned, the input parameter
//       'ePrefix' will be inserted or prefixed at the beginning
//       of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_ExtractDataField_01() "
//  endOfLineDelimiters := []string{"\n"}
//  commentDelimiters := []string{"#"}
//  leadingFieldDelimiters := []string{
//  "\t",
//  "\r",
//  "\f",
//  "\v",
//  " "}
//
//  trailingFieldDelimiters := []string{
//  "\t",
//  "\r",
//  "\f",
//  "\v",
//  " "}
//
//  targetStr := " Zone:\t America/Chicago\t Link:\t US/Central\t\n"
//  startIdx := 0
//  leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
//
//  datDto,
//  err :=
//    StrOps{}.Ptr().
//        ExtractDataField(
//           targetStr,
//           leadingKeyWordDelimiters,
//           startIdx,
//           leadingFieldDelimiters,
//           trailingFieldDelimiters,
//           commentDelimiters,
//           endOfLineDelimiters,
//           ePrefix)
//
//  -----------------------------------------------
//  datDto.DataFieldStr is now equal to:
//          "America/Chicago"
//
//
func (sops *StrOps) ExtractDataField(
	targetStr string,
	leadingKeyWordDelimiters []string,
	startIdx int,
	leadingFieldSeparators []string,
	trailingFieldSeparators []string,
	commentDelimiters []string,
	endOfLineDelimiters []string,
	ePrefix string) (
	DataFieldProfileDto,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ExtractDataField() "

	sOpsAtom := strOpsAtom{}

	return sOpsAtom.extractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldSeparators,
		trailingFieldSeparators,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)
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
//  targetStr           string
//     - The target string to be searched for the first instance of
//       a number string. A number string is usually defined as a
//       string comprised of one or more consecutive numeric digits.
//       Additional parameters provided by this method will allow
//       the caller to insert specified non-numeric characters at
//       the beginning, end or interior of a number string.
//
//
//  startIdx            int
//     - The starting index in input parameter 'targetStr'
//       from which the search for a number string will be
//       initiated. This useful in extracting multiple number
//       strings form a single 'targetStr'.
//
//       If 'startIndex' is less than zero or if 'startIndex' exceeds
//       the last character index in 'targetStr', an error will be
//       returned.
//
//
//  keepLeadingChars    string
//     - This string contains non-numeric characters which will be
//       retained as a prefix to the final number string extracted
//       from the 'targetStr' parameter. To be included, these characters
//       must exist in 'targetStr' and must immediately precede the
//       first instance of a number string.
//
//       For example, if the target string is "Hello $123789 world" and
//       parameter 'keepLeadingChars' includes the USA currency character,
//       '$', the returned number string would be '$123789'.  If no currency
//       character was included in 'keepLeadingChars', the returned number
//       string would be '123789'. It is worth noting that if the target
//       string was '$ 123789' and a currency symbol, '$', was included
//       in 'keepLeadingChars', the returned number string would still be
//       '123789' because 'keepLeadingChars' characters must immediately
//       precede the string of numeric digits in 'targetStr'.
//
//       Specifically, if the plus ('+') and minus ('-') sign are NOT
//       included in 'keepLeadingChars' those leading number signs will
//       never be included in the final number string.
//
//       Leading characters will not be repeated. If for some reason you
//       wanted to retain two leading currency symbols ("$$") it would be
//       necessary to include two currency characters in 'keepLeadingChars'.
//
//
//  keepInteriorChars   string
//     - This string contains non-numeric characters which will be retained
//       as valid characters within the final extracted number string. The
//       characters must exist withing the first instance of a number string
//       located in 'targetStr'. Such interior characters might include
//       thousands separators (commas) or decimal points (periods).
//
//       For example, if a comma and a period are included in 'keepInteriorChars'
//       and the target string is "Hello word 123,456,789.25 !", the returned
//       number string would be "123,456,789.25".  If the comma character was
//       NOT included in the 'keepInteriorChars' string, the returned number
//       string would be '123', since the number string extraction parser
//       would break on the comma, a non-numeric digit.
//
//       'keepInteriorChars' will NOT allow multiple non-numeric characters
//       to exist within the interior of the final extracted number string.
//       Only single non-numeric characters are allowed within a number string.
//
//
//  keepTrailingChars   string
//     - This string contains non-numeric characters which should be retained
//       at the end of the final number string. By default, a non-numeric
//       character will mark the end of a number string. However, if the caller
//       elects to use parameter 'keepTrailingChars' to retain non-numeric
//       characters such as a trailing right-parenthesis, then those non-numeric
//       characters will be retained in the final extracted number string.
//
//       Trailing characters will not be repeated. If for some reason you
//       wanted to retain two closing parentheses symbols ("))") it would be
//       necessary to include closing parentheses characters in 'keepTrailingChars'.
//
//       It should be emphasized that 'keepTrailingChars' must immediately
//       follow the first instance of a number string in parameter, 'targetStr'.
//
//       Example #1:
//         Target String = "Hello world, (1234). Today is new day."
//         keepLeadingChars = "("
//         keepInteriorChars = ""
//         keepTrailingChars= ")"
//         Extracted Number String = "(1234)"
//
//       Example #2:
//         Target String = "Hello world, USA GDP growth is projected at 1.8%."
//         keepLeadingChars = ""
//         keepInteriorChars = "."
//         keepTrailingChars= "%"
//         Extracted Number String = "1.8%"
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrProfileDto    - If successful, this method will return a type 'NumStrProfileDto'
//                        populated with the extracted number string and additional profile
//                        information related to the extracted number string.
//
//     type NumStrProfileDto struct {
//
//           TargetStr            string   //  The original target string which is scanned for a
//                                         //    number string
//
//           TargetStrStartIndex  int      //  The starting index in 'TargetStr' from which the
//                                         //    number string search was initiated.
//
//           LeadingSignIndex     int      //  The string index of a leading sign in 'NumStr' below. If a
//                                         //    leading sign character is NOT present in 'NumStr' this
//                                         //    value is set to -1
//
//           LeadingSignChar      string   //  If a leading sign character (plus '+' or minus '-')
//                                         //    exists in data field 'NumStr' (below), it is stored
//                                         //    in this string.
//
//           FirstNumCharIndex    int      //  The index in 'TargetStr' (above) where the first character
//                                         //    of the extracted number string is located.
//
//           NextTargetStrIndex   int      //  The index of the next character in 'TargetStr' immediately
//                                         //    following the extracted number string.
//
//           NumStrLen            int      //  The length of the extracted number string.
//
//           NumStr               string   //  The number string extracted from 'TargetStr'.
//     }
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If 'startIndex' is less than zero or if 'startIndex'
//       exceeds the last character index in 'targetStr', an error
//       will be returned.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_ExtractNumericDigits_06() "
//  targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
//  startIndex := 0
//  keepLeadingChars := "$("
//  keepInteriorChars := ",."
//  keepTrailingChars := ")"
//
//
//  nStrDto,
//  err :=
//    StrOps{}.Ptr().
//         ExtractNumericDigits(
//             targetStr,
//             startIndex,
//             keepLeadingChars,
//             keepInteriorChars,
//             keepTrailingChars,
//             ePrefix)
//
//  ----------------------------------------
//
//  nStrDto.NumStr is now equal to:
//     "$(1,250,364.33)"
//
//
func (sops *StrOps) ExtractNumericDigits(
	targetStr string,
	startIndex int,
	keepLeadingChars string,
	keepInteriorChars string,
	keepTrailingChars string,
	ePrefix string) (
	NumStrProfileDto,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ExtractNumericDigits() "

	sOpsAtom := strOpsAtom{}

	return sOpsAtom.extractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)
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
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string to be searched for the first non-space character.
//
//
//  startIdx            int
//     - Since the search is forwards from left to right, this is
//       the starting index for the search.
//
//
//  endIdx              int
//     - Since the search is forwards from left to right, this is
//       the ending index for the search.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the index of the first non-space
//       character in the target string segment using a left
//       to right search. If the target string is an empty string
//       or consists of entirely of space characters, this method
//       returns a value of minus one (-1).
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (sops *StrOps) FindFirstNonSpaceChar(
	targetStr string,
	startIndex,
	endIndex int,
	ePrefix string) (
	int,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.FindFirstNonSpaceChar() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.findFirstNonSpaceChar(
		targetStr,
		startIndex,
		endIndex,
		ePrefix)
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
// If the last non-space character is found, that string index is returned. If the string
// segment consists entirely of space characters, the return value is -1.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string to be searched for the last non-space character.
//
//  startIdx            int
//     - Since the search is backwards from right to left, this is
//       the ending index for the search.
//
//
//  endIdx              int
//     - Since this is a backwards search from right to left, this
//       is actually the starting index for the search.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - The index of the last non-space character in input
//       parameter 'targetStr' within the range specified by
//       the staring and ending indexes.
//
//       If all the characters within the specified range are
//       space characters, this parameter returns a value of
//       minus one (-1).
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
func (sops *StrOps) FindLastNonSpaceChar(
	targetStr string,
	startIdx int,
	endIdx int,
	ePrefix string) (
	int,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.FindLastNonSpaceChar() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.findLastNonSpaceChar(
		targetStr,
		startIdx,
		endIdx,
		ePrefix)
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string segment to be searched for the last space
//       character resides in this string.
//
//  startIdx            int
//     - Defines the actual ending index for the string segment to
//       be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is input parameter 'endIdx' while
//       the ending point for the string segment is this 'startIdx'.
//
//
//  endIdx              int
//     - Defines the actual beginning index of the string segment
//       to be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is defined by this 'endIdx'
//       parameter while the ending point for the string segment is
//       marked by the input parameter, 'startIdx'.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - If this method completes successfully, this returned
//       integer value will constitute the string index of the last
//       space character in the string segment marked by input
//       parameters 'startIdx' and 'endIdx'.
//
//       If a space character is NOT found in the specified string
//       segment, a value of minus one (-1) is returned.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//                 012345678901234
//       testStr := "xx       Hxgltx"
//
//       lastSpaceIdx, err :=
//       StrOps{}.Ptr().FindLastSpace(
//       testStr,
//       0,
//       14,
//       ePrefix)
//
//  lastSpaceIdx is now equal to '8'
//
//
func (sops *StrOps) FindLastSpace(
	targetStr string,
	startIdx int,
	endIdx int,
	ePrefix string) (
	int,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.FindLastSpace() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.findLastSpace(
		targetStr,
		startIdx,
		endIdx,
		ePrefix)
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
//   Example (1)
//     In the text string segment:
//
//     "The cow jumped over the moon."
//
//     The last word would be defined as "moon."
//
//     Example (2)
//       In the text string segment:
//
//       "  somewhere over the rainbow  "
//
//       The last word would be defined as "rainbow"
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
//  targetStr           string
//     - The string containing the string segment which
//       will be searched to identify the last word
//       in the string segment.
//
//
//  startIndex          int
//     - The index marking the beginning of the string
//       segment in 'targetStr'.
//
//
//  endIndex       int
//     - The index marking the end of the string segment
//       in 'targetStr'.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  beginWrdIdx         int
//     - The index marking the beginning of the last word
//       in the string segment identified by input parameters
//       'startIndex' and 'endIndex'. If the string segment
//       consists of all spaces or is empty, this value is
//       set to -1.
//
//
//  endWrdIdx           int
//     - The index marking the end of the last word in the
//       string segment identified by input parameters 'startIndex'
//       and 'endIndex'. If the string segment consists of all
//       spaces or is empty, this value is set to -1.
//
//
//  isAllOneWord        bool
//     - If the string segment identified by input parameters
//       'startIndex' and 'endIndex' consists entirely of non-space
//       characters (characters other than ' '), this value is set
//       to 'true'.
//
//
//  isAllSpaces         bool
//     - If the string segment identified by input parameters
//       'startIndex' and 'endIndex' consists entirely of space
//       characters (character = ' '), this value is set to 'true'.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (sops *StrOps) FindLastWord(
	targetStr string,
	startIndex,
	endIndex int,
	ePrefix string) (beginWrdIdx,
	endWrdIdx int,
	isAllOneWord,
	isAllSpaces bool,
	err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.FindLastWord() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.findLastWord(
		targetStr,
		startIndex,
		endIndex,
		ePrefix)
}

// FindRegExIndex - returns a two-element slice of integers defining the location
// of the leftmost match in targetStr of the regular expression (regex).
//
// ------------------------------------------------------------------------
//
// Return Value
//
// The return value is an array of integers. If no match is found the return
// value is 'nil'.  If regular expression is successfully matched, the match
// will be located at targetStr[loc[0]:loc[1]]. Again, a return value of 'nil'
// signals that no match was found.
//
func (sops *StrOps) FindRegExIndex(
	targetStr string,
	regex string) []int {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.findRegExIndex(
		targetStr,
		regex)
}

// GetReader - Returns an io.Reader which will read the private
// member data element StrOps.stringData.
func (sops *StrOps) GetReader() io.Reader {
	var stringData string

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	stringData = sops.stringData

	return strings.NewReader(stringData)
}

// GetSoftwareVersion - Returns the software version for package 'strops'.
func (sops StrOps) GetSoftwareVersion() string {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	return "3.0.0"
}

// GetStringData - Returns the current value of internal
// member string, StrOps.stringData
func (sops *StrOps) GetStringData() string {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	var output string

	output = sops.stringData
	sops.cntBytesWritten = 0
	sops.cntBytesRead = 0

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
//  targetBytes         [] byte
//     - An array of characters (bytes) which will be examined
//       for valid characters. The list of valid characters is
//       found in input parameter 'validBytes'. Valid characters
//       in targetBytes will be returned by this method as an
//       array of bytes. Invalid characters will be discarded.
//
//
//  validBytes          [] byte
//     - An array of bytes containing valid characters. If a character
//       (byte) in 'targetBytes' is also present in 'validBytes' it will
//       be classified as 'valid' and included in the returned array of
//       bytes. Invalid characters will be discarded.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  [] byte
//     - An array of bytes which contains bytes that are present in both 'targetBytes'
//       and 'validBytes'. Note: If all characters in 'targetBytes' are classified as
//       'invalid', the returned array of bytes will be a zero length array.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    ePrefix := "TestStrOps_GetValidBytes_01() "
//    validBytes := []byte{'v', 'a', 'l', 'i', 'd'}
//    testBytes := []byte{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}
//
//    actualBytes, err := StrOps{}.Ptr().GetValidBytes(
//    testBytes,
//    validBytes,
//    ePrefix)
//
//    'actualBytes' is now equal to "valid"
//
func (sops *StrOps) GetValidBytes(
	targetBytes,
	validBytes []byte,
	ePrefix string) ([]byte, error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.GetValidBytes() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.getValidBytes(
		targetBytes,
		validBytes,
		ePrefix)
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
//  targetRunes         [] rune
//     - An array of characters (runes) which will be examined
//       for valid characters. The list of valid characters is
//       found in input parameter 'validRunes'. Valid characters
//       in targetRunes will be returned by this method as an
//       array of runes. Invalid characters will be discarded.
//
//
//  validRunes          [] rune
//     - An array of runes containing valid characters. If a character
//       (rune) in targetRunes is also present in 'validRunes' it will
//       be classified as 'valid' and included in the returned array of
//       runes. Invalid characters will be discarded.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  [] rune
//     - An array of runes which contains runes that are present in 'targetRunes' and
//       'validRunes'. Note: If all characters in 'targetRunes' are classified as
//       'invalid', the returned array of runes will be a zero length array.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     ePrefix := "TestStrOps_GetValidRunes_01() "
//
//     validRunes := []rune{'v', 'a', 'l', 'i', 'd'}
//
//     testRunes := []rune{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}
//
//     expected := "valid"
//
//     actualRunes, err := StrOps{}.Ptr().GetValidRunes(
//     testRunes,
//     validRunes,
//     ePrefix)
//
//     actualRunes is now equal to string(actualRunes) == "valid"
//
//
func (sops *StrOps) GetValidRunes(
	targetRunes []rune,
	validRunes []rune,
	ePrefix string) (
	[]rune,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.GetValidRunes() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.getValidRunes(
		targetRunes,
		validRunes,
		ePrefix)
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
//  targetStr           string
//     - The string which will be screened for valid characters.
//
//
//  validRunes []rune
//     - An array of type rune containing valid characters. Characters
//       which exist in both 'targetStr' and 'validRunes' will be
//       returned as a new string. Invalid characters are discarded.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This string will be returned containing valid characters extracted
//       from 'targetStr'. A character is considered valid if it exists in
//       both 'targetStr' and 'validRunes'. Invalid characters are discarded.
//       This means that if no valid characters are identified, a zero length
//       string will be returned.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    ePrefix := "TestStrOps_GetValidString_01() "
//
//    validRunes := []rune{'v', 'a', 'l', 'i', 'd'}
//
//    testStr := "xjvmRaJlZiFdS"
//
//    actualStr, err := StrOps{}.Ptr().GetValidString(
//                        testStr,
//                        validRunes,
//                        ePrefix)
//
//  'actualStr' is now equal to "valid"
//
//
func (sops *StrOps) GetValidString(
	targetStr string,
	validRunes []rune,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.GetValidString() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.getValidString(
		targetStr,
		validRunes,
		ePrefix)
}

// IsEmptyOrWhiteSpace - If a string is zero length or consists solely of
// white space (contiguous spaces), this method will return 'true'.
//
// Otherwise, a value of false is returned.
func (sops *StrOps) IsEmptyOrWhiteSpace(targetStr string) bool {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.isEmptyOrWhiteSpace(targetStr)
}

// justifyTextInStrField - Creates a and returns a new string text
// field with text 'strToJustify' positioned inside that new string
// in accordance with the string justification formatting passed in
// input parameter, 'textJustify'.
//
// 'textJustify' will specify either 'Right-Justify', 'Left-Justify'
// or 'Center'. Based on this specification, the newly created and
// returned text string will contain 'strToJustify' properly
// formatted as 'Right-Justified', 'Left-Justified' or 'Centered'.
//
// If the length of 'strToJustify' is greater than the output field
// length value, 'fieldLen', this method will increase the value of
// 'fieldLen' to match the length of 'strToJustify'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//    - The string content or text which will be positioned and
//      formatted inside the returned output string according to
//      the text justification specification contained in input
//      parameter, 'textJustify'.
//
//      If 'strToJustify' is a zero length empty string, this
//      method will return a string of white-space characters
//      equal in length to the value of input parameter,
//      'fieldLen'.
//
//
//  fieldLen            int
//     - The total length of the returned output string in which
//       'strToJustify' will be positioned and formatted. If this
//       'fieldLen' value is less than the length of 'strToJustify',
//       'fieldLen' will be automatically increased to a value equal
//       to the length of 'strToJustify'.
//
//
//  textJustify         TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to 'strToJustify' when
//       it is positioned inside of the returned output string.
//       This enumeration value must be one of the three following
//       format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The output string resulting from the text justification
//       operation described above. Input parameter, 'strToJustify'
//       will be formatted in this output string according to the
//       format specification defined by input parameter,
//       'textJustify'. The length of this output string is
//       controlled by input parameter, 'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrJustify_01() "
//  strToJustify := "12345"
//  fieldLen := 10
//
//  su := StrOps{}
//  strJustified, err :=
//   su.justifyTextInStrField(
//               strToJustify,
//               fieldLen,
//               TextJustify(0).Right()
//               ePrefix)
//
//  --------------------------------------------------------
//                                  1234567890
//  'strJustified' is now equal to "     12345"
//  The string length of 'strJustified' is 10
//
func (sops *StrOps) JustifyTextInStrField(
	strToJustify string,
	fieldLen int,
	textJustify TextJustify,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.JustifyTextInStrField() "

	sOpsNanobot := strOpsNanobot{}

	return sOpsNanobot.justifyTextInStrField(
		strToJustify,
		fieldLen,
		textJustify,
		ePrefix)
}

// LowerCaseFirstLetter - Finds the first alphabetic character
// in a string (a-z A-Z) and converts it to lower case.
//
func (sops *StrOps) LowerCaseFirstLetter(str string) string {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.lowerCaseFirstLetter(str)
}

// MakeSingleCharString - Creates a string of length 'strLen' consisting of
// a single character passed through input parameter, 'charRune' as type
// 'rune'.
//
// Example Usage:
//
//     sUtil := StrOps{}
//     requestedLen := 5
//     charRune := '='
//     outputStr, err := sUtil.MakeSingleCharString(charRune, requestedLen)
//
//     outputStr is now equal to "====="
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  charRune            rune
//     - The text character which will be repeated throughout the
//       length of the returned string.
//
//
//  strLen              int
//     - The length of the returned string containing the repeated
//       characters.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This returned string will have a length of 'strLen' and
//       contain 'strLen' characters all of which will be equal to
//       'charRune'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (sops *StrOps) MakeSingleCharString(
	charRune rune,
	strLen int,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.MakeSingleCharString() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.makeSingleCharString(
		charRune,
		strLen,
		ePrefix)
}

// NewPtr - Returns a pointer to a new instance of
// StrOps. Useful for cases requiring io.Reader
// and io.Writer.
func (sops StrOps) NewPtr() *StrOps {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sopsNew := StrOps{}

	sopsNew.stringDataMutex = new(sync.Mutex)

	return &sopsNew
}

// Ptr - Returns a pointer to a new instance of
// StrOps. Useful for cases requiring io.Reader
// and io.Writer.
//
// This method is identical to method StrOps.NewPtr().
//
// Example Usage:
//
// StrOps{}.Ptr().GetReader()
//
//
func (sops StrOps) Ptr() *StrOps {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sopsNew := StrOps{}

	sopsNew.stringDataMutex = new(sync.Mutex)

	return &sopsNew
}

// Read - Implements io.Reader interface. Read reads up to len(p)
// bytes into 'p'. This method supports buffered 'read' operations.
//
// The internal member string variable, 'StrOps.stringData' is written
// into 'p'. When the end of 'StrOps.stringData' is written to 'p',
// the method returns error = 'io.EOF'.
//
// 'StrOps.stringData' can be accessed through Getter an Setter methods,
// GetStringData() and SetStringData()
//
func (sops *StrOps) Read(p []byte) (n int, err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix := "StrOps.Read() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.readBytes(
		sops,
		p,
		ePrefix)

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
//  bytes               []byte
//     - An array of bytes from which a string will be extracted
//       and returned.
//
//  startIdx            int
//     - The starting index in input parameter 'bytes' where the string
//       extraction will begin. The string extraction will cease when
//       a carriage return ('\r'), a vertical tab ('\v') or a new line
//       character ('\n') is encountered.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  extractedStr        string
//     - The string extracted from input parameter 'bytes' beginning
//       at the index in 'bytes' indicated by input parameter 'startIdx'.
//
//  nextStartIdx        int
//     - The index of the beginning of the next string in the byte array
//       'bytes' after 'extractedString'. If no more strings exist in the
//       the byte array, 'nextStartIdx' will be set to -1.
//
func (sops *StrOps) ReadStringFromBytes(
	bytes []byte,
	startIdx int) (
	extractedStr string,
	nextStartIdx int) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.readStringFromBytes(
		bytes,
		startIdx)
}

// RemoveStringChar - Removes or deletes a specified character
// from a string and returns a new string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetStr                  string
//     - The target string containing the character to be removed.
//       If this is a zero length or empty string, an error will
//       be returned.
//
//
//  charToRemove               rune
//     - The character which will be removed from 'targetStr'. If
//       this is an empty character or zero value rune, an error
//       will be returned.
//
//
//  maxNumOfCharDeletions      int
//     - If this parameter is set to minus one (-1), all instances
//       of 'charToRemove' in 'targetStr' will be deleted in the
//       returned string. If this parameter is greater than zero,
//       it will limit the maximum number of character deletions
//       in this operation. Remember that the search for
//       'charToRemove' proceeds from left to right starting at
//       index zero (0).
//
//       If this parameter is set to zero, an error will be
//       returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  newStr                     string
//     - If this method completes successfully, a new string
//       will be returned containing all the characters in
//       input parameter 'targetStr' except those removed by
//       the deletion operation.
//
//
//  numOfDeletions             int
//     - This parameter will record the number of character
//       deletions performed by this operation.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (sops *StrOps) RemoveStringChar(
	targetStr string,
	charToRemove rune,
	maxNumOfCharDeletions int,
	ePrefix string) (
	newStr string,
	numOfDeletions int,
	err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.RemoveStringChar() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.removeStringChar(
		targetStr,
		charToRemove,
		maxNumOfCharDeletions,
		ePrefix)
}

// ReplaceBytes - Replaces characters in a target array of bytes ([]bytes) with those specified in
// a two dimensional slice of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetBytes         []byte
//     - The byte array which will be examined. If characters ('bytes') eligible
//       for replacement are identified by replacementBytes[i][0] they will be
//       replaced by the character specified in replacementBytes[i][1].
//
//
//  replacementBytes    [][]byte
//     - A two dimensional slice of type byte. Element [i][0] contains the
//       target character to locate in 'targetBytes'. Element[i][1] contains
//       the replacement character which will replace the target character
//       in 'targetBytes'. If the replacement character element [i][1] is
//       a zero value, the target character will not be replaced. Instead,
//       it will be eliminated or removed from the returned byte array ([]byte).
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  []byte
//     - The returned byte array containing the characters and replaced characters
//       from the original 'targetBytes' array.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "1a2b3c4d5e6"
//  testBytes := []byte(testStr)
//
//  replaceBytes := make([][]byte, 5, 10)
//
//  for i := 0; i < 5; i++ {
//   replaceBytes[i] = make([]byte, 2, 5)
//  }
//
//  replaceBytes[0][0] = 'a'
//  replaceBytes[0][1] = 'A'
//
//  replaceBytes[1][0] = 'b'
//  replaceBytes[1][1] = 'B'
//
//  replaceBytes[2][0] = 'c'
//  replaceBytes[2][1] = 'C'
//
//  replaceBytes[3][0] = 'd'
//  replaceBytes[3][1] = 'D'
//
//  replaceBytes[4][0] = 'e'
//  replaceBytes[4][1] = 'E'
//
//  ePrefix := "TestStrOps_ReplaceBytes_01() "

//  actualRunes, err := StrOps{}.Ptr().ReplaceBytes(
//  testBytes,
//  replaceBytes,
//  ePrefix)
//
//  actualRunes = "1A2B3C4D5E6"
//
func (sops *StrOps) ReplaceBytes(
	targetBytes []byte,
	replacementBytes [][]byte,
	ePrefix string) (
	[]byte,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ReplaceBytes() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.replaceBytes(
		targetBytes,
		replacementBytes,
		ePrefix)
}

// ReplaceMultipleStrs - Replaces all instances of string replaceArray[i][0] with
// replacement string from replaceArray[i][1] in 'targetStr'.
//
// Input parameter 'replaceArray' should be passed as a two-dimensional slice.
// If the length of the 'replaceArray' second dimension is less than '2', an
// error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of sub-string replaceArray[i][0]. Instances of
//       this string found in 'targetStr' will be replaced with the
//       string from replaceArray[i][1].
//
//  replaceArray        [][]string
//     - A two dimensional array of type 'string'. 'targetStr' will
//       be searched for instances of replaceArray[i][0]. If found
//       in 'targetStr' this string will be replaced with the string
//       from replaceArray[i][1].
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This is the 'result' string. It is identical to 'targetStr'
//       except for the string substitutions made from replaceArray[i][1].
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_ReplaceMultipleStrs_08() "
//  rStrs := make([][]string, 3, 5)
//
//  for i := 0; i < 3; i++ {
//   rStrs[i] = make([]string, 2, 5)
//  }
//
//  testStr := "HeFFxJWxrFd"
//
//  rStrs[0][0] = "x"
//  rStrs[0][1] = "o"
//  rStrs[1][0] = "J"
//  rStrs[1][1] = " "
//  rStrs[2][0] = "F"
//  rStrs[2][1] = "l"
//
//
//  actualStr, err := StrOps{}.Ptr().ReplaceMultipleStrs(
//                      testStr,
//                      rStrs,
//                      ePrefix)
//
//  'actualStr' is now equal to "Hello World"
//
//
func (sops *StrOps) ReplaceMultipleStrs(
	targetStr string,
	replaceArray [][]string,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ReplaceMultipleStrs() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.replaceMultipleStrs(
		targetStr,
		replaceArray,
		ePrefix)
}

// ReplaceNewLines - Replaces New Line characters from string. If the specified
// replacement string is empty, the New Line characters are simply removed
// from the input parameter, 'targetStr'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetStr                  string
//     - The target string containing the new line characters to be
//       removed. If this is a zero length or empty string, no action
//       will be taken.
//
//  replacementStr             string
//     - The string which will replace the new line character. If
//       this parameter is an empty string, the new line characters
//       will simply be deleted from the returned string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - The returned string which contains substitutions of
//       'replacementStr' for the new line character in
//       'targetStr'.
//
func (sops *StrOps) ReplaceNewLines(
	targetStr string,
	replacementStr string) string {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	if len(targetStr) == 0 {
		return targetStr
	}

	sOpsQuark := strOpsQuark{}
	var newStr string

	if len(replacementStr) == 0 {
		newStr,
			_,
			_ =
			sOpsQuark.removeSubString(
				targetStr,
				"\n",
				-1,
				"")

		return newStr
	}

	newStr,
		_,
		_ = sOpsQuark.replaceSubString(
		targetStr,
		"\n",
		replacementStr,
		-1,
		"")

	return newStr
}

// ReplaceRunes - Replaces individual characters in a target array
// of runes ([]rune) with those specified in a two-dimensional
// slice of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetRunes         []rune
//     - The rune array which will be examined. If target characters
//       ('runes') eligible for replacement are identified by
//       replacementRunes[i][0], they will be replaced by the
//       character specified in replacementRunes[i][1].
//
//
//  replacementRunes    [][]rune
//     - A two dimensional slice of type 'rune'. Element [i][0]
//       contains the target character to locate in 'targetRunes'.
//       Element[i][1] contains the replacement character which will
//       replace the target character in 'targetRunes'. If the
//       replacement character element [i][1] is a zero value, the
//       target character will not be replaced. Instead, it will be
//       eliminated or removed from the returned rune array
//       ([]rune).
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  []rune
//     - The returned rune array containing the characters and
//       replaced characters from the original 'targetRunes' array.
//
//
//  error
//     - If the method completes successfully this value is 'nil'.
//       If an error is encountered this value will contain the
//       error message. Examples of possible errors include a zero
//       length 'targetRunes' array or 'replacementRunes' array.
//
//       In addition, if any of the replacementRunes[][x] 2nd
//       dimension elements have a length less than two, an
//       error will be returned.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_ReplaceRunes_02() "
//
//  testStr := "1a2b3c4d5e6"
//  testRunes := []rune(testStr)
//
//  replaceRunes := make([][]rune, 5, 10)
//
//  for i := 0; i < 5; i++ {
//    replaceRunes[i] = make([]rune, 2, 5)
//  }
//
//  replaceRunes[0][0] = 'a'
//  replaceRunes[0][1] = 0
//
//  replaceRunes[1][0] = 'b'
//  replaceRunes[1][1] = 0
//
//  replaceRunes[2][0] = 'c'
//  replaceRunes[2][1] = 0
//
//  replaceRunes[3][0] = 'd'
//  replaceRunes[3][1] = 0
//
//  replaceRunes[4][0] = 'e'
//  replaceRunes[4][1] = 0
//
//  actualRunes, err := StrOps{}.Ptr().ReplaceRunes(
//  testRunes,
//  replaceRunes,
//  ePrefix)
//
//  actualStr := string(actualRunes)
//  --------------------------------------
//  Original testStr := "1a2b3c4d5e6"
//  actualStr is now equal to "123456"
//
func (sops *StrOps) ReplaceRunes(
	targetRunes []rune,
	replacementRunes [][]rune,
	ePrefix string) (
	[]rune,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ReplaceRunes() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.replaceRunes(
		targetRunes,
		replacementRunes,
		ePrefix)
}

// ReplaceStringChar - Replaces a specific character found anywhere
// in a string with another specified substitute character.
//
// The replacement operation proceeds from left to right within the
// 'targetStr' beginning with index zero (0).
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr                  string
//     - The string containing the character to be replaced.
//       If this is an empty string, an error will be returned.
//
//
//  charToReplace              rune
//     - The character within input parameter string 'targetStr'
//       which will be replaced. If this parameter is set to zero
//       signaling an empty character, this method will return an
//       error.
//
//
//  replacementChar            rune
//     - The character which will replace 'charToReplace' in
//       'targetStr'. If this parameter is set to zero signaling an
//       empty character, this method will return an error.
//
//
//  maxNumOfReplacements       int
//     - The maximum number of replacements allowed for this
//       operation. If this parameter is set to minus one (-1), all
//       instances of 'charToReplace' in 'targetStr' will be
//       replaced with 'replacementChar'.
//
//       If this parameter is set to zero ('0'), an error will be
//       returned.
//
//       The replacement operation proceeds from left to right
//       within the 'targetStr' beginning with index zero (0).
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, a new string will be
//       returned with the designated replacement characters.
//
//
//  int
//     - The integer value records the number of characters
//       replaced in 'targetStr'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (sops *StrOps) ReplaceStringChar(
	targetStr string,
	charToReplace rune,
	replacementChar rune,
	maxNumOfReplacements int,
	ePrefix string) (
	string,
	int,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ReplaceStringChar() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.replaceStringChar(
		targetStr,
		charToReplace,
		replacementChar,
		maxNumOfReplacements,
		ePrefix)
}

// ReplaceStringChars - Replaces string characters in a target
// string ('targetStr') with those specified in a two dimensional
// slice of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be examined. If target string characters
//       eligible for replacement are identified by replacementRunes[i][0],
//       they will be replaced by the character specified in
//       replacementRunes[i][1].
//
//  replacementRunes    [][]rune
//     - A two dimensional slice of type 'rune'. Element [i][0] contains
//       the target character to locate in 'targetStr'. Element[i][1]
//       contains the replacement character which will replace the target
//       character in 'targetStr'. If the replacement character
//       element [i][1] is a zero value, the target character will not
//       be replaced. Instead, it will be eliminated or removed from the
//       returned string.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string  - The returned string containing the characters and replaced characters
//            from the original target string, ('targetStr').
//
//  error   - If the method completes successfully this value is 'nil'. If an error is
//            encountered this value will contain the error message. Examples of possible
//            errors include a zero length 'targetStr' or 'replacementRunes[][]' array.
//            In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//            a length less than two, an error will be returned.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_ReplaceStringChars_06() "
//
//  testStr := "1a2b3c4d5e6"
//
//  replaceRunes := make([][]rune, 5, 10)
//
//  for i := 0; i < 5; i++ {
//  replaceRunes[i] = make([]rune, 2, 5)
//  }
//
//  replaceRunes[0][0] = 'a'
//  replaceRunes[0][1] = 0
//
//  replaceRunes[1][0] = 'b'
//  replaceRunes[1][1] = 0
//
//  replaceRunes[2][0] = 'c'
//  replaceRunes[2][1] = 0
//
//  replaceRunes[3][0] = 'd'
//  replaceRunes[3][1] = 0
//
//  replaceRunes[4][0] = 'e'
//  replaceRunes[4][1] = 0
//
//  actualStr, err := StrOps{}.Ptr().ReplaceStringChars(
//  testStr,
//  replaceRunes,
//  ePrefix)
//
//  ---------------------------------------------------
//  Original testStr := "1a2b3c4d5e6"
//
//  actualStr is now equal to = "123456"
//
func (sops *StrOps) ReplaceStringChars(
	targetStr string,
	replacementRunes [][]rune,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.ReplaceStringChars() "

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.replaceStringChars(
		targetStr,
		replacementRunes,
		ePrefix)
}

// SetStringData - Sets the value of internal
// string data element, StrOps.stringData. It
// also zeros internal fields sops.cntBytesWritten
// and sops.cntBytesRead.
func (sops *StrOps) SetStringData(str string) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sops.stringData = str

	sops.cntBytesWritten = 0

	sops.cntBytesRead = 0
}

// StrCenterInStrLeft - returns a string which includes a left pad blank string
// plus the original string. It does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively center the original string
// in a field of specified length.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//     - This string will be centered in a text field. The text
//       field length is defined by input parameter, 'fieldLen'.
//
//
//  fieldLen            int
//     - Defines the length of a text field in which 'strToCenter'
//       will be centered.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This returned string contains 'strToCenter' with the
//       necessary left-pad number of spaces required for
//       centering.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//   In this example the total field length is 15. The length of the
//   test string, "Hello", is 5. In order to center the test string
//   in a field with length of 15, there will be 5-spaces on the
//   left and 5-spaces on the right. This method will compute the
//   left-pad spaces necessary to center the string in a field with
//   length of 15, but will only include the padded left margin of
//   5-spaces. It will NOT include the trailing 5-spaces on the
//   right.
//
//   In the following example, the final result string will
//   substitute the'@' character for the white space character
//   (0x20) in order to illustrate the padding added by this method.
//
//    strToCenter     = "Hello"
//    fieldLen        = 15
//    Returned String = "@@@@@Hello" or "     Hello"
//
//
func (sops *StrOps) StrCenterInStrLeft(
	strToCenter string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.StrCenterInStrLeft() "

	sOpsNanobot := strOpsNanobot{}

	return sOpsNanobot.strCenterInStrLeft(
		strToCenter,
		fieldLen,
		ePrefix)
}

// StrCenterInStr - returns a string which includes a left pad blank string plus
// the original string ('strToCenter'), plus a right pad blank string.
//
// The returned string will effectively center the original string ('strToCenter')
// in a field of specified length ('fieldLen').
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//     - This string will be centered in a text field. The text
//       field length is defined by input parameter, 'fieldLen'.
//
//
//  fieldLen            int
//     - Defines the length of a text field in which 'strToCenter'
//       will be centered.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This returned string contains 'strToCenter' with the
//       necessary left-pad and right-pad number of spaces
//       required for centering. The total length of this string
//       will be equal to input parameter, 'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_StrCenterInStr_02() "
//  strToCenter := "Hello"
//  fieldLen := 15
//
//  su := StrOps{}
//  centeredStr, err := su.StrCenterInStr(
//  strToCenter,
//  fieldLen,
//  ePrefix)
//
//  ---------------------------------------------
//                               123456789012345
//  centeredStr is now equal to "     Hello     "
//  'Hello' is centered in a field of length 15
//  with left and right pad of 5-spaces.
//
func (sops *StrOps) StrCenterInStr(
	strToCenter string,
	fieldLen int,
	ePrefix string) (string, error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.StrCenterInStr() "

	sOpsMolecule := strOpsMolecule{}

	return sOpsMolecule.strCenterInStr(
		strToCenter,
		fieldLen,
		ePrefix)
}

// StrGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (sops *StrOps) StrGetRuneCnt(
	targetStr string) int {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.getRuneCountInStr(targetStr)
}

// StrGetCharCnt - Uses the 'len' method to
// return the number of rune characters in a
// string.
func (sops *StrOps) StrGetCharCnt(
	targetStr string) int {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.getCharCountInStr(targetStr)
}

// StripBadChars - Removes/deletes specified sub-strings from a
// parent or host string. The targeted sub-strings are deleted
// wherever found in the parent or host string.
//
// The sub-strings to be removed are identified in a string array
// passed as input parameter, 'badChars'.
//
// All instances of 'badChars' sub-strings are deleted from the
// target string which is passed as input parameter, 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be searched for the sub-strings
//       identified in the 'badChars' array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from input parameter,
//       'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       sub-strings identified in the 'badChars' array which are
//       deleted.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//   badChars := []string{"@@"}
//
//   expectedStrLen := len(expectedStr)
//                  12    123456789    12  12
//   testString := "@@Some@@@@@@@@@Stri@@ng@@"
//
//   actualString, actualStrLen :=
//         StrOps{}.Ptr().StripBadChars(
//                            testString,
//                            badChars)
//
//  -----------------------------------------------
//                                 12345678901
//   actualString is now equal to "Some@String"
//   actualStrLen is now equal to 11
//
func (sops *StrOps) StripBadChars(
	targetStr string,
	badChars []string) (
	cleanStr string,
	strLen int) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.stripBadChars(
		targetStr,
		badChars)
}

// StripLeadingChars - Strips or deletes sub-strings from the front
// of a parent or host string labeled here as 'targetStr'. The
// sub-strings to be deleted are specified by input parameter and
// string array, 'badChars'.
//
// The method then returns a string which does not contain leading
// sub-strings identified in the 'badChars' array. In addition, the
// length of the final 'clean string' is also returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be searched for leading sub-strings
//       identified in the 'badChars' array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from input parameter,
//       'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       sub-strings identified in the 'badChars' array which are
//       deleted from the front of 'targetStr'.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  badChars := []string{
//                 " ", // A single white-space character
//                 "/",
//                 "//",
//                 "\\\\",
//                 "\\",
//                 ".\\",
//                 "../",
//                 ".",
//                 "..\\",
//                 "\\\\\\",
//                 "..",
//                 "./",
//                 "//",
//                 "///",
//                 "////",
//                 "..."}
//
//  testString :=
//   "..........      ./../.\\.\\..\\////   SomeString"
//
//  actualString, actualStrLen :=
//      StrOps{}.Ptr().StripLeadingChars(
//                       testString,
//                       badChars)
//
//  ----------------------------------------------------
//                                1234567890
//  actualString is now equal to "SomeString"
//  actualStrLen is now equal to 10
//
func (sops *StrOps) StripLeadingChars(
	targetStr string,
	badChars []string) (
	cleanStr string,
	strLen int) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.stripLeadingChars(
		targetStr,
		badChars)
}

// StripTrailingChars - Strips or deletes sub-strings from the
// end of a parent or host string. The sub-strings to be deleted
// are identified in a string array input parameter labeled,
// 'badChars'.  The parent string to be searched is passed as input
// parameter, 'targetStr'. The targeted sub-strings are only deleted
// if they exist at the end of 'targetStr'.
//
// Upon completion, this method returns the cleaned string and the
// length of the cleaned string to the caller.  The cleaned string
// is equivalent to input parameter, 'targetStr', minus the trailing
// sub-strings identified by string array 'badChars'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of trailing sub-strings identified in the
//       'badChars' string array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from the end of 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       trailing sub-strings identified for deletion in the
//       'badChars' array.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  badChars := []string{
//                 " ", // Single white space character
//                 "/",
//                 "//",
//                 "\\\\",
//                 "\\",
//                 ".\\",
//                 "../",
//                 ".",
//                 "..\\",
//                 "\\\\\\",
//                 "..",
//                 "./",
//                 "//",
//                 "///",
//                 "////",
//                 "..."}
//
//  testString :=
//   "SomeString..........      ./../.\\.\\..\\////   "
//
//  actualString, actualStrLen :=
//    StrOps{}.Ptr().StripTrailingChars(
//                      testString,
//                      badChars)
//
//  -------------------------------------------------------------
//
//                                1234567890
//  actualString is now equal to "SomeString"
//  actualStrLen is now equal to 10
//
func (sops *StrOps) StripTrailingChars(
	targetStr string,
	badChars []string) (cleanStr string, strLen int) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.stripTrailingChars(
		targetStr,
		badChars)
}

// StrLeftJustify - Creates a new string, left-justified, within a
// with a wider text field or output string. The text to be left
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the
// output string and spaces are padded to the right in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//     - The text content which will be left-justified in the
//       output string returned by this method.
//
//  fieldLen            int
//     - Defines the length of the output string in which input
//       parameter 'strToJustify' will be left-justified.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The output string resulting from the 'left-justify'
//       operation. Input parameter, 'strToJustify' will be
//       left-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrLeftJustify_01() "
//  fieldLen = 15
//  strToJustify    = "Hello World"
//  su := StrOps{}
//  justifiedStr, err := su.StrLeftJustify(
//                           strToJustify,
//                           fieldLen,
//                           ePrefix)
//
//  ------------------------------------------------
//                                  123456789012345
//  'justifiedStr' is now equal to "Hello World    "
//  The string length of 'justifiedStr' is 15
//
//
func (sops *StrOps) StrLeftJustify(
	strToJustify string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.StrLeftJustify() "

	sOpsMolecule := strOpsMolecule{}

	return sOpsMolecule.strLeftJustify(
		strToJustify,
		fieldLen,
		ePrefix)
}

// StrPadLeftToCenter - Returns a blank string which allows
// centering of the target string in a fixed length field. A blank
// string is defined here as a string consisting entirely of white
// space characters (' ' 0x20). This string can be prefixed or
// added to the beginning of a text string to achieve a centering
// effect when displayed.
//
// Therefore the blank string returned by this method only
// constitutes the left-padding of white space necessary to center
// the input parameter, 'strToCenter'. It does not include the
// right padding of white space.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//    - The content or text string which will be centered.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be centered.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The output string resulting from the string centering
//       operation. This string will consist entirely of white
//       space (' ' 0x20 characters). It will represent the left
//       padding necessary to center the text string,
//       'strToCenter'. See 'Example Usage' section below.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// Assume that total field length ('fieldlen') is 70. Further
// assume that the string to Center ('strToCenter') is
// 10-characters in length. In order to center a 10-character
// string in a 70-character field, 30-space characters would need
// to be positioned on each side of the string to center. This
// method only returns the left margin, or in this example, a
// string consisting of 30-spaces.
//
//
//  ePrefix := "TestStrOps_StrPadLeftToCenter_02() "
//  //              12345
//  strToCenter := "Hello"
//  fieldLen := 15
//  su := StrOps{}
//  padStr, err := su.StrPadLeftToCenter(
//                   strToCenter,
//                   fieldLen,
//                   ePrefix)
//
//  -------------------------------------------------------------
//                          12345
//  'padStr' is now equal to "     "
//  'padStr' consists of 5-spaces.
//  padStr + strToCenter will yield a centered string.
//
//
func (sops *StrOps) StrPadLeftToCenter(
	strToCenter string,
	fieldLen int,
	ePrefix string) (string, error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.StrPadLeftToCenter() "

	sOpsMolecule := strOpsMolecule{}

	return sOpsMolecule.strPadLeftToCenter(
		strToCenter,
		fieldLen,
		ePrefix)
}

// StrRightJustify - Creates a new string, right-justified, within
// a wider text field or output string. The text to be right
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the right side of
// the output string and spaces are padded to the left in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//    - The content or text string which will be right justified.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be right-justified.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The output string resulting from the 'right-justify'
//       operation. Input parameter, 'strToJustify' will be
//       right-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  If the total field length ('fieldLen') is specified as
//  10-characters and the length of string to justify
//  ('strToJustify') is 5-characters, then this method would return
//  a string consisting of 5-space characters plus the
//  'strToJustify'.
//
//  ePrefix := "TestStrOps_StrRightJustify_05() "
//  strToJustify := "12345"
//  fieldLen := 10
//
//  su := StrOps{}
//  strRightJustified, err :=
//   su.StrRightJustify(
//               strToJustify,
//               fieldLen,
//               ePrefix)
//
//  --------------------------------------------------------
//                                       1234567890
//  'strRightJustified' is now equal to "     12345"
//  The string length of 'strRightJustified' is 10
//
//
func (sops *StrOps) StrRightJustify(
	strToJustify string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.StrRightJustify() "

	sOpsMolecule := strOpsMolecule{}

	return sOpsMolecule.strRightJustify(
		strToJustify,
		fieldLen,
		ePrefix)
}

// SwapRune - Swaps, or replaces, instances of the 'oldRune'
// character with the 'newRune' character in the input parameter
// target string ('targetStr'). The number of character
// substitutions performed is controlled by input parameter
// 'maxNumOfSwaps'.
//
// If input parameter 'targetStr' is an empty string, no error will
// be generated and the empty 'targetStr' will be returned.
//
// For more information on 'runes', reference:
//   https://www.geeksforgeeks.org/rune-in-golang/
//   https://golangbyexample.com/understanding-rune-in-golang/
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - This string will be searched for instances of 'oldRune'.
//       When 'oldRune' is located in this string, it will be
//       replaced with 'newRune'.
//
//
//  oldRune             rune
//     - This rune or text character will be used search the parent
//       string, 'targetStr'. When instances of 'oldRune' are
//       located in 'targetStr', they will be replaced with
//       'newRune'.
//
//
//  newRune             rune
//     - This rune or text character will be used to replaced
//       instances of 'oldRune' located in a search of 'targetStr'.
//
//
//  maxNumOfSwaps       int
//     - This integer value constitutes the upper limit for the
//       number of character substitutions performed by this method.
//       If this value is less than one (+1), no limit will be
//       imposed and all instances of 'oldRune' found in 'targetStr'
//       will be replaced with 'newRune'.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This string will be populated with the contents of
//       'targetStr' after instances of 'oldRune' have been
//        replaced by 'newRune'. Effectively, this string is the
//        result of the character substitution operation.
//
//  int
//     - The value of this returned integer records the number of
//       character substitutions performed on 'targetStr'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     tStr := "  Hello   World  "
//     ePrefix := "theCallingMethodName "
//
//     su := StrOps{}
//        resultStr,
//        numOfReplacements,
//        err := su.SwapRune(
//                 tStr,
//                 ' ',
//                 '!',
//                 -1,
//                 ePrefix)
//
//
//    resultSt is now equal to "!!Hello!!!World!!"
//    numOfReplacements is equal to '7'
//
func (sops *StrOps) SwapRune(
	targetStr string,
	oldRune rune,
	newRune rune,
	maxNumOfSwaps int,
	ePrefix string) (
	string,
	int,
	error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.SwapRune() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.swapRune(
		targetStr,
		oldRune,
		newRune,
		maxNumOfSwaps,
		ePrefix)
}

// TrimMultipleChars - Performs the following operations on strings:
//
//  1. Trims Right and Left ends of 'targetStr' for all instances
//     of 'trimChar'
//
//  2. Within the interior of a string, multiple instances of
//     'trimChar' are reduced to a single instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of the character 'trimChar'.
//
//
//  trimChar            rune
//     - 'targetStr' will be searched for instances of this
//       character. If this character is a leading or trailing
//       character in 'targetStr', those instances will be deleted.
//       If this character is found in the interior of 'targetStr',
//       it will be reduced to one single instance of the
//       character.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  rStr                string
//     - This is the 'result' string. It is comprised of all the
//       characters in the original 'targetStr' minus those
//       'trimChar' characters deleted in the 'trim' operation.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "theCallingFunctionName() "
//  targetStr = "       Hello          World        "
//  trimChar  = ' ' (One Space Character)
//  sops := StrOps{}
//
//  rStr,
//  err := sops.TrimMultipleChars(
//           targetStr,
//           trimChar,
//           ePrefix)
//
//  returned string (rStr) now equal to "Hello World"
//
//
func (sops StrOps) TrimMultipleChars(
	targetStr string,
	trimChar rune,
	ePrefix string) (
	rStr string,
	err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.TrimMultipleChars() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.trimMultipleChars(
		targetStr,
		trimChar,
		ePrefix)
}

// TrimStringEnds - Removes all instances of input parameter
// 'trimChar' from the beginning and end of input parameter string
// 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of the character 'trimChar'.
//
//
//  trimChar            rune
//     - 'targetStr' will be searched for instances of this
//       character. If the character is found to be either a
//       trailing character or a leading character, it will be
//       deleted. If this character exists in the interior of
//       'targetStr' it will be ignored and NOT deleted.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  rStr                string
//     - This is the 'result' string. It is comprised of all the
//       characters in the original 'targetStr' minus those
//       'trimChar' characters deleted in the 'trim' operation.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     ePrefix := "TestStrOps_TrimStringEnds_01() "
//     tStr := "XXXHello WorlXdXXX"
//     trimChar := 'X'
//
//     sops := StrOps{}
//
//     result,
//     err := sops.TrimStringEnds(
//              tStr,
//              trimChar,
//              ePrefix)
//
//  result is now equal to "Hello WorlXd"
//
//
func (sops StrOps) TrimStringEnds(
	targetStr string,
	trimChar rune,
	ePrefix string) (
	rStr string,
	err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	ePrefix += "StrOps.TrimStringEnds() "

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.trimStringEnds(
		targetStr,
		trimChar,
		ePrefix)
}

// UpperCaseFirstLetter - Finds the first alphabetic character in a string
// (a-z A-Z) and converts it to upper case.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  str                 string
//     - The first character in this string will be converted to
//       an Upper Case Letter.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string              string
//     - This is copy of input parameter 'str' with the first
//       character converted to upper case (a.k.a. a capital
//       letter).
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//   str := "how now brown cow."
//
//   sops := StrOps{}
//
//   actualStr := sops.UpperCaseFirstLetter(str)
//
//  'actualStr' is now equal to "How now brown cow."
//
//
func (sops *StrOps) UpperCaseFirstLetter(
	str string) string {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsQuark := strOpsQuark{}

	return sOpsQuark.upperCaseFirstLetter(str)
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
//
func (sops *StrOps) Write(p []byte) (n int, err error) {

	if sops.stringDataMutex == nil {
		sops.stringDataMutex = new(sync.Mutex)
	}

	sops.stringDataMutex.Lock()

	defer sops.stringDataMutex.Unlock()

	sOpsElectron := strOpsElectron{}

	return sOpsElectron.write(
		sops,
		p,
		"StrOps.Write() ")
}
