package main

import "fmt"

func main() {
	//  For klasik
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// For dengan kondisi saja 
	n := 1
	for n <= 5 {
		fmt.Println(n)
		n++
	}

	// For range 
	names := []string{"Aris", "Wijaya", "Nur"}
	for index, name := range names {
		fmt.Println(index, name)
	}
}