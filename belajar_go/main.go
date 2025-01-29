package main

import (
	"fmt"
)

func main() {
	//mencetak cika index bilangan genap
	//mencetak jika iterasi nya huruf vokal

	teks := "Golang the bast languange"

	for index, iterasi := range teks {

		if index%2 == 0 {
			
			switch iterasi {
			case 'a', 'i', 'u', 'e', 'o':
				fmt.Println("index :", index, "car :", string(iterasi))

				// case 'a':
				// 	fmt.Println("index :", index, "car :", string(iterasi))
				// case 'i':
				// 	fmt.Println("index :", index, "car :", string(iterasi))
				// case 'u':
				// 	fmt.Println("index :", index, "car :", string(iterasi))
				// case 'e':
				// 	fmt.Println("index :", index, "car :", string(iterasi))
				// case 'o':
				// 	fmt.Println("index :", index, "car :", string(iterasi))
			}

		}
	}
}
