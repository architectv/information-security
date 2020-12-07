package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"lw6/lzw"
	"os"
	"strings"
)

const (
	Path           = "data/"
	CompressFlag   = "c"
	DecompressFlag = "d"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v file flag\n", os.Args[0])
		fmt.Printf("flag:\n\ts - compress file\n\tv - decompress file\n")
		return
	}

	fileName := os.Args[1]
	flag := os.Args[2]

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	path := strings.Split(fileName, "/")
	fileName = path[len(path)-1]

	switch flag {
	case CompressFlag:
		compressedData := lzw.Compress(data)
		compressedFile := Path + "comp/" + fileName
		ioutil.WriteFile(compressedFile, compressedData, 0666)
		fmt.Println("Successfully compressed!")
	case DecompressFlag:
		decompressedData := lzw.Decompress(data)
		decompressedFile := Path + "decomp/" + fileName
		ioutil.WriteFile(decompressedFile, decompressedData, 0666)
		fmt.Println("Successfully decompressed!")
	default:
		fmt.Println("Wrong flag!")
	}
}
