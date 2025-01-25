package main

import (
	"fmt"
	"learn-go/calculation"
)

func main()  {
  fmt.Println("hallo word")

  result := calculation.Add(10, 10)

  fmt.Println(result)
}