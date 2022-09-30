package main

import (
	"fmt"
	"time"
)

func main() {
	var globalCount = 0
	go hello("First", &globalCount)
	go hello("Second", &globalCount)
	go hello("Third", &globalCount)
	go hello("Forth", &globalCount)

	time.Sleep(time.Second)
}

func hello(name string, count *int) {
	(*count)++
	fmt.Printf("Hello %s %d \n", name, *count)
}

/**

Race condition   is when several goroutines share
common data and modify it dependig timing of other
uncontrollable events, which leads to different from
expected value.

A race condition in Go occurs when two or more goroutines
have shared data and interact with it simultaneously.
In my case it occured when several goroutines started their
function, edited 'count' but did not finish printing


*/
