package main

import "fmt"

func main() {
    s:= make([]string, 3) // Unlike arrays, slices are typed only by the elemnts they contain
    fmt.Println("emp:", s) //emp: [  ]

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:" ,s) //set: [a b c]
    fmt.Println("get:", s[2]) //get: c

    fmt.Println("len:", len(s)) //len: 3

    s = append(s, "d") //returns a slice containing d
    s = append(s, "e", "f")
    fmt.Println("apd:", s) //apd: [a b c d e f]

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c) //cpy: [a b c d e f]

    l := s[2:5]
    fmt.Println("sl1:", l) //sl1: [c d e]

    t := []string{"g", "h", "i"} // declare and intialize
    fmt.Println("dcl:", t) //dcl: [g h i]

    twoD := make([][]int, 3)
    for i := 0; i<3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) //2d:  [[0] [1 2] [2 3 4]]
}
