package channelTwo

import (
	"fmt"
	"math/rand"
)

func ClosingChannel() {

	randChannel := make(chan int)
	go bufferOne(randChannel)
	var count int
	for {
		count++
		result, ok := <-randChannel
		if !ok {
			break
		}
		fmt.Printf("index %d and value %d \n", count, result)
	}
	fmt.Println("out of loop")

}

func bufferOne(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Int()
	}
	close(c)
}
