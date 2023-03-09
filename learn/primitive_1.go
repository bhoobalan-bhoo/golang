package main

import "fmt"


func main() {


	// Boolean 
	var var1 bool = true

	m1 := 1 == 1 // declaring logical way
	m2 := 2 == 1
	fmt.Printf("%v <-> %T \n", var1, var1)
	fmt.Printf("%v <-> %T \n", m1, m1)
	fmt.Printf("%v <-> %T \n", m2, m2)


	// Operations

	a:= 5999
	b:= 100
	fmt.Println( a+b )
	fmt.Println( a-b )
	fmt.Println( a*b )
	fmt.Println( a/b )
	

	fmt.Println( a%b )
}
