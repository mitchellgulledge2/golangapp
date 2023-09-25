package main

import (
	"fmt"
)

// creating example function
// it will return a boolean result
func returnTrue() bool{
	fmt.Println("returning true")
	return true
}

func main() {
	stateDb := map[string]int{
		"entry_one": 1,
		"entry_two": 2,
	}
	// if statment to search for a key in the map
	// if key found (ok) then print the value for the key
	if pop, ok := stateDb["entry_one"]; ok {
		fmt.Println(pop)
	}
		// you can only use the pop variable within if statement
		// you cannot use pop statement anywhere else in main func

	// we have compare operators we can use with if statements
	// <, >, ==, <=, >=, !=
	number := 50
	guess := 30
	// use logical operators to add additional validations
	// using or operator to ensure variable is within range
	if guess <1 || guess >100 {
			// GO has the concept of short circuiting
			// if one condition fails when multiple are listed then GO doesnt continue
		fmt.Println("guess is outside range 1-100")
	} else if guess >=1 && guess <=100{
	// using and operator to ensure variable is not outside range
		if guess < number {
		fmt.Println("guess is less than number")
		// executing function to return True
		fmt.Println(returnTrue())
		}
	} else {
		fmt.Println("example")
	}

	// moving onto switch statements
	// special purpose if statement
	// use switch keyword
	switch i := 2+3;i { // using initializers to create tag and switch on that
	// first case that passes is one that executes
	case 1, 5, 6:
		fmt.Println(1)
	// multiple tests can be in a single case
	case 2:
		fmt.Println(2)
	// all test cases must be unique
	default:
		fmt.Println("None")
	// easiest way to compare one variable to multiple possibilities
	}

	ii := 10 // establishing variable
	switch{ // stand alone switch statement with no tag
	case ii <= 10: // cases standing in as if statements
		fmt.Println("less or equal to 10")
		// we can use the concept of fallthrough if we dont want to break
		fallthrough
	case ii >= 10:
		fmt.Println("greator or equal to 10")
		// because fallthrough statement this also executes
		// fallthrough is stateless so statement after will execute no matter t/f
	default:
		fmt.Println("neither")
	}

	// type switch statements below
	// interface type can be assigned to integer, string, bool, func etc
	var iii interface{} = 1
	// dot operator and tell go to pull underlying data type
	switch iii.(type) {
	// using GO types as cases
	case int:
		fmt.Println("int")
		// you can exit case early with break keyword
		// break
		fmt.Println("this also prints")
	default:
		fmt.Println("not an int")
	}
}

// note that when doing equality operations with floating point numbers be careful
