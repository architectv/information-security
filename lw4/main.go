package main

import (
	"fmt"
	"io/ioutil"
	"log"
	rsa "lw4/rsa/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Path = "data/"
)

func main() {
	logFile, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer logFile.Close()
	// ioutil.Discard
	log.SetOutput(logFile)
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v input_file bits\n", os.Args[0])
		return
	}

	inputFile := os.Args[1]
	bits, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
		return
	}

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	path := strings.Split(inputFile, "/")
	fileName := path[len(path)-1]

	log.Printf("INPUT [%s] %d bytes\n", inputFile, len(data))

	rsa := rsa.NewRSA(bits)

	start := time.Now()
	encData := rsa.Encode(data)
	duration := time.Since(start)
	fmt.Println("Encode time:", duration)
	log.Println("Encode time:", duration)

	encFile := Path + "enc/" + fileName
	ioutil.WriteFile(encFile, encData, 0666)

	start = time.Now()
	decData := rsa.Decode(encData)
	duration = time.Since(start)
	fmt.Println("Decode time:", duration)
	log.Println("Decode time:", duration)

	decFile := Path + "dec/" + fileName
	ioutil.WriteFile(decFile, decData, 0666)

	fmt.Println("Successfully done!")
}
