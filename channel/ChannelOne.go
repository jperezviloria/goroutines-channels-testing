package channel

import "fmt"

func ChannelOne() {
	c := make(chan string)

	go func() {
		c <- "hello from channel"
	}()

	res := <-c

	fmt.Println(res)
}
