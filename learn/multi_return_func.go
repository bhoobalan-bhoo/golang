package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    // fmt.Println(_)
	// "_" its used to ignore values. if try to print phrases error
    fmt.Println(c)	
}