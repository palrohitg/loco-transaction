package main

import (
	"fmt"
	"sync"
)

/*

	Elastic Search Important Concepts to Remember:
		1. Apache lucene, now the integrations has the pipeline to work on other stuff as wells.
		2. Uses the Inverted Indices: to maps to the words of the data.
		3. key name --> Document Id present where is the data is presents.

	Dependency Injection:
		- Creating objects that can be configurable is ones the ability to solve the
			dependency injection problems
		- Find the changes


	Synchronization Techniques:
	1. Mutex Lock and Unlock methods for the sames. (Shared resources) to avoid the unpredictable outcomes and finally
		avoid the same the resources , which has been access by multiple go-threads.
	2. Atomic Operations so others goroutines can't modify the data.
	3. WaitGroups : to synchronize the data between all the goroutines.


	Channels :
		1. Unbuffered : No capacity sender --> receiver should be ready to receive the data.
			ch := make(chan int) // Creating an unbuffered channel
			go func() {
				ch <- 42 // Sending a value to the channel
			}()

    		value := <-ch /
		2. Buffered: Having a finite capacity to send and receive the data.
			    ch := make(chan int, 3) // Creating a buffered channel with capacity 3
				ch <- 1
				ch <- 2
				ch <- 3
		3. Find the Greatest elements you have till nows.
*/

/*
	Data structures Indexes
	1. BTree:
		a. Key, and pointers : data is store key leaf + internal nodes, pointer on the leaf node (Index)
		b. Each node is same : BP, key and other nodes.
		c. Searching is slow,
	2. B+tree:
		a. Data is stored only in the leaf nodes.
		b. Searching is faster, deletion is fasters.
		c. leaf node is attached with each others : which is used in the range querys.
		d.


*/

// No of People in the groups we have
var ph = []string{"v1", "v1", "v1", "v1"}

// Process Dining Problem : Processors Synchronizations in Operating Systems
var dining sync.WaitGroup

// Mutex  helps us to prevents the shared resources using in the concurrency
// Environments
func main() {
	dining.Add(5)
	fmt.Println(ph)
	return
}
