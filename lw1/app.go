package main

import "fmt"

func main() {
	check := IsLicensed()
	if check {
		fmt.Println("It's OK! You got a license for this program.")
	} else {
		fmt.Println("ERROR! There is no license!")
	}
}
