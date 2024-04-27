package gorountine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Digunakan saat ada sharing variable dlm gorountine
func TestMutex(t *testing.T) {
	var mutex sync.Mutex

	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(x)
}

type Account struct {
	lock    sync.RWMutex
	balance int
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance -= amount

}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

func TestRwMutex(t *testing.T)
