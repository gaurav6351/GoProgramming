package main

import "fmt"

func main() {
    slice1 := []int{10, 20, 30}
    slice2 := []int{0, 0, 0, 1000}
    // Display slice1 and slice2.
    fmt.Println(slice1)
    fmt.Println(slice2)

    // Copy elements from slice1 into slice2.
    copy(slice2, slice1)

    // Slice2 now has values from slice1.
    fmt.Println(slice2)
}
