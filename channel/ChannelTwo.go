package channel

import "fmt"

func BufferChannel() {

	c := make(chan string, 2)
	c <- "one"
	c <- "two"

	fmt.Println(<-c)
	fmt.Println(<-c)

}
