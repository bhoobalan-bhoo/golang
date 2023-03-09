package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)
	
    s[0] = "bhoo"
    s[1] = "grey"
    s[2] = "jol"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "dikj")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("Checking len twoD[0]",  len(twoD[0]))
	fmt.Println("Checking len twoD[1]",  len(twoD[1]))
	fmt.Println("Checking len twoD[2]",  len(twoD[2]))
    fmt.Println("2d: ", twoD)
}