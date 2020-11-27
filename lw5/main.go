package main

import (
	"fmt"
	"lw5/dsa"
	"os"
)

const (
	SignFlag   = "s"
	VerifyFlag = "v"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v file flag\n", os.Args[0])
		fmt.Printf("flag:\n\ts - sign file\n\tv - verify sign\n")
		return
	}

	fileName := os.Args[1]
	flag := os.Args[2]
	var err error
	switch flag {
	case SignFlag:
		err = dsa.SignFile(fileName)
		if err != nil {
			fmt.Println("ERROR =>", err.Error())
			return
		}
		fmt.Println("Successfully signed!")
	case VerifyFlag:
		err = dsa.VerifyFile(fileName)
		if err != nil {
			fmt.Println("ERROR =>", err.Error())
			return
		}
		fmt.Println("Successfully verified!")
	default:
		fmt.Println("Wrong flag!")
	}
}
