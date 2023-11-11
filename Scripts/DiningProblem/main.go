package main

import (
	"fmt"
	"sync"
)

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
