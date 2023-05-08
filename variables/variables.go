package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "string variable"
	fmt.Println(name)

	var number = 1
	fmt.Println(number)

	var ifTrue bool //if not initialized, bool variables are initialized to false, String to empty string and int to 0.
	fmt.Println(ifTrue)

	pi := 3.14 // variables can be declared in this format without specifying their type. Variable type is defined by the expression.
	fmt.Println(pi)
	fmt.Println(reflect.TypeOf(pi)) //get type of variable

	// pi = "some-string" will error out as pi is already set to float64

	x := 1
	ptr := &x // pointer variable contains the address of a variable.
	fmt.Println(ptr)
	fmt.Println(*ptr) // *ptr can be used to get or set the value of x

}
