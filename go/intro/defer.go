package main

import (
    "fmt"
    "os"
)

func main() {

  f := createFile("deferExample.txt")
  defer closeFile(f) // will be executed at the end of the enclosing function
  writeFile(f)
}

func createFile(p string) *os.File {
   fmt.Println("creating")
   f, err := os.Create(p)
   if err != nil {
       panic(err)
   }
   return f
}

func writeFile(f *os.File) {
   fmt.Println("writing")
   fmt.Fprintln(f, "blah blah")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
       fmt.Fprintf(os.Stderr, "error: %v\n", err)
       os.Exit(1)
    }
}
// creating
// writing
// closing
