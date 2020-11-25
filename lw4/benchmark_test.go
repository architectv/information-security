package main

import (
	"io/ioutil"
	"log"
	v1 "lw4/rsa/v1"
	v2 "lw4/rsa/v2"
	"testing"
)

const (
	Bits = 50
	File = "data/in/cat.jpg"
)

func BenchmarkRSAv1(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		inputFile := File
		rsa := v1.NewRSA(Bits)
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		encData := rsa.Encode(data)
		rsa.Decode(encData)
	}
}

func BenchmarkRSAv2(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		inputFile := File
		rsa := v2.NewRSA(Bits)
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		encData := rsa.Encode(data)
		rsa.Decode(encData)
	}
}
