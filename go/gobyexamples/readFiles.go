package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := ioutil.ReadFile("defer.txt") //slurp a file's entire contents into memory
    check(err)
    fmt.Print(string(dat)) //blah blah

    f, err := os.Open("defer.txt") //obtain an os.File value
    check(err)

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) //5 bytes: blah

    o2, err := f.Seek(6, 0) //to a known location and Read from there
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2) //2 bytes @ 6
    fmt.Printf("%v\n", string(b2[:n2])) //: la

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)

    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))//2 bytes @ 6: la

    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    // bufio package implements a bueffered reader, useful both for efficienty with many small reads
    // and for additional reading methods it provides
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))  //5 bytes: blah

    f.Close() // usually this would be scheduled immediately after Opening with defer
}
