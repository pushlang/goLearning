package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	//int array length 3, assigned the zero value of the array type, array index starts at 0
	var a [3]int
	a[0], a[1], a[2] = 12, 78, 50
	//a[3] = 1 //invalid array index - out of bounds
	

	//short hand declaration to create array, not necessary that all elements in an array have to be assigned a value, remaining elements are assigned 0 automatically
	b := [3]int{12, 78, 50}

	// ... compiler determines the length
	c := [...]int{12, 78, 50}

	//not possible since [3]int and [5]int are distinct types
	d := [3]int{5, 78, 8}
	var e [5]int
	//d = e

	// a copy of a is assigned to b
	f := [...]string{"USA", "China", "India", "Germany", "France"}
	g := f

	//num is passed by value
	num := [...]int{5, 6, 7, 8, 8}
	changeLocal(num)

	//length of an array
	fmt.Println("length of a is", len(num))

	//range returns both the index and value
	a2 := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a2 {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}

	for _, v := range a2 { //ignores index
		fmt.Printf("the element of a is %.2f\n", v)
	}

	fmt.Println(a, b, c, d, e, g, sum)

	//multidimensional array
	a3 := [3][2]string{
		{"lion", "tiger"}, {"cat", "dog"}, {"pigeon", "peacock"},
	} //this comma is necessary
	printarray(a3)
	var b3 [3][2]string
	b3[0][0], b3[0][1] = "apple", "samsung"
	b3[1][0], b3[1][1] = "microsoft", "google"
	b3[2][0], b3[2][1] = "AT&T", "T-Mobile"

	fmt.Printf("\n")
	printarray(b3)

	//SLICES
	//A slice does not own any data of its own, just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.
	//length - the number of elements in the slice
	//capacity - the number of elements in the underlying array starting from the index from which the slice is created
	a4 := [5]int{76, 77, 78, 79, 80}
	var b4 []int = a4[1:4] //creates a slice from a[1] to a[3]
	c4 := []int{6, 7, 8}   //creates and array and returns a slice reference
	fmt.Println(b4, c4)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                               //re-slicing till its capacity
	fmt.Println("length", len(fruitslice), "capacity", cap(fruitslice))

	//creating a slice using make
	i1 := make([]int, 5, 5) //output [0 0 0 0 0]
	fmt.Println(i1)

	//appending to a slice
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	//appending to nil slice
	var names []string
	if names == nil {
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	//appending fruits to veggies.
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	//When a slice is passed to a function, pointer variable will refer to the same underlying array. Hence changes made inside the function are visible outside the function too.
	nos := []int{8, 7, 6}
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	//multidimensional slices
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	country := make([][]string, 2)
	var carss1,carss2 []string

	carss1 = append(carss1, "audi", "bmw", "mercedes")
	country[0] = append(country[0], carss1...)

	carss2 = append(carss2, "renault", "pegault", "citroen")
	country[1] = append(country[1], carss2...)
	
	fmt.Println(country)

	//memory optimisation, Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice.
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

}
