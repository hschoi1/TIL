package main

import(
  "fmt"
  "math"
)

const s string ="constant"
// A const statement can appear anywhere a var statement can
func main() {

   var a = "initial"
   fmt.Println(a) //initial

   var b,c int = 1, 2
   fmt.Println(b,c) // 1,2

   var d = true
   fmt.Println(d) //true

   var e int
   fmt.Println(e) //0

   f := "apple" // := syntax is shorthand for declaring and initializing a variable
   fmt.Println(f) // apple

   const n = 500

   const g = 3e20 / n
   fmt.Println(g)  //6e+17

   fmt.Println(int64(g)) //600000000000000000
   fmt.Println(math.Sin(n)) // -0.46777180532247614

}
