package main

import (
  "fmt" )

func main() {

  fmt.Println("Testing main")
  lib := StrOps{}
  versionNo := lib.GetSoftwareVersion()

  fmt.Printf("Version Number= %v", versionNo)
}
