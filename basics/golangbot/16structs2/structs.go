package main

import (
	"fmt"

	"github.com/golangbot/structs/computer"
)

//Person struct with anonymous field
type Person struct {
	string
	int
}

//Employee named structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

//Address structure
type Address struct {
	city, state string
}

//Person1 nested struct
type Person1 struct {
	name    string
	age     int
	address Address
}

//Person2 struct with promoted fields
type Person2 struct {
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

func main() {
	//creating named structure using field name
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating named structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1\n", emp1)
	fmt.Println("\nEmployee 2\n", emp2)

	//creating anonymous structure
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("\nEmployee 3\n", emp3)

	//zero valued structure
	var emp4 Employee
	fmt.Println("\nEmployee 4\n", emp4)
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("\nEmployee 5\n", emp5)

	//accessing individual fields of a structure
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("\nEmployee 6\nFirst Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("\nEmployee 7:", emp7)

	//Pointer to a struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("\nEmployee 8\nFirst Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	emp9 := &emp7
	fmt.Println("\nEmployee 9\nFirst Name:", emp9.firstName)
	fmt.Println("Age:", emp9.age)

	//creating structure with anonymous field
	p := Person{"Naveen", 50}
	fmt.Println("\np", p)

	//accessing anonymous fields using their type as name
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println("\np1", p1)

	//nested struct
	var p2 Person1
	p2.name = "Naveen"
	p2.age = 50
	p2.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("\np2")
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.address.city)
	fmt.Println("State:", p2.address.state)

	//promoted fields
	var p3 Person2
	p3.name = "Naveen"
	p3.age = 50
	p3.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("\nP3")
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)   //city is promoted field
	fmt.Println("State:", p3.state) //state is promoted field

	//accessing exported fields from another package
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "Mac Mini"//trying to access unexported field will result in compilation error
	fmt.Println("Spec:", spec)

	//structs equality
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("\nname1 and name2 are equal")
	} else {
		fmt.Println("\nname1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("\nname3 and name4 are equal")
	} else {
		fmt.Println("\nname3 and name4 are not equal")
	}

}
