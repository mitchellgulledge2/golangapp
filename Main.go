package main

import (
	"fmt"
// in GO a string is just an alias for a stream of bytes
// to do this you need to import the strconv package
	"strconv"
)

// declaring package variable outside main function
// using var block () to clean up format
var (
	mitchellAGE int = 42
	geniusMITCHELL string = "Mitchell"
)

// enumerated constant example
const (
	// initial value of iota is the same as a 0 value for an integer
	aConstant = iota // iota is a counter we can use when making enumerated constants
	bConstant = 1 << iota // using bitshifting with iota
	cConstant = iota

	// there is a write only variable that you can use in GO by typing _
	// we dont need to store this value in memory
	_ = iota + 5
	// value of constant must be able to be determined at compile time
	// you can do things like addition (used for fixed offsets)

)

// lower case variables are scoped to the package
// current package is main
// if we switch a variable to an uppercase letter I instead of i
// then all packages can access and use the variable

func main() {
	// when declaring variables in main function they must be used
	// this is true when you are using them within main func, this is considered Blocked scope
	// 3 different ways to declare variables in GO
	// way 1, start with var key word, name of variable, and type of variable
	var i int //decalred variable
	i = 42 // set variable value

	// another way to initialize a variable
	// this is actually REDECLARING a variable that was already created
	// overwriting the package level value of a variable inside main function is called shadowing
	var mitchellAGE int = 42
	// var k float32 = 27  
	// float32 and float64 both exist

	// final way to declare variable
	j := 42.
	// can redeclare a variable by just setting j equal to 43
	j = 43

	// need to use all variables in order to not get an error
	fmt.Println(j)

	// another way to print formatted strings in GO
	// prints variable and type (%T)
	fmt.Printf("%v, %T", geniusMITCHELL, geniusMITCHELL)

	// below we are going to convert from one variable type to another
	fmt.Printf("%v, %T\n", i, i)

	var e float32
	// below is a conversion function (float32 is the func here)
	// destinationType(variable)
	e = float32(i)
	fmt.Printf("%v, %T\n", e, e)

	// going to convert integer to string
	// mitchellAGE is a int set at 42
	// going to convert this to a string
	var t string
	t = strconv.Itoa(mitchellAGE)
	fmt.Printf("%v, %T\n", t, t)

	// summarizing three ways to declare variables
	// var foo int
	// var foo int = 42
	// foo := 42

	// summarizing three different levels of variable visibility
	// lower case = package scope
	// upper case first letter = global
	// blocked scope = w/in function only

	// going to start focusing on PRIMITIVES
	// Boolean
	// Numeric types (integers, floating point, complex numbers)
	// Text Types

	// Boolean is just true or false
	var booleanEXAMPLE bool = true
	fmt.Printf("%v, %T\n", booleanEXAMPLE, booleanEXAMPLE)

	// using equal operator == to see if value equals one another
	n := 11 == 22
	fmt.Printf("%v, %T\n", n, n)
	// you dont have to initialize a variable value in GO, all values default with 0 or false

	// now moving to numeric types
	// the 0 value for numeric types is....0
	// int8 ranges from -128 - 127
	// int16 ranges from -32768 - 32767
	// int32 ranges in the 2m mark
	// int64 when you need really big number

	// uint8, uint16, uint32

	// there are signed and unsigned integers, they are basically the same type
	var unsignedINT uint16 = 42
	fmt.Printf("%v, %T\n", unsignedINT, unsignedINT)

	aa := 10
	bb := 3
	fmt.Println(aa - bb)
	// +, -, *, /, % percent sign being used for remainder
	// primitive types cannot change during operation
	// integer divided by integer will equal integer

	// we also cannot add integers of different types int + int8
	var cc int = 10
	var dd int8 = 3
	// need to convert one variable to match the other primitive
	fmt.Println(cc + int(dd))

	// there are also 4 bit operators, & | ^ &^
	// like subnetting _ _ _ _ 8 4 2 1
	// so cc which is equal to 10 is actually 1010
	// and dd which is 3 is equal to 0011
	// example below uses and operator which looks for matches
	// what bits are set in first and second number
	// so matching 1010 and 0011 you get 0010 which is 2
	fmt.Println(cc & int(dd))

	// or operator looks for one or the other bit being set btwn 2 ints
	fmt.Println(cc | int(dd)) // 1101 which eqals 11

	// exclusive or
	fmt.Println(cc ^ int(dd)) // 1001

	// and not
	fmt.Println(cc &^ int(dd)) // 0100

	// example bit shifting
	xx := 8 // 8 = 2^3 
	fmt.Println(xx << 3) // bit shift xx left 3 places 
	// 2^3 * 2^3 = 64

	fmt.Println(xx >> 3) // bit shift xx right 3 places
	// 2^3 / 2^3 = 1

	// in order for GO to store decimals we need floating point
	// there are 32 and 64 bit floating numbers
	exampleFLOATINGpoint := 3.14
	fmt.Println(exampleFLOATINGpoint)
	// cant do operations between float64 and float32s
	// default value for floats is float64

	// complex primitive type example
	var nn complex64 = 1 + 2i
	// + - * / all available
	fmt.Printf("%v, %T\n", nn, nn)

	// there are 2 built in functions in GO called real and imag
	// real pulls real values from complex variables (1)
	// imag pulls fake values from complex variables (2i)
	fmt.Printf("%v, %v\n", real(nn), imag(nn))

	// there is also a complex function built into GO
	// This function takes in 2 numbers, see example below
	var yy complex128 = complex(5, 12)
	// the first number is the real part and the second number is the imaginary part
	fmt.Printf("%v, %T\n", yy, yy)

	// next primitive data type is text
	// string is a text type
	// string in GO is any utf8 character
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	// we can treat string of text as collection of letter
	// meaning we can index characters in a string
	fmt.Printf("%v, %T\n", s[2], s[2])
	// strings in GO are aliass for bytes, so i = 105
	// need to convert to string
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	// strings are immutable

	// strings can be concatantated
	s2 := " this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// you can also convert strings into a collection of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
	// a lot of functions in GO work with bytes
	// going to convert strings to bytes a lot

	// rune represents a utf32 character
	// any utf8 character can be a valid utf32 character
	// when declaring a rune you use single quotes instead of double quotes
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
	// data type comes back int32

	// primitive data type summary
	// boolean type = true or false
	// numeric types (integers)
		// integers (signed)  (int8 through int64)
		// integers (unsigned) (8 bit uint8 - uint32)
		// arithmetic operations (+, -, *, /, %)
		// bitwise operators (and, or, xor and not)
		// zero value is 0
	// numeric types (floating point numbers)
		// zero value is 0
		// 32 and 64 bit versions 3.14, 13.7e5 etc
		// + , - , * , /
	// numeric type (complex numbers)
		// zero value is (0+0i)
		// 64 and 128 bit versions
		// built in functions 
			// complex - make complex number from 2 floats
			// real - get real part of float
			// imag - get imaginary part of the float
		// + , - , * , /
	// Text Types (strings) 
		// utf8 and immutable
		// can be concatenated +
		// can be converted to []byte
	// Text Type (Rune)
		// UTF32 alias for int32

	// Moving onto Constants in GOLANG
	// creating constant below
	const myConst int =42
	// typed constant created by declaring int
	fmt.Printf("%v, %T\n", myConst,myConst)
	// contants can be int, str, float32, bool etc
	// constants can be shadowed just like variables can
		// meaning we can create constant outside function and redefine value inside function
		// inner declaration constant value always wins

	var bc int = 27
	// adding constant and variable
	// do not have to declare the constant type, GO can infer the constant type int at compile time
	fmt.Printf("%v, %T\n", myConst + bc, myConst + bc)

	// printing enumerated constants created outside function
	// noting that since the below constants were made outside main function they can be shadowed
	fmt.Printf("%v\n", aConstant) // iota set at 0
	fmt.Printf("%v\n", bConstant) // iota infers 1 after iota = 1 (because in same constant block ())
	fmt.Printf("%v\n", cConstant) // iota infers 1 after iota = 2 (because in same constant block ())

	// setting variable equal to constant 
	var testVariable int = aConstant
	fmt.Printf("%v,\n", testVariable== aConstant)

	const kk = 1 << (10 * iota) // you can use bitshifting with iota
	fmt.Printf("%v,\n", kk)

	// you can use bit shifting to set boolean flags inside of a single byte
	// OR'ing together 3 different constants in same constant flag using iota
	var roles byte = aConstant | bConstant | cConstant
	fmt.Printf("%b\n", roles)

	fmt.Printf("Is True? %v\n", aConstant == bConstant)

	// constants are immutable but can be shadowed
	// constants are replaced by compiler at compile time
		// value must be calculable at compile time
	// Named just like variables
	// Typed constants work just like immutables
		// can only operate with same type
	// untyped contants work like literals
		// can interoperate with similar types

	// enumerated constants
		// iota allows related constants to be easily created
		// iota starts at 0 in each constant block and increments by 1
	// enumerated expressions
		// operations can be determined at compile time (+, -, *, /, >>, <<, and, or, xor and not)
	
}
