package main

import (
	"fmt"
	"time"
)

func task(name string, channel chan int, executions *int) {
	for {
		value := <-channel
		fmt.Println("Executing task", value, "by", name)
		*executions++
	}
}

func main() {

	channel := make(chan int)
	executions := 0

	go task("goroutine 1", channel, &executions)
	go task("goroutine 2", channel, &executions)
	go task("goroutine 3", channel, &executions)
	go task("goroutine 4", channel, &executions)
	go task("goroutine 5", channel, &executions)

	for i := 0; i < 100; i++ {
		channel <- i
	}
	// wait all gorotines to finish
	time.Sleep(time.Second)

	fmt.Println(executions)
}
