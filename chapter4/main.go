package main

import "fmt"

func main() {
    fmt.Print("Enter a Fahrenheit temp: ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)
    
    centigrade := (fahrenheit - 32) * 5 / 9
    fmt.Println(centigrade)
}
