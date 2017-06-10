package main

import "fmt"

func main() {
    // Map animal names to color strings.
    // ... Create a map with composite literal syntax.
    colors := map[string]string{
        "bird":  "blue",
        "snake": "green",
        "cat":   "black",
    }

    // Get color of snake.
    c := colors["snake"]

    // Display string.
    fmt.Println(c)
}
