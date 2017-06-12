package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a Date in 2006.
    t := time.Date(2006, 1, 1, 12, 0, 0, 0, time.UTC)
    fmt.Println(t)
}
