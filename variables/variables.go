package main

import "fmt"

func main() {
	firstName := "Mehedi"
	lastName := "Hasan"

	const mod uint32 = 1e9+7

	var age int = 32

	fmt.Print(firstName," ",lastName," -> ", age ,"\n")

	firstName = "Jahid"

	fmt.Printf("%s %s\n",firstName,lastName)

	lastName = "Hossain"

	fmt.Println(firstName,lastName,mod)
}

/*
 ---Output ---

Mehedi Hasan
Mehedi Hasan
Mehedi Hasan

*/