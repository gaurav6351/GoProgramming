package main

import "fmt"

func main()  {
  sum := 0
	for i := 0; i < 10; i++ {
		sum += i
    fmt.Println(i)
	}
	fmt.Println(sum)


  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
    }

}
