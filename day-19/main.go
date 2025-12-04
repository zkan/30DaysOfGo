package main

import (
	"fmt"
	"time"
)

func doSomething() {
	time.Sleep(3 * time.Second)
	fmt.Println("do something")
}

func main() {
	fmt.Println("start")
	go doSomething()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("inside local func")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
	time.Sleep(5 * time.Second)
}
