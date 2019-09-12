package strops

import "testing"

func TestStrOps_FindFirstNonSpaceChar_01(t *testing.T) {
	//012345678901234
	testStr := "   Hello World"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, 4)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if 3 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='3'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}

}

func TestStrOps_FindFirstNonSpaceChar_02(t *testing.T) {

	//012345678901234
	testStr := "       Hello"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, 6)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if -1 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='-1'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}

}

func TestStrOps_FindFirstNonSpaceChar_03(t *testing.T) {

	//012345678901234
	testStr := "Hello  There"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 7, 9)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if 7 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='7'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}

}

func TestStrOps_FindFirstNonSpaceChar_04(t *testing.T) {
	//012345678901234
	testStr := "xx       H"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 2, 9)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if 9 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='9'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}

}

func TestStrOps_FindFirstNonSpaceChar_05(t *testing.T) {
	//012345678901234
	testStr := "      Hello World"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, 9)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if 6 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='6'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}

}

func TestStrOps_FindFirstNonSpaceChar_06(t *testing.T) {
	//012345678901234
	testStr := "Hello World"

	_, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, -1, 9)

	if err == nil {
		t.Error("Expected an ERROR return. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindFirstNonSpaceChar_07(t *testing.T) {

	//          012345678901234
	testStr := "        Hello "

	_, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 10, 9)

	if err == nil {
		t.Error("Expected an ERROR return. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindFirstNonSpaceChar_08(t *testing.T) {

	//          012345678901234
	testStr := "012345 78901234"

	_, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, 15)

	if err == nil {
		t.Error("Expected an ERROR return. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindFirstNonSpaceChar_09(t *testing.T) {

	//          012345678901234
	testStr := "012345 78901234"

	_, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, -1)

	if err == nil {
		t.Error("Expected an ERROR return. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindFirstNonSpaceChar_10(t *testing.T) {
	//012345678901234
	testStr := "xx       Hxglt"

	firstNonSpaceIdx, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 2, 9)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4). "+
			"Error='%v' ", err.Error())
	}

	if 9 != firstNonSpaceIdx {
		t.Errorf("Error: Expected firstNonSpaceIdx='9'. Instead, Idx='%v' ",
			firstNonSpaceIdx)
	}
}

func TestStrOps_FindFirstNonSpaceChar_11(t *testing.T) {
	//012345678901234
	testStr := "  "

	index, err :=
		StrOps{}.FindFirstNonSpaceChar(testStr, 0, 4)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4)\n"+
			"Error='%v'\n", err.Error())
	}

	if index != -1 {
		t.Errorf("Expected 'index' returned by StrOps{}.FindFirstNonSpaceChar(testStr,0,4)\n"+
			"would be equal to -1 because parameter, 'testStr' consists entirely of spaces.\n"+
			"However, the returned index was %v.", index)
	}
}

func TestStrOps_FindLastNonSpaceChar_01(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "xx       Hxgltx     "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 19)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 14 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='14'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_02(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "Now is the time for all good men to come to the aid of their country."

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 68)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 68 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='68'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_03(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "Now is the time for all good men to come to the aid of their country.           "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 79)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 68 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='68'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_04(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "Now is the time for all good men to come to the aid of their country.           "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 68, 79)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 68 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='68'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_05(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "Now is the time for all good men to come to the aid of their country.           "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 59, 79)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 68 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='68'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_06(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "                                                                                 "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 80)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if -1 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='-1'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_07(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "        XX        XX          XX             XXX      XXX.                       "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 40, 80)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 57 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='57'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_08(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := " The cow jumped over the moon."

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 10, 29)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 29 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='57'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastNonSpaceChar_09(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := " The cow jumped over the moon.    "

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, 34, 33)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_10(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := ""

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 0)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_12(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "The cow jumped over the dark side of the moon.   "

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, 49)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_13(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "The cow jumped over the dark side of the moon.   "

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, 49, 50)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_14(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "The cow jumped over the dark side of the moon.   "

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, -1, 47)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_15(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "The cow jumped over the dark side of the moon.   "

	_, err := StrOps{}.FindLastNonSpaceChar(tStr, 0, -1)

	if err == nil {
		t.Error("Expected an Error return.  NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastNonSpaceChar_16(t *testing.T) {
	//                 1         2         3         4         5         6         7         8
	//       012345678901234567890123456789012345678901234567890123456789012345678901234567890
	tStr := "Now is the time for all good men to come to the aid of their country.           "

	lastNonSpaceChar, err := StrOps{}.FindLastNonSpaceChar(tStr, 10, 62)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastNonSpaceChar(...). "+
			"Error='%v' ", err.Error())
	}

	if 62 != lastNonSpaceChar {
		t.Errorf("Error: Expected last non-space char idx='62'.  Instead, idx='%v' ",
			lastNonSpaceChar)
	}

}

func TestStrOps_FindLastSpace_01(t *testing.T) {

	//          012345678901234
	testStr := "xx       Hxgltx"

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 0, 14)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 0, 14). "+
			"Error='%v' ", err.Error())
	}

	if 8 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='8'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_02(t *testing.T) {

	//          012345678901234
	testStr := "xx1111111Hxglt "

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 0, 14)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 0, 14). "+
			"Error='%v' ", err.Error())
	}

	if 14 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='14'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_03(t *testing.T) {

	//          012345678901234
	testStr := " x1111111Hxgltf"

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 0, 14)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 0, 14). "+
			"Error='%v' ", err.Error())
	}

	if 0 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='0'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_04(t *testing.T) {

	//          012345678901234
	testStr := " x1111111Hxglt "

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 1, 13)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 1, 13). "+
			"Error='%v' ", err.Error())
	}

	if -1 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='-1'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_05(t *testing.T) {

	//          012345678901234
	testStr := " x1111   Hxgl  "

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 5, 12)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 5, 12). "+
			"Error='%v' ", err.Error())
	}

	if 8 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='8'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_06(t *testing.T) {

	//          012345678901234
	testStr := "fx1111rg3luHxgl"

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 0, 14)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 0, 14). "+
			"Error='%v' ", err.Error())
	}

	if -1 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='-1'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_07(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	lastSpaceIdx, err :=
		StrOps{}.FindLastSpace(testStr, 0, 14)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastSpace(testStr, 0, 14). "+
			"Error='%v' ", err.Error())
	}

	if 7 != lastSpaceIdx {
		t.Errorf("Error: Expected lastSpaceIdx='7'. Instead, Idx='%v' ",
			lastSpaceIdx)
	}
}

func TestStrOps_FindLastSpace_08(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	_, err :=
		StrOps{}.FindLastSpace(testStr, -1, 14)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_09(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	_, err :=
		StrOps{}.FindLastSpace(testStr, 0, -1)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_10(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	_, err :=
		StrOps{}.FindLastSpace(testStr, 0, 15)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_11(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	_, err :=
		StrOps{}.FindLastSpace(testStr, 19, 14)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_12(t *testing.T) {

	//          012345678901234
	testStr := "fx1111  3luHxgl"

	_, err :=
		StrOps{}.FindLastSpace(testStr, 5, 3)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_13(t *testing.T) {

	//          012345678901234
	testStr := ""

	_, err :=
		StrOps{}.FindLastSpace(testStr, 0, 0)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastSpace_14(t *testing.T) {

	//          012345678901234
	testStr := ""

	_, err :=
		StrOps{}.FindLastSpace(testStr, 0, 9)

	if err == nil {
		t.Error("Error: Expected an Error return from FindLastSpace(). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_FindLastWord_01(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          01234567890123456789012345678901234567890123456789012345678901234567890
	testStr := "Now is the time for all good men to come to the aid of their country.  "

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 0, 70)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 61 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='61'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 68 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='68'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_02(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          01234567890123456789012345678901234567890123456789012345678901234567890
	testStr := "Now is the time for all good men to come to the aid of their country.  "
	//                    xxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 10, 15)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 11 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='11'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 14 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='14'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_03(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country.  "
	//          xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 0, 72)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 63 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='63'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 70 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='70'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_04(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "                                                                         "
	//          xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 0, 72)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if -1 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='-1'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if -1 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='-1'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if true != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='true'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_05(t *testing.T) {
	//          0         1         2         4         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//          xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 0, 72)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 0 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='0'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 72 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='72'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if true != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='true'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_06(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "                xxxxxxxxxxxxxxxxxxxxxxxxxxxxx                            "
	//            xxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 2, 8)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if -1 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='-1'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if -1 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='-1'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if true != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='true'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_07(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country . "
	//                                                       xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 55, 72)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 71 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='71'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 71 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='71'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_08(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	begWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, 40, 67)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if 63 != begWrdIdx {
		t.Errorf("Error: Expected begWrdIdx='63'. Instead, begWrdIdx='%v' ", begWrdIdx)
	}

	if 67 != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='67'. Instead, endWrdIdx='%v' ", endWrdIdx)
	}

	if false != isAllOneWord {
		t.Errorf("Error: Expected isAllOneWord='false'. Instead, isAllOneWord='%v' ", isAllOneWord)
	}

	if false != isAllSpaces {
		t.Errorf("Error: Expected isAllSpaces='false'. Instead, isAllSpaces='%v' ", isAllSpaces)
	}

}

func TestStrOps_FindLastWord_09(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, -1, 67)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindLastWord_10(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, 40, -1)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindLastWord_11(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := "  Now is the time for all good men to come to the aid of their country "
	//                                                  xxxxxxxxxxxxxxxxxxxxxxxxxxxx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, 72, 72)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

func TestStrOps_FindLastWord_12(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          0123456789012345678901234567890123456789012345678901234567890123456789012
	testStr := ""
	//           xx

	_, _, _, _, err :=
		StrOps{}.FindLastWord(testStr, 1, 2)

	if err == nil {
		t.Error("Expected an error to be returned. NO ERROR RETURNED!")
	}

}

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

func TestStrOps_FindLastWord_15(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          01234567890123456789012345678901234567890123456789012345678901234567890
	testStr := "Now is the time for all good men to come to the aid of their country.  "

	startIndex := 67

	beginWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, startIndex, startIndex)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if startIndex != beginWrdIdx {
		t.Errorf("Error: Expected beginWrdIdx='%v'.\n"+
			"Instead, beginWrdIdx='%v'.\n",
			startIndex, beginWrdIdx)
	}

	if startIndex != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='%v'.\n"+
			"Instead, endWrdIdx='%v'.\n",
			startIndex, endWrdIdx)
	}

	if isAllOneWord != true {
		t.Error("Error: Expected isAllOneWord='true'.\n" +
			"Instead, isAllOneWord='false'.\n")
	}

	if isAllSpaces != false {
		t.Error("Error: Expected isAllSpaces='false'.\n" +
			"Instead, isAllSpaces='true'.\n")
	}
}

func TestStrOps_FindLastWord_16(t *testing.T) {
	//          0         1         2         3         4         5         6         7
	//          01234567890123456789012345678901234567890123456789012345678901234567890
	testStr := "Now is the time for all good men to come to the aid of their country.  "

	startIndex := 60 // A space character

	beginWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err :=
		StrOps{}.FindLastWord(testStr, startIndex, startIndex)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.FindLastWord(). \n "+
			"Error='%v' ", err.Error())
	}

	if startIndex != beginWrdIdx {
		t.Errorf("Error: Expected beginWrdIdx='%v'.\n"+
			"Instead, beginWrdIdx='%v'.\n",
			startIndex, beginWrdIdx)
	}

	if startIndex != endWrdIdx {
		t.Errorf("Error: Expected endWrdIdx='%v'.\n"+
			"Instead, endWrdIdx='%v'.\n",
			startIndex, endWrdIdx)
	}

	if isAllOneWord != false {
		t.Error("Error: Expected isAllOneWord='false'.\n" +
			"Instead, isAllOneWord='true'.\n")
	}

	if isAllSpaces != true {
		t.Error("Error: Expected isAllSpaces='true'.\n" +
			"Instead, isAllSpaces='false'.\n")
	}
}
