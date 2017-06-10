package main

import "fmt"

func main() {
    // Create map with three string keys.
    sizes := map[string]int{
        "XL": 20,
        "L":  10,
        "M":  5,
    }

    // Loop over map and append keys to empty slice.
    keys := []string{}
    for key, _ := range sizes {
        keys = append(keys, key)
    }

    // This is a slice of the keys.
    fmt.Println(keys)
}
