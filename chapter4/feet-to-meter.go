package main

import "fmt"

func main() {
    var ft float64
    fmt.Print("Insert measurement in feet: ")
    fmt.Scanf("%f", &ft)

    meters := ft * 0.3048
    fmt.Println(meters)
}
