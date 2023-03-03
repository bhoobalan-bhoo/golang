package main

import (
	"fmt"
	"strconv"
)

// variable conversion

func main() {

	var var2 float64 = 1.8963
	var var3 int = 8
	fmt.Printf("%v <-> %T \n", var2, var2)
	fmt.Printf("%v <-> %T \n", var3, var3)
	var2 = float64(var3)

	fmt.Printf("%v <-> %T \n", var2, var2)
	fmt.Printf("%v <-> %T \n", var3, var3)

	var var4 string = string(var3) // returns with ASCII value

	fmt.Printf("%v <-> %T \n", var3, var3)
	fmt.Printf("%v <-> %T \n", var4, var4)

	jun := 956
	var var5 string = strconv.Itoa(jun) // To get actual Number to string
	fmt.Printf("%v <-> %T \n", jun, jun)
	fmt.Printf("%v <-> %T \n", var5, var5)

}
