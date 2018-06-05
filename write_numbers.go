package main

import (
	"fmt"
	"time"
)

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	N := 10

	go func() {
		for {
			v := <-oddChan
			fmt.Println(v)
			if ( v <= N) {
				evenChan <- v + 1
			} else {
				break
			}
		}
	}()

	go func() {
		for {
			v := <-evenChan
			fmt.Println(v)
			if (v <= N) {
				oddChan <- v + 1;
			} else {break}
		}
	}()

	oddChan <- 1
	time.Sleep(2)
	fmt.Print("Hello WorldXXXX")
}
