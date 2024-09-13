package main

import "fmt" 

type Student struct {
	id int
	address string
	age int
}

func displayInfo(s Student) {
	fmt.Println(s.id)
	fmt.Println(s.address)
	fmt.Println(s.age)
}

func main() {

	
	mehedi := Student{101, "lal khan bazar", 30}
	displayInfo(mehedi)
}