package main

import (
	"fmt"
	"time"
)

//channel types
// chan T
// bidirectional channel type
// allows both recieving and sending values to bidirection channels
//
//
// chan<- T
// send only channel
//
// <-chan T
// recieve only channel

// bidirectional can be changed to send or recived but never the opposite

func mainChannel() {

	// ch := make(chan int, 6)
	//
	// ch <- 3
	// ch <- 4
	// ch <- 5
	// ch <- 6
	// ch <- 7
	//
	// close(ch)
	//
	// v := <-ch
	// fmt.Println(v)
	// channelPrinter(ch)
	// channelPrinter(ch)
	tasks := []string{"t1", "t2", "t3"}

	c := make(chan string, 3)

	for _, task := range tasks {
		c <- task
	}

	//c <- "t4" this panics

	go worker(c)
	//chanSyncExample()

	bufChan := make(chan int, 3)

	for i := 0; i < 5; i++ {
		go bufWorker(i, bufChan)
		bufChan <- i
	}

	time.Sleep(time.Second * 7)

	close(bufChan)

	// ranging over channels

	chr := make(chan int, 3)

	chr <- 1
	chr <- 2
	chr <- 3

	close(chr)

	for v := range chr {
		fmt.Println(v)
	}

	selectBuf()
	//v, open := <-ch Open bool is also sent when ch is recieved
	// if !open {
	// 	fmt.Println("ch v is closed")
	// }

}

func selectBuf() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- 1
		time.Sleep(time.Second * 3)
		ch1 <- 5
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("rec from ch1: ", msg1)

		case msg2 := <-ch2:
			fmt.Println("rec from ch2: ", msg2)
		default:
			fmt.Println("no msg")
		}
	}
}

func bufWorker(id int, ch chan int) {

	fmt.Printf("\n worker %d started here \n", id)
	time.Sleep(time.Second)
	fmt.Printf("\n worker %d finished here \n", id)
	<-ch

}

func chanSyncExample() {
	message := make(chan string)

	for i := 0; i < 10; i++ {

		time.Sleep(time.Second * 2)
		go func() {

			message <- "hi there"

		}()

		msg := <-message
		fmt.Println(msg)
	}
	close(message)
}

func channelPrinter(ch <-chan int) {
	channelData := []int{}
	for value := range ch {
		channelData = append(channelData, value)
	}

	fmt.Printf("Channel data: %v\n", channelData)
}

func worker(ch chan string) {
	for {
		t := <-ch
		process(t)
	}
}

func process(t string) {
}
