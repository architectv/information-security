package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Rotor struct {
	pos  int
	ring [Length]int
}

func makeRotor(settings string) Rotor {
	var r Rotor

	if settings != "" {
		s := strings.Split(settings, " ")
		for k, v := range s {
			r.ring[k], _ = strconv.Atoi(v)
		}
	} else {
		var arr [Length]int
		for i := 0; i < Length; i++ {
			arr[i] = i
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(Length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		m := make(map[int]int)
		var check [Length]bool

		for i := 0; i < Length; i++ {
			if !check[i] && !check[arr[i]] {
				m[i] = arr[i]
				m[arr[i]] = i
				check[i] = true
				check[arr[i]] = true
			}
		}

		for i := 0; i < Length; i++ {
			if !check[i] {
				m[i] = i
			}
		}

		for k, v := range m {
			r.ring[k] = v
		}
	}

	return r
}

func (r *Rotor) inc() {
	r.pos++
}

func (r *Rotor) upd() {
	r.pos %= Length
}

func (r *Rotor) isFullTurn() bool {
	return r.pos == Length
}

func (r *Rotor) getVal(in int) int {
	return r.ring[(in+r.pos)%Length]
}

func (r *Rotor) getKey(val int) int {
	for i := 0; i < Length; i++ {
		if r.ring[i] == val {
			key := (Length + i - r.pos) % Length
			return key
		}
	}
	return -1
}

func (r Rotor) getRing() [Length]int {
	return r.ring
}
