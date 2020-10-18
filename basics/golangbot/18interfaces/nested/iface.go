package main

import "fmt"

type Salaried interface {
	getSalary() int
	setSalary(b, i, a int)
}

type Salary struct {
	basic     int
	insurance int
	allowance int
}

func (s Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

func (s *Salary) setSalary(b, i, a int) {
	s.basic, s.insurance, s.allowance = b, i, a
}

type Employee struct {
	firstName, lastName string
	//salary              Salaried
	Salaried // with promoted field
}

func main() {
	rossSalary := &Salary{1100, 50, 50}
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		//salary:    rossSalary,
		Salaried:    rossSalary,
	}


	//v := ross.salary.(*Salary)
	//fmt.Println("Ross's salary is", v.basic)

	//fmt.Println("Ross's salary is", ross.salary.getSalary())
	//ross.salary.setSalary(100, 10, 1)
	//fmt.Println("Ross's salary is", ross.salary.getSalary())

	v := ross.Salaried.(*Salary)
	fmt.Println("Ross's salary is", v.basic)

	fmt.Println("Ross's salary is", ross.getSalary())
	ross.setSalary(100,10,1)
	fmt.Println("Ross's salary is", ross.getSalary())
}
