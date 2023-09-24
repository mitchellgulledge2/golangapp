package main

import (
	"fmt"
	"reflect" // needed in order to use TAGS in structs
)

// creating struct outside main function (declaring struct)
// its really a list of field names with a type associated with the field
// struct has no constraints on data that can be inside struct unlike maps
// if we start variable with structure with upper case it will be available to all packages
// if the first letter of struct is lower case it is only available within local package
type exampleStruct struct {
	number int
	name string
	companions [] string // this is a slice within a struct
}

type Animal struct {
	// structures also support the concept of TAGS
	// to use a tag for instance, a max length count for name
	Name string `required max: "100"`
	Origin string
}

type Bird struct {
	// embed a struct within a struct
	// embed Animal struct within Bird struct
	Animal
	SpeedKPH float32
	CanFly bool
}

func main() {
	// Maps and Structs are 2 collection types in GO
	// MAP example below, takes 2 arguements to make key value pair
	// in this map example we have string data type for key and int for value
	mapExampleVariable := map[string]int{
		"Dorothy": 19,
		"Mitchell": 49,
	}
	// the above example is great when we need to map one key type to one value type

	// some limitations like an array is a valid key type but a slice is not
	m := map[[3]int]string{}
	fmt.Println(m)
	// above way is a literal syntax

	// you can also use make function to create a map
	mapExampleVariableTwo := make(map[string]int)
	fmt.Println(mapExampleVariableTwo)

	// you can print the values for specific keys in a map
	fmt.Println(mapExampleVariable["Dorothy"])

	// you can also add keys and values to existing maps
	mapExampleVariable["Maggie"] = 5

	// return order of a map is not guaranteed unlike a slice or array
	// entries in map will almost always be different/random

	// built in delete functions for removing entries from a map
	delete(mapExampleVariable, "Mitchell")
	fmt.Println(mapExampleVariable)

	// to validate if a key is inside a map
	// this is done using the comma ok syntax , ok
	pop, ok := mapExampleVariable["Buttons"]

	fmt.Println(pop, ok)
	// good way to check if key is in map

	// you can use len to find length of map
	fmt.Println(len(mapExampleVariable))

	// manipulating one map that references another map will influence all references to those maps

	// moving onto structs and this will be most relevant for dealing with jsons
	// example struct value
	aEntry := exampleStruct{
		number: 3,
		name: "Mitchell",
		companions: []string{
			"Babe",
			"Slam",
		},
	}
	// you can also pull specific values for a reference name within a struct
	fmt.Println(aEntry.companions[1])

	// another way to create a struct (anonymous struct)
	// defining structure of struct and second {} is entry in struct
	aExampleStruct := struct{name string}{name: "corn"}
	fmt.Println(aExampleStruct)
	// used for short lived data in web apps to organize API responses...

	aStructEntry := aExampleStruct
	aStructEntry.name = "Penny"
	fmt.Println(aStructEntry) // values for name is updated to Penny in new struct
	fmt.Println(aExampleStruct) // still stays as Corn
	// each struct is stored in memory as independent data set
	// meaning changing one that references another will not affect both

	// if we want to point to same reference of underlying data we can use pointers
	// (&)
	aaStructEntry := &aExampleStruct
	fmt.Println(aaStructEntry) 

	// GO uses composition through embedding
	bStruct := Bird{}
	bStruct.Name = "Dorothy"
	bStruct.Origin = "Taiwan"
	bStruct.SpeedKPH = 44
	bStruct.CanFly = false
	fmt.Println(bStruct.Name)

	cStruct := Bird{
		// explicitly need to define internal Animal strut
		Animal: Animal{Name:"Babe", Origin: "USA"},
		SpeedKPH: 40,
		CanFly: true,
	}
	fmt.Println(cStruct.Name)

	// using tags within structs
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	// asking for TAG property of field in struct below
	fmt.Println(field.Tag)
	// all TAGS do is provide a string of text, nothing more
	// if tag is there apply this logic

	// maps are collections of value types that are accessed via keys
	// created via literals or make function
	// members accessed via [key] syntax / myMap[key] = "value"
	// maps check for presence with "value, ok" form of result
	// multiple assignments of map refer to same underlying data
	
	// structs are collections of disparate data types to describe single concept
	// keyed by named fields
	// normally created as Types, but anonymous structs are allowed
	// Structs are value types
	// No inheritence, but can use embedding (one struct within another)
	// Tags used to describe fields within structs

}
