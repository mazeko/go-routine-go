package golang_goroutine

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func RuncAsync(group *sync.WaitGroup, number int){
	defer group.Done()

	group.Add(1)
	fmt.Println("Hey", number)
	time.Sleep(1 * time.Second)
}

func TestWaitrGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RuncAsync(group, i)
	}

	group.Wait()
	fmt.Println("complete")
}

