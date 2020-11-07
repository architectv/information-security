package main

import (
	"fmt"
	"io/ioutil"
	"log"
	des "lw3/des"
	"os"
	"strings"
)

const (
	Key  = "config/key.txt"
	Path = "data/"
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

	path := strings.Split(inputFile, "/")
	fileName := path[len(path)-1]
	encData := desEcb.Encode(data)
	encFile := Path + "enc/" + fileName
	ioutil.WriteFile(encFile, encData, 0666)

	decData := desEcb.Decode(encData)
	decFile := Path + "dec/" + fileName
	ioutil.WriteFile(decFile, decData, 0666)

	fmt.Println("Successfully done!")
}
