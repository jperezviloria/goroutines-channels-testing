package goroutineOne

import (
	"fmt"
	"runtime"
	"time"
)

func GoroutineOne() {
	cpus := runtime.NumCPU() //4
	runtime.GOMAXPROCS(cpus)
	fmt.Println(cpus)

	numGR := runtime.NumGoroutine()
	fmt.Println(numGR)

	go func() {

		for i := 0; i < 20; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i)
		}
	}()

	go func() {

		for j := 21; j < 40; j++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(j)
		}
	}()
	go func() {

		for j := 41; j < 60; j++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(j)
		}
	}()

	numGR2 := runtime.NumGoroutine()
	fmt.Println(numGR2)

	go func() {

		for j := 61; j < 80; j++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(j)
		}
	}()
	go func() {

		for j := 81; j < 100; j++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(j)
		}
	}()

	//time.Sleep(5 * time.Second)
	numGR3 := runtime.NumGoroutine()
	fmt.Println(numGR3)
}
