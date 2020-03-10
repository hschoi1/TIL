package main

import "fmt"

func main() {
   i := 1
   for i <=3 {
     fmt.Println(i)
     i = i + 1
   }  //1, 2, 3

   for j := 7; j<=9; j++{
     fmt.Println(j)
   }  //7,8,9

   for { //loop repeatedly until you break out of the loop or retrun
     fmt.Println("loop")
     break
   }  //loop

   for n :=0; n <= 5; n++ {
     if n%2 == 0 {
        continue
     } // to the next iteration of the loop
     fmt.Println(n)
   } //1,3,5
}
