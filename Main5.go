package main

import (
	"fmt"
)

func main() {
	// there is only one way we can loop in Go
	// For statement
		// simple loops
		// exiting early
		// looping through collections
	// creating loop with 3 statements
	// first statement is initializer
	// second statement generates a boolean result
	// last statement is incrementer
	//for i := 0; i < 5; i++ {
	
	// going to initiate multiple variables
	for i, j := 0, 0; i <5; i, j = i+1, j+1{
		fmt.Println(i, j)
	}

	// dont have incrementer within loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// below if statement checks to see if i is even or odd number
		if i%2 == 0 {
			i/= 2
		} else {
			i = 2*i + 1
		}
	}

	// we dont need all 3 statements in for loop
	// first statement can be empty if variable is already initialized
	ii := 0
	//for ; ii<5; ii++{
	// we can even leave out incrementor value
	//for ; ii<5; {
	// you can even leave out both ;
	for ii<5{
		fmt.Println(ii)
		ii++
	}

	// there is also infinite for loop
	// for {
	// need way to break from loop
	// you can use break keyword 
	// continue statement can also be used
	for iii := 0; iii <5; iii++ {
		if iii%2==0{
			continue
			// only print odd numbers less than 5
		}
		fmt.Println(iii)
	}

	// nested loop
	// for statement within a for statement
	Loop: // label for loop to reference later in break statements
		for k := 1; k<=3; k++{
			for y := 1; y <=3; y++{
				fmt.Println(k * y)
				// want to leave loop as soon as we get a value that is greater than 3
				if k * y >= 3{
					//break // only breaks out of inner loop 
					// to break outer loop use the Loop label
					break Loop
					
				}
			}
		}
	
	// looping through collections with for statements
	s := []int{1,2,3} // initializes slice of 3 integers
	// this can also be used for arrays, maps etc
	fmt.Println(s)
	// going to iterate through the slice using a for range loop
	// you will always get a key, value result when iterating
	for k, v := range s {
		fmt.Println(k, v)
	}

	// you can also perform for loops over channels in GO

	// For statement for all loops
	// simple loops
		// for initializer; test; incrementer {}
		// for test {}
		// for ()
	// 3 ways to exit for loop early
		// break, continue, labels
	// looping over collections
		// arrays, slices, maps, strings, channels
		// for k,v := range collection {}
}

