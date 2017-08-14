package main

import "fmt"

func main() {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83
    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))

    // AppendSlice()
    // CreateArray()
    // CopySlice()
    Maps()
}

// Function to create an array with fixed length
func CreateArray() {
    x := [5]float64{
        98,
        93,
        48,
        57,
        21,
    }

    fmt.Println(x)
}

// Function to create a slice with a max length of 5
func CreateSlice() {
    x := make([]float64, 5, 10)

    fmt.Println(x)
}

// Function to append a slice with an existing array
func AppendSlice() {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)

    fmt.Println(slice1, slice2)
}

// Function to copy an array, limited to the declared length of the container (in this case 2)
func CopySlice() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)

    fmt.Println(slice1, slice2)
}

// Function to map a key-value pair
func Maps() {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements["Kim"])
}
