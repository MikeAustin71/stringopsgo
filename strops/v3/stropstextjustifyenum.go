package strops

import (
	"fmt"
	"strings"
	"sync"
)

var mStrOpsTextJustifyCodeToString = map[StrOpsTextJustify]string{
	StrOpsTextJustify(0): "None",
	StrOpsTextJustify(1): "Left",
	StrOpsTextJustify(2): "Right",
	StrOpsTextJustify(3): "Center",
}

var mStrOpsTextJustifyStringToCode = map[string]StrOpsTextJustify{
	"None":     StrOpsTextJustify(0),
	"Left":     StrOpsTextJustify(1),
	"Right":    StrOpsTextJustify(2),
	"Center":   StrOpsTextJustify(3),
	"Centered": StrOpsTextJustify(3),
}

var mStrOpsTextJustifyLwrCaseStringToCode = map[string]StrOpsTextJustify{
	"none":     StrOpsTextJustify(0),
	"left":     StrOpsTextJustify(1),
	"right":    StrOpsTextJustify(2),
	"center":   StrOpsTextJustify(3),
	"centered": StrOpsTextJustify(3),
}

// StrOpsTextJustify - An enumeration of text justification designations.
// StrOpsTextJustify is used to specify 'Right-Justified',
// 'Left-Justified' and 'Centered' string positioning within text
// fields.
//
// Since Go does not directly support enumerations, the 'StrOpsTextJustify'
// type has been adapted to function in a manner similar to classic
// enumerations. 'StrOpsTextJustify' is declared as a type 'int'. The
// method names effectively represent an enumeration of text
// justification formats. These methods are listed as follows:
//
// None            (0) - Signals that 'StrOpsTextJustify' value has NOT
//                       been initialized. This is an error condition.
//
//
// Left            (1) - Signals that the text justification format is
//                       set to 'Left-Justify'. Strings within text
//                       fields will be flush with the left margin.
//
//                           Example: "TextString      "
//
//
// Right           (2) - Signals that the text justification format is
//                       set to 'Right-Justify'. Strings within text
//                       fields will terminate at the right margin.
//
//                           Example: "      TextString"
//
//
// Center          (3) - Signals that the text justification format is
//                       is set to 'Centered'. Strings will be positioned
//                       in the center of the text field equidistant
//                       from the left and right margins.
//
//                           Example: "   TextString   "
//
//
// For easy access to these enumeration values, use the global variable
// 'TxtJustify'. Example: TxtJustify.Right()
//
// Otherwise you will need to use the formal syntax.
// Example: StrOpsTextJustify(0).Right()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the StrOpsTextJustify methods in alphabetical order. Be advised that all 'StrOpsTextJustify'
// methods beginning with 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration values.
//
type StrOpsTextJustify int

var lockStrOpsTextJustify sync.Mutex

// None - Signals that 'SOpsTextJustify' value has NOT been initialized.
// This is an error condition.
//
// The 'None' StrOpsTextJustify integer value is zero (0).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify StrOpsTextJustify) None() StrOpsTextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return StrOpsTextJustify(0)
}

// Left - Signals that the text justification format is set to
// 'Left-Justify'. Strings within text fields will be flush with
// the left margin.
//
//        Example: "TextString      "
//
// The 'Left' text justification has a StrOpsTextJustify
// integer value of one (+1).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify StrOpsTextJustify) Left() StrOpsTextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return StrOpsTextJustify(1)
}

// Right - Signals that the text justification format is
// set to 'Right-Justify'. Strings within text fields will
// terminate at the right margin.
//
//        Example: "      TextString"
//
// The 'Right' text justification has a StrOpsTextJustify
// integer value of two (+2).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify StrOpsTextJustify) Right() StrOpsTextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return StrOpsTextJustify(2)
}

// Center - Signals that the text justification format is
// is set to 'Center'. Strings will be positioned in the
// center of the text field equidistant from the left and
// right margins.
//
//        Example: "   TextString   "
//
// The 'Center' text justification has a StrOpsTextJustify
// integer value of three (+3).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify StrOpsTextJustify) Center() StrOpsTextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return StrOpsTextJustify(3)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'StrOpsTextJustify'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= StrOpsTextJustify(0).Center()
// str := t.String()
//     str is now equal to 'Center'
//
func (sopsTxtJustify StrOpsTextJustify) String() string {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	result, ok :=
		mStrOpsTextJustifyCodeToString[sopsTxtJustify]

	if !ok {
		return "Error: StrOpsTextJustify code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// StrOpsTextJustify value is valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  textJustification := StrOpsTextJustify(0).Right()
//
//  isValid := textJustification.XIsValid()
//
func (sopsTxtJustify StrOpsTextJustify) XIsValid() bool {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	if sopsTxtJustify > 3 ||
		sopsTxtJustify < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of StrOpsTextJustify is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//                        and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'gregorian' will NOT
//                        match the enumeration name, 'Gregorian'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'gregorian'
//                        will match match enumeration name 'Gregorian'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// StrOpsTextJustify
//     - Upon successful completion, this method will return a new
//       instance of StrOpsTextJustify set to the value of the enumeration
//       matched by the string search performed on input parameter,
//       'valueString'.
//
// error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := StrOpsTextJustify(0).XParseString("Right", true)
//
//     t is now equal to StrOpsTextJustify(0).Right()
//
func (sopsTxtJustify StrOpsTextJustify) XParseString(
	valueString string,
	caseSensitive bool) (StrOpsTextJustify, error) {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	ePrefix := "StrOpsTextJustify.XParseString() "

	if len(valueString) < 4 {
		return StrOpsTextJustify(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strOpsTxtJustification StrOpsTextJustify

	if caseSensitive {

		strOpsTxtJustification, ok = mStrOpsTextJustifyStringToCode[valueString]

		if !ok {
			return StrOpsTextJustify(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid StrOpsTextJustify Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strOpsTxtJustification, ok = mStrOpsTextJustifyLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return StrOpsTextJustify(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid StrOpsTextJustify Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strOpsTxtJustification, nil
}

// XValue - This method returns the enumeration value of the current
// StrOpsTextJustify instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (sopsTxtJustify StrOpsTextJustify) XValue() StrOpsTextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return sopsTxtJustify
}

// XValueInt - This method returns the integer value of the current
// StrOpsTextJustify instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (sopsTxtJustify StrOpsTextJustify) XValueInt() int {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return int(sopsTxtJustify)
}

// TxtJustify - public global variable of
// type StrOpsTextJustify.
//
// This variable serves as an easier, short hand
// technique for accessing StrOpsTextJustify values.
//
// Usage:
// TxtJustify.None(),
// TxtJustify.Left(),
// TxtJustify.Right(),
// TxtJustify.Center(),
//
var TxtJustify StrOpsTextJustify
