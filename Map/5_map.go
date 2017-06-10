package main

import "fmt"

func main() {
    meds := map[string]int{
        "Ativan": 15,
        "Xanax": 20,
        "Klonopin": 30,
    }

    // The ok variable is set to true.
    if dosage, ok := meds["Xanax"]; ok {
        fmt.Println("Xanax", dosage)
    }

    // The ok variable is set to false.
    // ... The string does not exist in the map.
    if dosage, ok := meds["Atenolol"]; ok {
        fmt.Println("Atenolol", dosage)
    }
}
