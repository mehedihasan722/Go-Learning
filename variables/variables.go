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

	var a = "initial"
	fmt.Println(a)

	var b,c int = 1, 2
	fmt.Println(b, c)

	var d  = true 
	fmt.Println(d)

	var e int 
	fmt.Println(e)

	f := "Appple"
	fmt.Println(f)
}

/*
 ---Output ---

Mehedi Hasan
Mehedi Hasan
Mehedi Hasan

*/