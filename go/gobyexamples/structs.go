package main

import "fmt"

type person struct {
  name string
  age int
}

func NewPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  return &p
}

func main() {
  fmt.Println(person{"Bob", 20}) // {Bob 20}
  fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}
  fmt.Println(person{name: "Fred"}) // {Fred 0}

  fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}
  // & prefix yields a pointer to the struct
  fmt.Println(NewPerson("Jon")) // &{Jon 42}

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name) //Sean
  s.name = "Shawn"
  fmt.Println(s.name, s.age) //Shawn, 50

  sp := &s
  fmt.Println(sp) // &{Shawn 50}
  fmt.Println(sp.age) //50

  sp.age = 51
  fmt.Println(sp.age) //51
}
