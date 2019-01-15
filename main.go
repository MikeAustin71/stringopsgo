package main

import (
	strOps "./strops/v2"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	ExampleIoCopy02()

}

func ExampleRead01() {

	fmt.Println("ExampleRead01() - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strOps.StrOps{}.NewPtr()
	sops1.StrOut = originalStr

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
	fmt.Println("         sops1.StrOut: ", sops1.StrOut)
	fmt.Println("  sops1.StrOut Length: ", len(sops1.StrOut))
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
	sops1.StrOut = originalStr

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
	fmt.Println("        sops1.StrOut: ", sops1.StrOut)
	fmt.Println(" sops1.StrOut Length: ", len(sops1.StrOut))
	fmt.Println("    sops1 Bytes Read: ", n)
	fmt.Println("**********************************************")

}

func ExampleIoCopy02() {
	fmt.Println("ExampleIOCopy_02() - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strOps.StrOps{}.NewPtr()
	sops1.StrOut = originalStr
	sops2 := strOps.StrOps{}.NewPtr()

	n, err := io.Copy(sops2, sops1)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(sops2, sops1). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("       Original Str: ", originalStr)
	fmt.Println("Original Str Length: ", len(originalStr))
	fmt.Println("       sops1.StrOut: ", sops1.StrOut)
	fmt.Println("       sops2.StrOut: ", sops2.StrOut)
	fmt.Println("sops2.StrOut Length: ", len(sops2.StrOut))
	fmt.Println("      Bytes Written: ", n)
	fmt.Println("**********************************************")
	fmt.Println("      Copying sops2 To StdOut")
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
