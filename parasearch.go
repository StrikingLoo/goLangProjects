package main

import (
	"fmt"
	"time"
	//	"sync"
)

func isHere(where []int, what int) bool {
	start := time.Now()
	res := false
	//defer fmt.Println(time.After(start))
	for _, e := range where {
		/*
			for j := 0 ; j < 10000 ; j++ {

				continue
			}*/
		//		time.Sleep(1 * time.Millisecond)
		if e == what {
			//fmt.Println(time.Since(start))
			res = true
		}
	}
	fmt.Println(time.Since(start))
	return res
}

func found(where []int, what int, rch chan bool) {
	res := false
	for _, e := range where {
		//	time.Sleep(1 * time.Millisecond)
		/*for j := 0  ; j < 10000; j++ {
		continue
		}*/
		if e == what {
			res = true
		}
	}
	rch <- res
}

func isHereHalf(where []int, what int) bool {
	start := time.Now()
	rch := make(chan bool)
	res := false
	divs := 4
	for i := 0; i < divs; i++ {
		chunkSize := len(where) / divs
		go found(where[i*chunkSize:(i+1)*chunkSize], what, rch)
	}

	for j := 0; j < divs; j++ {
		res = res || <-rch
	}
	fmt.Println(time.Since(start))
	return res
}

func main() {
	l := make([]int, 240000)
	for i := range l {
		l[i] = (191 + 283*i) % 307
		//		l = append(l,44)
	}
	//l = []int{1,2,3,4,5,1,2,3,4,5,1,2,3,4,5,44}
	fmt.Println(isHere(l, 44))
	fmt.Println(isHereHalf(l, 44))

}
