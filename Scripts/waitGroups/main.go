package main

import (
	"fmt"
	"sync"
)

/*
	WaitGroups helps us to
	ADD : Add methods is used to add a  counters
	Done : Decrement the counter by one single no this is how it is working on golang
	Wait
*/

/*
Higher Order Functions in Golang
1.

Formatting in the Golang Code:
1. %d
2. %t
3. %g
4. %e

In case of multiple defer statement there
as the tools which is going to LIFO order is the actual

a. What is Functions Closures In golang ?
b. GoRoutine helps to do the things are faster than

	concurrently we have in the systems, which we are doing in the ways
	to execute the tasks, if more than processors are busy in the
	I/O Tasks then others tasks are  schedule to handle those
	stuff efficiently.

Threads : 1MB
Go : 2KB

Threads Setup and TearDown Costs is higher
costs  because it required the Operating systems and return's
once its done.

Mutex :

	- Mutex : Mutual Exclusion
		a.Synchronization primitive in Golang, that allows mutliple goroutine
			to corrdinate and protect the shared resources.
		b. yeh ensure krta hai ke kio ek

Goroutine ko kio resource access krna tho
acquire lock krta hai functions pr
otherwise unlock

two main methods hai
lock()
unlock()

kio bhi shared resources pr kaam krne ke.


GoRoutine Pool in Go
*/
var wg sync.WaitGroup

func responseSize(url string) {
	defer wg.Done() // this will be decrement the counter by on
	fmt.Println(url)

}

func main() {

	wg.Add(3)

	fmt.Println("Started Goroutines")

	go responseSize("testing")
	go responseSize("testing")
	go responseSize("testing")

	wg.Wait()

	return
}
