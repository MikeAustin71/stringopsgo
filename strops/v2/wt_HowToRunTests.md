# Running Tests
##### Open a command prompt in this directory (strops/v2) and run the following command:

### `go test -v > xx_tests.txt`

This will generate test results in the "strops/v2" directory.  The tests utilize
some library routines stored in the "appLibs" directory.

## Running Tests with code coverage

First pull down and install the `cover` package.
 
  `go get golang.org/x/tools/cmd/cover`
  
Next, follow the test execution protocol.  
  
## Test Execution with Code Coverage
Run this in *strops/v2* directory

 `go test -cover -v > xx_tests.txt`  
     

## Cover Profile

Generate the code coverage detail.
Run this in *strops/v2* directory:


`go test -coverprofile=xx_coverage.out`


The following provides for code coverage display in your
browser. Run this on the terminal command line and run it
in the *strops/v2* directory:

`go tool cover -html=xx_coverage.out`