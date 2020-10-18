package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "set" {
		e := NewEnigma("")
		settings := e.GetSettings()
		ioutil.WriteFile(Settings, []byte(settings), 0666)
		fmt.Println("New settings were successfully set!")
		return
	}
	if len(os.Args) < 3 {
		fmt.Printf("Usage:\n")
		fmt.Printf("\tcode:\n\t\t%v input output\n", os.Args[0])
		fmt.Printf("\tset new settings:\n\t\t%v set\n", os.Args[0])
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

	// e2 := NewEnigma("")

	// for i := 0; i < RotorCount; i++ {
	// 	cnt := 0
	// 	for k, v := range e2.rotors[i].ring {
	// 		if k == v {
	// 			cnt++
	// 		}
	// 	}
	// 	fmt.Println(Length-cnt, cnt)
	// }

	fmt.Println("Successfully done!")
}
