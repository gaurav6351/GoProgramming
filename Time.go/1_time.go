package main

import (
    "fmt"
    "time"
)

func main() {
    // Get the current time.
    t := time.Now()
    fmt.Println(t)

    // Print year, month and day of the current time.
    fmt.Println(t.Year())
    fmt.Println(t.Month())
    fmt.Println(t.Day())
}
