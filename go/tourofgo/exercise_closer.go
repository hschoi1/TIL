//https://tour.golang.org/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  x, y := 0, 1
  i := 0
  return func() int{
     if i!=0{
       x, y = y, x+y
      }
      i++
      return x
   }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
/*
0
1
1
2
3
5
8
13
21
34
*/
