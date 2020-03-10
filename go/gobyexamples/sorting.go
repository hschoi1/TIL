package main

import (
    "fmt"
    "sort"
)

type byLength []string // to sort by a custom function, we need a corresponding type

//Implement sort.Interface - Len, Less, and Swap
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
} // want to sort in order of increasing string length

func main() {
    strs := []string{"c", "z", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs) //Strings: [a b c z]

    ints := []int{-1,2,4,0}
    sort.Ints(ints)
    fmt.Println("Ints:  ", ints) //Ints:   [-1 0 2 4]

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted  ", s) //Sorted   true

    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)[kiwi peach banana]
}
