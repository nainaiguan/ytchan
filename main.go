package main

import (
	"fmt"
	"sync"
	"time"
	"ytChan/api/sub"
	"ytChan/subchan"
)

func main() {
	ch, shut := subchan.New(sub.NewSubArgs{
		Size:           10000,
		MaxSendProcess: 10000,
	})

	c := ch.Subscribe("cyt", 1024)

	go func() {
		for {
			x := <-c
			fmt.Println(x)
		}
	}()

	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10000; i++ {
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
			err := ch.Send(i)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
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
