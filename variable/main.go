package main

// import fmt
import (
	"fmt"
)

// bikin function main disini
func main() {

	// bikin variabel name dan age disini
	var name string = "Aris Wijaya"
	age := 25

	// kita print
	fmt.Println(name)
	fmt.Println(age)
	Latihan()
}

// latihan coding
func Latihan() {

	var name string = "Aris"
	var age int = 25
	var city string = "surabaya"
	var balance int = 1500000
	var active bool = true

	fmt.Println("Nama : ", name)
	fmt.Println("Umur :", age)
	fmt.Println("City :", city)
	fmt.Println("Saldo :", balance)
	fmt.Println("Status :", active)

}
