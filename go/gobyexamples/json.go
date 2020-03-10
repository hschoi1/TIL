package main

import (
    "encoding/json"
    "fmt"
    "reflect"
    "os"
)

type response1 struct {
   Page int
   Fruits []string
   Dummy int
}

type response2 struct { // to decode JSON into custom data types
  Page int `json:"page"`
  Fruits []string `json:"fruits"`
} //fields must start with capital letters

func main() {
   bolB, _ := json.Marshal(true)
   fmt.Println(json.Marshal(true)) // [116 114 117 101] <nil>
   fmt.Println(string(bolB)) // true

   intB, _ := json.Marshal(1)
   fmt.Println(string(intB)) //1

   fltB, _ := json.Marshal(2.34)
   fmt.Println(string(fltB)) //2.34

   strB, _ := json.Marshal("gopher")
   fmt.Println(string(strB)) // "gopher"

   slcD := []string{"apple", "peach", "pear"}
   slcB, _ := json.Marshal(slcD) // slices encode to JSON arrays
   fmt.Println(string(slcB)) // ["apple", "peach", "pear"]

   mapD := map[string]int{"apple":5, "lettuces":7} // maps encode to objects
   mapB, _ := json.Marshal(mapD)
   fmt.Println(reflect.TypeOf(mapD)) //map[string]int
   fmt.Println(reflect.TypeOf(mapB)) //[]uinit8
   fmt.Println(string(mapB)) //{"apple"5, "lettuces":7}

   res1D := &response1{
       Page: 1,
       Fruits: []string{"apple", "peach", "pear"}}
   res1B, _ := json.Marshal(res1D) // will only include exported fields in the encoded output and use those names as JSON keys
   fmt.Println(string(res1B)) //{"Page":1, "Fruits":["apple", "peach", "pear"]}

   byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

   var dat map[string]interface{}

   if err := json.Unmarshal(byt, &dat); err != nil {
      panic(err)
   }
   fmt.Println(dat) // map[num:6.13 strs:[a b]]

   fmt.Println(reflect.TypeOf(dat["num"])) //float64
   num := dat["num"].(float64) // why is this necessary???
   fmt.Println(reflect.TypeOf(num)) //float64
   fmt.Println(num) //6.13

   strs := dat["strs"].([]interface{}) // acessing nested data
   str1 := strs[0].(string)
   fmt.Println(str1) //a

   str := `{"page": 1, "fruits": ["apple", "peach"]}`
   res := response2{}
   json.Unmarshal([]byte(str), &res)
   fmt.Println(res)  // {1 [apple peach]}
   fmt.Println(res.Fruits[0]) // apple

   enc := json.NewEncoder(os.Stdout) //can stream JSON encdoings directly to os.Writers or even HTTP response bodies
   d := map[string]int{"apple":5, "lettuce":7}
   enc.Encode(d) //{"apple":5,"lettuce":7}
}
