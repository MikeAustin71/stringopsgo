package main

import (
	"MikeAustin71/stringopsgo/app/examples"
	"fmt"
	"time"
)

func main() {

	var startTime, endTime time.Time
	startTime = time.Now()
	time.Sleep(50 * time.Millisecond)
	endTime = time.Now()

	totalNanoSeconds,
	elapsedTime := examples.MainTest{}.Timer(startTime, endTime)

	if totalNanoSeconds > 5000 {
		return
	}

	if elapsedTime != "Something" {
		return
	}

fmt.Println()

}












/* This Works
import (
  "MikeAustin71/stringopsgo/strops/v2"
  "fmt"
  "strings"

)

Reference format:
 strops.StrOps{}.ExtractNumericDigits(..)

Check-Out
https://goreportcard.com
import (
  "MikeAustin71/stringopsgo/strops/v2"
  "errors"
  "fmt"
  "io"
  "os"
  "regexp"
  "sort"
  "strconv"
  "strings"
  "time"
)

*/
