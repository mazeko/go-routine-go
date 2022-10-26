package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Oke")

	time.Sleep(3 * time.Second)
}
