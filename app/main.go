package main

import (
  "MikeAustin71/stringopsgo/app/examples"
  "fmt"
  "time"
)

func main() {
  var startTime, endTime time.Time

  startTime = time.Now()

  time.Sleep(500 * time.Millisecond)

  endTime = time.Now()

  totalNanoSecs,
  elapsedTime := examples.MainTest{}.Timer(startTime, endTime)



  fmt.Println("totalNanoSecs", elapsedTime)
  fmt.Println("Elapsed Time", totalNanoSecs)

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
