package main

import "fmt"

type Consumer struct {
	msgs *chan int
}

func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

func (c *Consumer) consume() {
	fmt.Println("consume : Started")
	for {
		msg := <-*c.msgs
		fmt.Println("consume : Received:", msg)
	}
}

// Producer Definition
type Producer struct {
	msgs *chan int
	done *chan bool
}

func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

func (p *Producer) produce(max int) {
	fmt.Println("produce : started")
	for i := 0; i < max; i++ {
		fmt.Println("produce : Sending", i)
		*p.msgs <- i
	}
	*p.done <- true
	fmt.Println("produce : Done")
}
func main() {

}
