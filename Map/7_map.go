package main

import "fmt"

func main() {
    // A simple map.
    birds := map[string]string{
        "finch":    "yellow",
        "parakeet": "blue",
    }

    // Place values in a string slice.
    values := []string{}
    for _, value := range birds {
        values = append(values, value)
    }

    // The values.
    fmt.Println(values)
}
