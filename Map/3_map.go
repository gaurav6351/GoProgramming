package main

import "fmt"

func main() {
    // Create an empty map and add three pairs to it.
    ids := map[string]int{}
    ids["steve"] = 10
    ids["mark"] = 20
    ids["adnan"] = 30
    fmt.Println(len(ids))

    // Delete one key from it.
    delete(ids, "steve")
    fmt.Println(len(ids))
}
