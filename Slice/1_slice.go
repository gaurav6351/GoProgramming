package main

import "fmt"

func main() {
    elements := []string{"cat", "dog", "bird"}
    elements = append(elements, "fish", "snake")
    fmt.Println(elements)
}
