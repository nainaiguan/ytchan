package main

import (
	"fmt"
	"sync"
	"time"
	"ytChan/subchan"
)

func main() {
	ch, shut := subchan.Default()

	c := ch.Subscribe("cyt", 1024)

	go func() {
		for {
			x := <-c
			fmt.Println(x)
		}
	}()

	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 1020; i++ {
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			ch.Send(i)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(3 * time.Second)
	cond.Broadcast()

	time.Sleep(5 * time.Second)

	ch.Close(shut)

	//ch, shut := slowchan.Default()
	//ch.Send(1)
	//ch.Send(2)
	//ch.Send(3)
	//ch.Send(4)
	//ch.Send(5)
	//fmt.Println(ch.History())
	//
	//c := ch.Start(1024)
	//go func() {
	//	for {
	//		x := <-c
	//		fmt.Println(x)
	//	}
	//}()
	//
	//time.Sleep(10 * time.Second)
	//
	//ch.Close(shut)
}
