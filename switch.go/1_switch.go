package main

import "fmt"

func main() {
    id := 10
   // Use switch with multiple values in each case.
    switch id {
    case 10, 12, 14:
        fmt.Println("Even")
    case 11, 13, 15:
        fmt.Println("Odd")
    }
}
