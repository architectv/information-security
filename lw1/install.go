package main

import "fmt"

func main() {
	check := WriteKey()
	if check {
		fmt.Println("Key was written successfully!")
	} else {
		fmt.Println("Error while writing key!")
	}
}
