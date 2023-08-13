# Chile Rut Tools
Go package that provides tools to format and validate chilean's RUN and RUT, the unique identification number for citizen and companies using the Module 11 algorithm

# Installation
```
go get github.com/ppastene/chile-rut-tools
```

# Usage
```
import github.com/ppastene/chile-rut-tools
```

## Get Verifier Digit
Will return the verifier digit in string format in a range of 0 to 9 including the letter K
```
ruttools.GetDigit("10118681") 4
ruttools.GetDigit("64643847") 0
ruttools.GetDigit("19255375") K
```

## Validate RUT
Will return true if the rut is valid, false otherwise
```
ruttools.IsValid("12693547-1") true
ruttools.IsValid("20973510-9") false
ruttools.IsValid("47704154-K") true
```

## Format RUT
Will return the rut with the correct format removing any unwanted character
```
ruttools.Format("99-030-307-K") 99030307-K
ruttools.Format("        12+488+882+4") 12488882-4
ruttools.Format("00000000849373080") 84937308-0
```
