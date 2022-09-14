package main

import (
	"fmt"
	"time"
	"ytChan/slowchan"
)

func main() {
	//ch, shut := subchan.Default()
	//
	//c := ch.Subscribe("cyt", 1024)
	//
	//for i := 0; i < 10; i++ {
	//	go func(x int) {
	//		_ = ch.Subscribe(strconv.Itoa(x), 1024)
	//	}(i)
	//	go func() {
	//		ch.Send("1")
	//	}()
	//}
	//go func() {
	//	for {
	//		x := <-c
	//		fmt.Println(x)
	//	}
	//}()
	//
	//ch.Send("1")
	//ch.Send("2")
	//ch.Send("3")
	//ch.Send("4")
	//ch.Send("5")
	//
	//time.Sleep(5 * time.Second)
	//
	//ch.Close(shut)

	ch, shut := slowchan.Default()
	ch.Send(1)
	ch.Send(2)
	ch.Send(3)
	ch.Send(4)
	ch.Send(5)
	fmt.Println(ch.History())

	c := ch.Start(1024)
	go func() {
		for {
			x := <-c
			fmt.Println(x)
		}
	}()

	time.Sleep(10 * time.Second)

	ch.Close(shut)
}
