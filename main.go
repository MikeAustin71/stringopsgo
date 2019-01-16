package main

import (
	strOps "./strops/v2"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	ExampleWrite02()

}

func ExampleWrite03() {

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

func ExampleWrite01() {

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

func ExampleWrite02() {

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

func ExampleRead01() {

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

func ExampleRead02() {
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

func ExampleIoCopy02() {
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
