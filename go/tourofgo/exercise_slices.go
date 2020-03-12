//https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
   incNums := make([][]uint8, dy)

   for i:=0; i<dy; i++{
  	  incNums[i] = make([]uint8, dx)
	   for j:=0; j<dx; j++ {
	      incNums[i][j] = uint8(i+j)
	   }
   }
   return incNums
}

func main() {
	pic.Show(Pic)
}
