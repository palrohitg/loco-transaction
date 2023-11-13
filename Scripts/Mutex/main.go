package main

import (
	"fmt"
	"sync"
)

/*

	Mutex is a lock for the sync.
	access of the resources we have in code
	type Shape interface {
		area() float64
	}

	type Rectangle struct {
		Width float64
		Height float64
	}


	func (r rectangle) area() float64 {
		return r.Width * r.Height
	}

	func calculateArea(s shape) {
		fmt.Printf("Area : %f\n",s.area())
	}
*/

var (
	mutex   sync.Mutex
	balance int
)

// this functions automatically calls when the before
// the main functions and initialize all the values
// which are required
func init() {
	balance = 1000
}

func withDraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("WithDraw %d to account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Deposting %d to account balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

// Higher Order functions which receive the
// Arguments and find the new ones etc.
func applyOperation(value int, operation func(int) int) int {
	return operation(value)
}

func double(x int) int {
	return x * 2
}

func main() {
	results := applyOperation(5, double)
	fmt.Println(results)
	fmt.Println("Go Mutex Example")
	var wg sync.WaitGroup // declared
	wg.Add(2)             // add, Done , wait()
	go withDraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait() // wait until all the go-routines will be done
	fmt.Printf("New balance %d\n", balance)
	return
}
