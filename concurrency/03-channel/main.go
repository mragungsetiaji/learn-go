package main

import (
	"fmt"
	"time"
)

func main() {
	data := [3]string{"a1", "a2", "a3"}
	ChanVar := make(chan string)

	// sender
	go func(mine [3]string) {
		for _, item := range mine {
			ChanVar <- item //send
		}
	}(data)

	// receiver
	go func() {
		for i := 0; i < 3; i++ {
			result := <-ChanVar //receive
			fmt.Println("Receiver: Received " + result + " from sender")
		}
	}()

	<-time.After(time.Second * 5) // wait for 5 seconds
}
