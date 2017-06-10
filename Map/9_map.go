package main

import "fmt"

func main() {
    // Create a map with a capacity of 200 pairs.
    // ... This makes adding the first 200 pairs faster.
    lookup := make(map[string]int, 200)

    // Use the new map.
    lookup["cat"] = 10
    result := lookup["cat"]
    fmt.Println(result)
}
