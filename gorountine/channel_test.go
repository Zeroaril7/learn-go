package gorountine

import (
	"fmt"
	"testing"
	"time"
)

/*
By Default channel merupakan by reference

Mendeklarasikan cukup dengan make(chan type data)
*/
func Test_Channel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Hello"
	}()

	fmt.Println(<-channel)

}

func ParamChannel(msg chan string, value string) {
	msg <- value
}

func Test_ParamChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		ParamChannel(channel, "Hi, World")
	}()

	fmt.Println(<-channel)
}

func OnlyIn(channel chan<- string) {
	channel <- "In Out"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func Test_InOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func Test_BufferChannel(t *testing.T) {

	// angka 3 adalah ukuran dari buffer
	channel := make(chan string, 3)

	go func() {
		channel <- "Test 1"
		channel <- "Test 2"
		channel <- "Test 3"
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

}

func Test_SelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		channel1 <- "Food"
	}()

	go func() {
		channel2 <- "beverage"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		default:
			continue
		}

		if counter == 2 {
			break
		}
	}

}
