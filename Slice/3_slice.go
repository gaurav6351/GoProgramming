package main

import "fmt"

func main() {
    // Create an empty slice.
    // ... Its length is 0.
    items := []string{}
    fmt.Println(len(items))

    // Append a string and the slice now has a length of 1.
    items = append(items, "cat")
    fmt.Println(len(items))
}
