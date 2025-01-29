package main

import "fmt"

func main() {

	// sama seperti variable, array juga punya nilai default
	var language [5]string

	language[0] = "php"
	language[1] = "jsht"
	language[2] = "java"
	language[3] = "c"
	language[4] = "ruby"

	fmt.Println(language)
}
