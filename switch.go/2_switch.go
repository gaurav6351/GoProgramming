package main

import "fmt"

func result(v int) int {
    // Return a value based on a switch.
    switch v {
    case 10, 20, 30:
        return v + 5
    case 15, 25, 35:
        return v - 5
    }
    return v
}

func main() {
    // Call the method that uses a switch.
    number := result(10)
    fmt.Println(number)

    number = result(25)
    fmt.Println(number)
}
