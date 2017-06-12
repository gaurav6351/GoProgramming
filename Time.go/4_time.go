package main

import (
    "fmt"
    "time"
)

func main() {
    // This date is used to indicate the layout.
    const layout = "Jan 2, 2006"
    // Format Now with the layout const.
    t := time.Now()
    res := t.Format(layout)
    fmt.Println(res)
}
