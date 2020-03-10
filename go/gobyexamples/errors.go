package main

import (
  "errors"
  "fmt"
)
// errors have type error, a built-in interface
func f1(arg int) (int, error) {
  if arg == 1 {
      return -1, errors.New("no 1!")
  }
  return arg + 3, nil
} // a nil value here indicates that there was no error.

type argError struct {
   arg int
   prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 1 {
        return -1, &argError{arg, "no no 1!"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{1, 7} {
        if r, e := f1(i); e != nil {
          fmt.Println("f1 failed:", e)  //f1 failed: no 1!
        } else {
            fmt.Println("f2 worked:", r) // f2 worked: 10
        }
    }

    _, e := f2(1)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg) //1
        fmt.Println(ae.prob) //"no no 1!"
    }
}
