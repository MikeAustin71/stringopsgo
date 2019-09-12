# Release Notes Package *strops* Version 2.0.4

## New Methods
Added methods:
  + StrOps{}.ExtractDataField()
  + StrOps{}.ExtractNumericDigits()
  
## Tests
  + Added tests for new methods.
  + Test coverage == 96%
  + Executed tests successfully on *Linux Mint 19.2* and 
  *Ubantu 18.04.3*. 
  
  

# Release Notes Package *strops* Version 2.0.3

## New Methods 
Added methods:
  + StrOps{}.StripBadChars()
  + StrOps{}.StripLeadingChars()
  + StrOps{}.StripTrailingChars()
  
## Tests
Added more tests. The *strops/v2* package now includes 210 unit
tests with a code coverage of 87%. 

## Module Requirements
All Version 2+ releases support *Go* modules.
With this release, *go.mod* now requires *Go*
version 1.12 or greater. 
  