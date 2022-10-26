package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
	"strconv"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Developer"
}

func TestChannelAsParams(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Eko"
	}()
	data := <-channel
	fmt.Println(data)
	close(channel)
}

func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Mazeko"
}

func OnlyOut(channel <-chan string){
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Eko"
	channel <- "Oke"
	channel <- "Oke"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i < 10; i++ {
			channel <- "Nomor " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Data ke ", data)
	}

	fmt.Println("Done")
}