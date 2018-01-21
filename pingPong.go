package main

import(
	"fmt"
	"os"
	)

func ping(ch,ch2 chan struct{}, v *int){
	<-ch
	fmt.Println("ping")
	*v++
	ch2<-struct{}{}
	}

func pong(ch2,done chan struct{}, v *int){
	fmt.Println("such concurrent")
	fmt.Println("much wow")
	<-ch2
	fmt.Println("pong")
	fmt.Println(*v)
	done<-struct{}{}
	}
func main(){
	var counter = 0
	ch := make(chan struct{},1)
	ch2 := make(chan struct{},1)
	done := make(chan struct{})

	for {
	ch <- struct{}{}
	go ping(ch,ch2,&counter)
	go pong(ch2,done,&counter) 
	<-done
	}
