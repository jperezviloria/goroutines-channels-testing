package main

import (
	"fmt"
	"goroutines-test/calcutation"
)

func main() {
	//First execution
	//goroutineOne.GoroutineOne()

	//Second execution
	//web.Web()

	//First Channel execution
	//channel.ChannelOne()

	//Second Buffer Channel execution
	//channel.BufferChannel()

	//Third Work Channel execution
	//channel.WorkChannel()

	//Closing Channel
	//channelTwo.ClosingChannel()

	//range channel

	//channel directions

	result := calcutation.Calculate(2, 2)
	fmt.Println(result)
}
