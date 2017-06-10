package main

import "fmt"

func main() {
    elements := []int{100, 200, 300}

    // Capacity is now 3.
    fmt.Println(cap(elements))

    // Append another element to the slice.
    elements = append(elements, 400)

    // Capacity is now 6—it has doubled.
    fmt.Println(cap(elements))
}
