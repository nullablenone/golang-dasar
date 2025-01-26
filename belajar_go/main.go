package main

import "fmt"

func main() {
	// else if
	score := 90
	var grade string

	if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else if score <= 80 {
		grade = "B"
	}else if score <= 100 {
		grade = "A"
	}else{
		grade = "NULL"
	}

	fmt.Println(grade)
}
