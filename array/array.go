package main

import "fmt"

func main() {

	var vowel = [5] string {"a","b","c","d","e"}

	for i:=0 ; i< len(vowel); i++ {
		fmt.Println(i+1,"vowel is =",vowel[i])
	}

	fmt.Println(vowel)
}