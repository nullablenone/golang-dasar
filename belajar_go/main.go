package main

import "fmt"

func main() {
	teks := "ingin menjadi programmer handal namun enggan mengcoding"

	for index, iterasi := range teks {
		fmt.Println("no :", index, "car :", string(iterasi))
	}

	// _: Kalau nggak perlu pakai index/key, biasanya ditaruh _ untuk mengabaikan.
	for _, iterasi := range teks {
		fmt.Println("car :", string(iterasi))
	}



}
