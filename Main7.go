package main

import (
	"fmt"
)

func main() {
	a := 42
	// this is a data copy, different data sets in memory
	b := a
	fmt.Println(a, b)

	// to point to same underlying data in memory use a pointer
	var aa int = 42
	// we initialize with *int and &aa
	// telling bb not to hold the value of aa but rather the memory location
	var bb *int = &aa
	// to figure out what value is in the memory location (*bb)
	// called dereferencing
	aa = 66
	fmt.Println(aa, *bb)

	// initializing array of 3 integers
	d := [3]int{1,2,3}
	// using pointers to reference indexed values of array to set to value of new variable
	e := &d[0]
	fmt.Printf("%v %p\n", d, *e)
	// math operators cannot be done on pointers

	// initializes struct that is defined outside main func
	// using dereferencing (*) for myStruct to get value located in memory vs reference value
	// * makes it pointer to myStruct
	var ms *myStruct
	// using object initialization syntax
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// you can use built in new function
	var mss *myStruct
	mss = new(myStruct)
	// need to dereference pointer in order to get value
	mss.foo = 55
	// same thing as (mss).foo
	fmt.Println(mss.foo)
	// the zero value for a pointer is nil

	// initializing array
	v := [3]int{1,2,3}
	// setting variable equal to array
	vv := v
	// printing both values
	fmt.Println(vv, v)
	// setting value for indexed element in array
	vv[1] = 5
	// printing both values
	fmt.Println(vv, v)
	//[1 2 3] [1 2 3]
	//[1 5 3] [1 2 3]


	// initializing slice
	kv := []int{1,2,3}
	// setting variable equal to slice
	kvv := kv
	// printing both values
	fmt.Println(kvv, kv)
	// setting value for indexed element in slice
	kvv[1] = 5
	// printing both values
	fmt.Println(kvv, kv)
	//[1 2 3] [1 2 3]
	//[1 5 3] [1 5 3]

	// a slice is a projection of an underlying array
	// slice is actually a pointer to an array
	// when sharing slices in application we are always pointing to same underlying data

	// initializing map with key value pair of string string
	abc := map[string]string{"foo": "bar", "ass": "grass"}
	bcd := abc
	fmt.Println(abc, bcd)
	// modifying abc affects bcd
	abc["foo"] = "go"
	fmt.Println(abc, bcd)
	// when working with slices and maps, data can change in unexpected ways

	// creating pointers with asterisk (*) - *int is ponter to integer
	// use address of operator (&) - get value of existing variable in memory
	// dereferencing pointers - (*) in front of pointer to drill down to value of pointer
	// Complex types (ex: structs) are automatically dereferenced
	// create pointers with objects - ex: ms := myStruct{foo:42} 
	// p := &ms <- address of operator used here

	// another way to create pointer to object:
		// use addressof operator befor initializer - &myStruct{foo: 42}
		// use the new keyword (cant initialize fields at same time)
	// Types with internal pointers
		// all assignment operations in GO are copy operations
		// slices and maps contain internal pointers so copies point to same underlying data



}

// defining struct outside main function
// struct is called myStruct
type myStruct struct {
	foo int
}
