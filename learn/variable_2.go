package main

import "fmt"

var var1 string = "bhoo_boy"

func main() {

	var var2 float64 = 1.8963
	var var3 int
	var3 = 2222
	var4 := "3333" // like python

	fmt.Printf("%v <-> %T \n", var1, var1)
	fmt.Printf("%v <-> %T \n", var2, var2)
	fmt.Printf("%v <-> %T \n", var3, var3)
	fmt.Printf("%v <-> %T \n", var4, var4)
}
