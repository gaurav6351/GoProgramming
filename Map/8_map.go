package main

import (
    "fmt"
    "time"
)

func main() {
    lookup := map[string]int{
        "cat":  0,
        "dog":  1,
        "fish": 2,
        "bird": 3}
    values := []string{"cat", "dog", "fish", "bird"}
    temp := 0

    t0 := time.Now()

    // Version 1: search map with lookup.
    for i := 0; i < 10000000; i++ {
        v, ok := lookup["bird"]
        if ok {
            temp = v
        }
    }

    t1 := time.Now()

    // Version 2: search slice with for-loop.
    for i := 0; i < 10000000; i++ {
        for x := range(values) {
            if values[x] == "bird" {
                temp = x
                break
            }
        }
    }

    t2 := time.Now()
    // Benchmark results.
    fmt.Println(temp)
    fmt.Println(t1.Sub(t0))
    fmt.Println(t2.Sub(t1))
}
