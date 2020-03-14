//https://tour.golang.org/concurrency/8
package main

import ("golang.org/x/tour/tree"
        "fmt")

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
  if t == nil{
    return
  }
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{

  ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:=0; i<cap(ch1); i++ {
	   a,b := <-ch1, <-ch2
	   if a!=b{
	      return false
	   }
	}
	return true
}
func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
}
