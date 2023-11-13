package main

import (
	"flag"
	"fmt"
)

type Consumer struct {
	msgs *chan int
	done *chan bool
}

func NewConsumer(msgs *chan int, done *chan bool) *Consumer {
	return &Consumer{msgs: msgs, done: done}
}

// Infinite Listening for the new events
func (c *Consumer) consume() {
	fmt.Println("consume : Started")
	// this is the for each loop we are  using it here
	for {
		msg, more := <-*c.msgs
		if !more {
			fmt.Println("consume : done")
			*c.done <- true
			return
		}
		//time.Sleep(1 * time.Second)
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

//func (p *Producer) produce(max int) {
//	fmt.Println("produce : started")
//	for i := 0; i < max; i++ {
//		*p.msgs <- i // writing message here
//		fmt.Println("produce : Sending", i)
//	}
//	*p.done <- true // writing to the done statements we have
//	fmt.Println("produce : Done")
//}

func (p *Producer) produce(max int) {
	fmt.Println("produce : started")
	for i := 0; i < max; i++ {
		*p.msgs <- i // writing message here
		fmt.Println("produce : Sending", i)
	}
	//*p.done <- true // writing to the done statements we have
	close(*p.msgs) // signals to close the channels now
	//fmt.Println("produce : Done")
}

func main() {
	max := flag.Int("n", 1, "defines the number of messages")
	var msgs = make(chan int)  // create one message channel
	var done = make(chan bool) // create one done channel

	go NewProducer(&msgs, &done).produce(*max)
	//time.Sleep(1 * time.Second)
	go NewConsumer(&msgs, &done).consume()

	<-done // Once all the Message is produce done
	fmt.Println("I'm here is the last statement here")
}
