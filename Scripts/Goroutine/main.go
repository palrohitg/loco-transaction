package main

import (
	"fmt"
	"time"
)

/*
	Find the concurrency and others topics of golang
	Go Routine : A goroutine is a functions is runs independently that functions is started here.
	Channel : A Pipeline of sending and receiving  the data.
	Concurrency  : Handles to run multiple tasks at once, but only one will be running each times.
	Parellisms :  There are many tasks runs at the same times.
	Gorilla/Websocket is using
	ch := make(chan int, 3) // Buffered channel with a capacity of 3
	Has some kinds of capacities while.
	while the unbuffered channels don't have the single data you can share during the whole
	processors.
*/

func execute(some string) {
	for i := 1; true; i++ {
		fmt.Println(i)
		fmt.Println(some)
		time.Sleep(time.Millisecond * 100)
	}
}

func worker2(id int) {
	fmt.Printf("Worker started %d \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker done %d\n", id)
}

// Channels in the code itself and usage of the same

func main() {
	//go execute("first")
	//go execute("second")
	//fmt.Println("Finally printed here")

	// WaitGroups Concepts and When all the GoRoutine is not executed Implementation here
	//var wg sync.WaitGroup
	//
	//for i := 1; i < 5; i++ {
	//	wg.Add(1)`
	//	first := i
	//	go func() {
	//		defer wg.Done()
	//		worker2(first)
	//	}()
	//}
	//wg.Wait() // waiting for the process to complete the tasks here
	//return
	data, abrr := getAmountOfRewards(5, 5, 125)
	fmt.Println(data)
	fmt.Println(abrr) // 8,5,5 7,5,5

}

func getAmountOfRewards(correctAnwerCount int, totalQuestion int, streakCount int) (int, string) {
	data := map[int]int{
		1: 10,
		2: 12,
		3: 14,
		4: 16,
		5: 18,
		6: 30,
		7: 100,
	}
	dataMapFinal := map[int]string{
		1: "C1DSV",
		2: "C2DSV",
		3: "C3DSV",
		4: "C4DSV",
		5: "C5DSV",
		6: "C6DSV",
		7: "C7DSVs",
	}

	if streakCount > 0 {
		streakCount = (streakCount % 7)
		if streakCount == 0 {
			streakCount = 7
		}
	}
	ml := data[streakCount]
	abbreviation := dataMapFinal[streakCount]
	return correctAnwerCount * ml, abbreviation
}

/**/
