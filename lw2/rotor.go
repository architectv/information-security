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
		var ring [Length]int
		for i := 0; i < Length; i++ {
			ring[i] = i
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(Length, func(i, j int) {
			ring[i], ring[j] = ring[j], ring[i]
		})

		for i := 0; i < Length; i++ {
			if ring[ring[i]] != i {
				tmp := ring[ring[i]]
				ring[ring[i]] = i
				for j := 0; j < Length; j++ {
					if ring[j] == i && ring[i] != j {
						ring[j] = tmp
					}
				}
			}
		}

		r.ring = ring
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
