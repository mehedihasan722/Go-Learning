package main

import "fmt"

func main()  {
	var decimalNumber int

	x := 2
	fmt.Scan(&x)
	fmt.Printf("The number x in binary %o\n",x)

	y := 3

	fmt.Printf("The float number in binary %b\n",y)

	fmt.Scanf("%d",&decimalNumber)

	name := "Mehedi Hasan"
	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)

	fmt.Printf("Your choosen number is = %b\n",decimalNumber)
	fmt.Printf("Your choosen number is = %o\n",decimalNumber)
	fmt.Printf("Your choosen number is = %x",decimalNumber)

}