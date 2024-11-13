package main

import (
	"fmt"
	"sync"
)

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan int
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan int),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan int {

	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan int, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic string, value int) {

	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subsriber := range ps.subscribers[topic] {
		subsriber <- value
	}

}

func (ps *PubSub) Close(topic string, subCh <-chan int) {

	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers, found := ps.subscribers[topic]

	if !found {
		fmt.Println("no subscribes with channels found")
		return
	}

	fmt.Printf("\n%+#v\n", ps)

	for i, subscriber := range subscribers {
		if subscriber == subCh {
			close(subscriber)

			ps.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

func mainPubSub() {

	ps := NewPubSub()

	subscriber := ps.Subscribe("t1")

	go func() {
		ps.Publish("t1", 1412)
	}()

	value := <-subscriber
	fmt.Println("recieved from sub: ", value)

	ps.Close("t1", subscriber)
	//ps.Close("t1", subscriber)

}
