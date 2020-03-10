package main

import "fmt"

func zeroval(init_val int) {
  init_val = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
} // dereferences the pointer from its memory address to the current value at that address

func main() {
  i := 1
  fmt.Println("initial:", i) // initial: 1

  zeroval(i)
  fmt.Println("zeroval:", i) //zeroval: 1

  zeroptr(&i) // & gives memory address
  fmt.Println("zeroptr:", i) //zeroval: 0

  fmt.Println("pointer:", &i)  // pointer: pointer: 0xc0000160a8
  //has a reference to the memory address for i
}
