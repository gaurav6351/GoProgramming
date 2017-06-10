package main

import "fmt"

func main() {
    // Create a slice of 5 integers.
    values := make([]int, 5)

    // Assign some elements.
    values[0] = 100
    values[4] = 200

    // Loop over elements in slice and display them.
    for v := range values {
        fmt.Println(values[v])
    }
}
