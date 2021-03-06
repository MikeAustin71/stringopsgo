# Running Tests

Open a command prompt in this directory (*stringopsgo/strops/v2*) and run the
following commands.

##### Windows Command
    'go test -v > xx_tests.txt`

##### Linux Command
    'go test -v | tee xx_tests.txt`

This will generate test results in the *stringopsgo/strops/v2* 
directory which are stored in the text file, `xx_tests.txt`. 

## Running Tests with code coverage

First pull down and install the `cover` package.
 
  `go get golang.org/x/tools/cmd/cover`
  
Next, follow the test execution protocol.  
  
## Test Execution with Code Coverage
Run this in *strops/v2* directory:

##### Windows Command
 `go test -cover -v > xx_tests.txt`  

##### Linux Command
  `go test -cover -v | tee xx_tests.txt`     

## Cover Profile

Generate the code coverage detail. Run this command
in the *stringopsgo/strops/v2* directory:

`go test -coverprofile=xx_coverage.out`


The following provides for code coverage display in your
browser. Run this on the terminal command line and run it
in the *stringopsgo/strops/v2* directory:

`go tool cover -html=xx_coverage.out`