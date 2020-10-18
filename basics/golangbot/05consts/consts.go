package main

import "fmt"

func main() {
	const a = 55 //allowed
	//a = 89       //reassignment not allowed
	//const b = math.Sqrt(4)//not allowed

	const hello = "Hello World" // string constant does not have any type.

	var name = "Sam"
	fmt.Printf("type %T value %v", name, name)
	/* how does the following program which assigns a variable name to an UNTYPED constant Sam work? The answer is untyped constants have a default type associated with them and they supply it if and only if a line of code demands it.*/
	const typedhello string = "Hello World" //typed constant

	var defaultName = "Sam"         //allowed,default type of the Sam is string
	type myString string            //new type myString is alias of string
	var customName myString = "Sam" //allowed,Sam is untyped
	//customName = defaultName //not allowed
	/* 	the assignment customName = defaultName is not allowed: cannot use defaultName (type string) as type myString in assignment */
}
