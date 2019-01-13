package main

import (
	strOps "./strops/v2"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("main() - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strOps.StrOps{}.NewPtr()
	sops1.StrOut = originalStr

	p := make([]byte, 100)

	n, err := sops1.Read(p)

	if err != nil && err != io.EOF {
		fmt.Printf("Error returned by sops1.Read(p). "+
			"Error='%v' \n", err.Error())
		return
	}

	sops2 := strOps.StrOps{}.NewPtr()

	n2, err := sops2.Write(p)

	if err != nil && err != io.EOF {
		fmt.Printf("Error returned by sops2.Write(p). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("        Original Str: ", originalStr)
	fmt.Println(" Original Str Length: ", len(originalStr))
	fmt.Println("        sops1.StrOut: ", sops1.StrOut)
	fmt.Println(" sops1.StrOut Length: ", len(sops1.StrOut))
	fmt.Println("    sops1 Bytes Read: ", n)
	fmt.Println("        sops2.StrOut: ", sops2.StrOut)
	fmt.Println(" sops2.StrOut Length: ", len(sops2.StrOut))
	fmt.Println(" sops2 Bytes Written: ", n2)
	fmt.Println("**********************************************")

}

func ExampleIoCopy01() {
	fmt.Println("ExampleIOCopy_01() - Version 2")

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
