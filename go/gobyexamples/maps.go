package main

import "fmt"

func main() {

  m := make(map[string]int)

  m["k_1"] = 5
  m["k_2"] = 10

  fmt.Println("map:", m) // map[k_1:5 k_2:10]

  v1 := m["k_1"]
  fmt.Println("v_1: ", v1) // v_1: 5

  fmt.Println("len: ", len(m)) // len: 2

  delete(m, "k_2")
  fmt.Println("map:", m) // map[k_2:10]

  _, prs := m["k_2"]
  fmt.Println("prs:", prs) //false

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n) //map[bar:2 foo:1]
}
