package main

import (
	"fmt"
	"log"
)

func main() {
	a:= 1
	// typically program is executed from top to bottom
	fmt.Println(a)
	defer fmt.Println(a) // executes after main function exits
	// doesnt move at end of main function but after main function
	fmt.Println("3rd")
	// defer function operates in a last in first out

	// typical use case for defer is with http calls
	// defer res.body.Close()
	// body may need to stay open for awhile during main function
	// otherwise youll get error that you tried to read the response body after closing

	// there are no exceptions in GO, you can use panic
	// we return error values not exceptions
	a, b := 1,1 // changed from 1/0
	ans := a/b
	//panic("something bad")
	fmt.Println(ans)
	// panic: runtime error: integer divide by zero

	// panics happen after defer statements are executed after main func
	// defer statements will succeed even if application panics
	// executing panicker func
	panicker()
}

func panicker(){
	// example anonymous func (aka func without name)
	defer func() {
		// recover func will return nil if func is okay
		if err := recover(); err != nil {
			log.Println("Error", err)
			// you can also re-panic your application
			panic(err) // this makes main func panic
		}
	}()// these empty () make anynymous func execute
	panic("something wrong")
	fmt.Println("end")

	// defer func to delay execution of statement until end of func
	// defer func are useful to group func together to open and close
	// defer statements run in last in first out order
	// defer arguements are evaluated at time defer is executed, not at time func is called

	// when go encounters an error, like a API call and getting no response back
	// 404 will return 
	// panic happens when GO has to recover an application (cant continue)
	// use for things like cannot obtain TCP port from web server
	// if func panics, it will stop executing and deferred statements will start
	// if nothing handles panic, program will exit

	// built in recover fun is used to recover from panics
	// only useful in deferred func

}

