package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Reflector struct {
	mapping map[int]int
}

func makeReflector(settings string) Reflector {
	var r Reflector
	r.mapping = make(map[int]int, Length)

	if settings != "" {
		s := strings.Split(settings, " ")
		for _, v := range s {
			elem := strings.Split(v, ":")
			k, _ := strconv.Atoi(elem[0])
			r.mapping[k], _ = strconv.Atoi(elem[1])
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

		for i := 0; i < Length; i += 2 {
			r.mapping[arr[i]] = arr[i+1]
		}

		for k, v := range r.mapping {
			r.mapping[v] = k
		}
	}

	return r
}

func (r *Reflector) getVal(key int) int {
	return r.mapping[key]
}

func (r Reflector) getMapping() map[int]int {
	return r.mapping
}
