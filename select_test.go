package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Developer"
}

func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
			case data := <-channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
			default:
				fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}