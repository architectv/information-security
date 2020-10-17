package main

import (
	"strconv"
	"strings"
)

type Enigma struct {
	rotors    [RotorCount]Rotor
	reflector Reflector
}

func NewEnigma(settings string) *Enigma {
	e := new(Enigma)

	var rotor_settings [RotorCount]string
	var reflector_settings string
	if settings != "" {
		s := strings.Split(settings, "\n")
		for i := 0; i < RotorCount; i++ {
			rotor_settings[i] = s[i]
		}
		reflector_settings = s[RotorCount]
	}

	for i := 0; i < RotorCount; i++ {
		e.rotors[i] = makeRotor(rotor_settings[i])
	}

	e.reflector = makeReflector(reflector_settings)

	return e
}

func (e *Enigma) Code(text []byte) []byte {
	var encText []byte
	for _, symbol := range text {
		encSymbol := e.getSymbol(symbol)
		encText = append(encText, encSymbol)
		e.rotate()
	}

	return encText
}

func (e *Enigma) getSymbol(in byte) byte {
	out := in
	for i := 0; i < RotorCount; i++ {
		out = byte(e.rotors[i].getVal(int(out)))
	}

	out = byte(e.reflector.getVal(int(out)))

	for i := RotorCount - 1; i >= 0; i-- {
		out = byte(e.rotors[i].getKey(int(out)))
	}

	return out
}

func (e *Enigma) rotate() {
	e.rotors[0].inc()
	for i := 0; i < RotorCount-1; i++ {
		if e.rotors[i].isFullTurn() {
			e.rotors[i+1].inc()
		}
	}

	for i := 0; i < RotorCount; i++ {
		e.rotors[i].upd()
	}
}

func (e *Enigma) GetSettings() string {
	settings := ""
	for i := 0; i < RotorCount; i++ {
		ring := e.rotors[i].getRing()
		for _, v := range ring {
			settings += strconv.Itoa(v) + " "
		}
		settings = settings[:len(settings)-1] + "\n"
	}

	mapping := e.reflector.getMapping()
	for i := 0; i < Length; i++ {
		settings += strconv.Itoa(i) + ":" + strconv.Itoa(mapping[i]) + " "
	}
	settings = settings[:len(settings)-1] + "\n"

	return settings
}
