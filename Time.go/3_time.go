package main

import (
    "fmt"
    "time"
)

func main() {
    // T1 is a time earlier than t2.
    t1 := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
    t2 := time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC)

    // This is true: the first time is before the second.
    if t1.Before(t2) {
        fmt.Println(true)
    }

    // This is true.
    if t2.After(t1) {
        fmt.Println(true)
    }

    // This returns false.
    fmt.Println(t1.After(t2))
}
