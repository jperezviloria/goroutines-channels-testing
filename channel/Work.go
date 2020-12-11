package channel

import "fmt"

func WorkChannel() {

	fmt.Println("Initialize Work")
	done := make(chan bool)
	go worker(done)
	result := <-done
	fmt.Println(result)

}

func worker(done chan bool) {
	fmt.Println("Inside the worker")
	done <- true
}
