package main

import "fmt"

func main() {
    var x string
    var y string
    x = "first"
    fmt.Println(x)
    x = x + "second"
    fmt.Println(x)
    y = "third"
    fmt.Println(x == y)
}
