package golang_goroutine

import (
	"fmt"
	"time"
	"testing"
	"sync"
)

func TestRaceCondition(t *testing.T){
	x := 0

	for i := 1; i <= 1000; i++ {
		go func(){
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

func TestRaceConditionWithMutext(t *testing.T){
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func(){
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println("Balance = ", account.GetBalance())
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Total Balance = ", account.GetBalance())
}