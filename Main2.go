package main

import (
	"fmt"
	// commenting out the below two libraries as they are used for making api calls
	//"net/http"
	//"io/ioutil"
)

func main() {
	// creating array below
	// [3] being size of the array so 3 entries
	// int being the type of data in the array
	// an array can only store one type of data
	// array of integers that has 3 entries
	// initializer syntax goes after where you enter values
	sample := [3]int{1,3,5}
	fmt.Printf("%v\n",sample)
	// technically you can replace the size of the array with [...] 
	// this makes array size match entries
	sampleTwo := [...]int{2,4,6}
	fmt.Printf("%v\n",sampleTwo)

	// creating new variable that is an array of strings
	var sampleThree [3]string
	// printing array but since no values have been assigned it will be blank array
	fmt.Printf("%v\n", sampleThree)
	// assigning value to index of array
	sampleThree[0] = "BABE"
	fmt.Printf("%v\n", sampleThree)

	// arrays are made up of continuous blocks of memory
	// GO has pointer (remember location) of beginning of array
	// you can index arrays just like python
	fmt.Printf("%v\n", sampleThree[0])

	// you can use built in function len (length) to find length of array
	fmt.Printf("Length of array: %v\n", len(sampleThree))

	// arrays can be made with more than just primitives
	// creating array with non primitives (array of arrays)
	// [[100][010][001]]
	var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// another way to write the above is to create the array and then add individual arrays after
	var identityMatrix2 [3][3]int 
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{1, 1, 0}
	identityMatrix2[2] = [3]int{1, 0, 1}
	fmt.Println(identityMatrix2)
	// much easier to read

	// in GO arays are considered values
	// when you copy an array, we are creating literal copy
	// meaning GO reassigns the entire length of the array
	aArray := [...]int{1,2,3}
	bArray := aArray
	bArray[1] = 5
	fmt.Println(aArray)
	fmt.Println(bArray)

	// POINTERS can be utilized if you do not want to create copies of Arrays
	// so instead of bArray := aArray
	// you can use bArray := &aArray
	
	// downside to arrays is that they have a fixed size that has to be calculated at compile time

	// moving into slices to get around limitations of arrays
	// creating slice below
	// slice is initialized as a literal by using square brackets []
	// type of data and some initialized data
	aSlice := []int{1,2,3}
	fmt.Println(aSlice)

	// anything we can do with an array we can do with a slice as well
	// for example the len function can be used on slices as well
	// in slices you can also use the built in capacity function cap
	fmt.Printf("capacity: %v\n", cap(aSlice))
	// a slice is a projection of an underlying array

	// slices are reference types
	bSlice := aSlice
	bSlice[1] = 5
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	// so unlike arrays when we set a slice = to another, it is referencing the same underlying data
	// changing b also affects aArray

	aaSlice := []int{1,2,3,4,5,6,7,8,9}
	bbSlice := aaSlice[:] // slice of all elements
	ccSlice := aaSlice[3:] // slice 4th element to end
	ddSlice := aaSlice[3:6] // slice 4th 5th 6th elements
	fmt.Println(aaSlice)
	fmt.Println(bbSlice)
	fmt.Println(ccSlice)
	fmt.Println(ddSlice)
	// note changing one element in any of these slices will affect all the other slices 
	// since they are referencing same underlying data

	// GO has a built in function for slices called make
	// takes 2 or 3 arguments (3rd arguement being capacity)
	fSlice := make([]int, 3, 100)
	fmt.Println(fSlice)
	fmt.Printf("length of slice: %v\n", len(fSlice))
	fmt.Printf("capacity of slice: %v\n", cap(fSlice))

	// just like python you can append eleents to a slice
	fSlice = append(fSlice, 1)
	fmt.Println(fSlice)
	fmt.Printf("capacity of slice: %v\n", cap(fSlice))
	// so now slice is [0 0 0 1] with capacity of 100

	// the append function can take 2 or 3 arguements
	// when inputting the 3rd arguement it is called a variatic function
	fSlice = append(fSlice, 1, 2, 3, 4, 5)
	fmt.Println(fSlice)
	// everything input past the first arguement will be seen as elements to append to slice
	// when not setting capacity manually, GO will calculate the size itself (try not to)
	// a slice with a data type of integers can only append same data type
	// you cannot append a slice of integers inside a slide of integers
	// example what you cant do a = append(a, []int{1,2,3})
	// to get around this you can add ... to spread the slice out to individual arguements
	fSlice = append(fSlice, []int{1,2,3}...)
	fmt.Println(fSlice)

	// stack operations are also available for slices
	// used for popping elements off slice
	// shift operation (indexing) can be used to trim elements from beginning and end of slice
	ffSlice := fSlice[1:len(fSlice)-2]
	fmt.Println(ffSlice)

	// in order to remove item from the middle of a slice, 
	// you have to concatanate two slices together
	// first slice will be all elements leading up to element
	// second slice will be everything past the pooped element to end
	// dont forget to use spread operation ...
	fffSlice := append(fSlice[0:3], fSlice[6:]...)
	fmt.Println(fffSlice)
	// if you remove elements from middle of slice, make sure you have no references to fSlice
	// this is because we are referencing the same object in memory so it will affect any func 

	// arrays are collections of items with same type
	// arrays are fixed in size, specify at compile time how big array is
	// 3 ways to delcare arrays
		// a := [3]int{1,2,3}
		// a:= [...]int{1,2,3}
		// var a [3]int
	// len fun returns size of array
	// copies refer to different underlying data

	// slices are backed by array
	// can create slice from existing array or slice
	// you can also create slice literal style
	// you can create slice from make function
		// a := make([]int, 10, 100)
		// length of 10 and capacity of 100
	// len func returns length of slice
	// cap func returns capacity of slice
	// append func adds elements to slice (see above examples)
	// unlike arrays, copies of slices refer to same underlying data

	


	// messing around with api calls
	//response, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json") //aws
	//response, err := http.Get("https://www.gstatic.com/ipranges/cloud.json") //gcp

	//fmt.Printf("%v\n", response)
	//if err != nil {
    //    fmt.Print(err.Error())
	//}

	// read response body
	//body, error := ioutil.ReadAll(response.Body)
	//if error != nil {
	//	fmt.Println(error)
	//	}
	// close response body
	//response.Body.Close()
   
	// print response body
	//fmt.Println(string(body))

}
