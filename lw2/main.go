package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	Settings = "settings.txt"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v input output\n", os.Args[0])
		return
	}

	in := os.Args[1]
	out := os.Args[2]

	settings, err := ioutil.ReadFile(Settings)
	if err != nil {
		log.Fatal(err)
	}

	e := NewEnigma(string(settings))
	data, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	encText := e.Code(data)
	ioutil.WriteFile(out, encText, 0666)

	fmt.Println("Successfully done!")
}
