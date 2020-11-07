package main

import (
	"fmt"
	"io/ioutil"
	"log"
	des "lw3/des"
	"os"
)

const (
	Key = "config/key.txt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v input_file\n", os.Args[0])
		return
	}

	inputFile := os.Args[1]

	key, err := ioutil.ReadFile(Key)
	if err != nil {
		log.Fatal(err)
	}

	desEcb := des.NewDES(string(key))
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	encData := desEcb.Encode(data)
	encFile := inputFile + ".enc"
	ioutil.WriteFile(encFile, encData, 0666)

	decData := desEcb.Decode(encData)
	decFile := inputFile + ".dec"
	ioutil.WriteFile(decFile, decData, 0666)

	fmt.Println("Successfully done!")
}
