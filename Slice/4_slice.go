package main

import "fmt"

func main() {
    animals := []string{"bird", "dog", "fish"}

    // Loop over the slice.
    for v := range animals {
        fmt.Println(animals[v])
    }
}
