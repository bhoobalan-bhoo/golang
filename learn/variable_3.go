package main

import "fmt"
// inside var block we can declare multiple varibales with [name, datatype]
var (
	var1 string = "bhoo_boy"
	var2 float64 = 85.5
	var3 string = "bhoo_boy"
	// var4 int = 85.36 // commit a error if float cannot applied to int )
	var4 int = 856
	) 

func main() {

	fmt.Printf("%v <-> %T \n", var1, var1)
	fmt.Printf("%v <-> %T \n", var2, var2)
	fmt.Printf("%v <-> %T \n", var3, var3)
	fmt.Printf("%v <-> %T \n", var4, var4)
}
