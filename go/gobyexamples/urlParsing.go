package main

import (
   "fmt"
   "net"
   "net/url"
)

func main() {

    s := "postgress://user:pass@host.com:12345/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    fmt.Println(u.Scheme) //postgress

    fmt.Println(u.User) //user:pass
    fmt.Println(u.User.Username()) //user
    p, _ := u.User.Password()
    fmt.Println(p) //pass

    fmt.Println(u.Host) //host.com:12345
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host) //host
    fmt.Println(port) //12345

    fmt.Println(u.Path) // /path
    fmt.Println(u.Fragment) // f (the fragment after #)

    fmt.Println(u.RawQuery) // k=v
    m, _ := url.ParseQuery(u.RawQuery) //The parsed query param maps are from strings to slices of strings
    fmt.Println(m) // map[k:[v]]
    fmt.Println(m["k"][0]) //v
    fmt.Println(m["k1"][0]) // panic: runtime error: index out of range
}
