package main

import (
	"fmt"
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

	ch.Send("1")
	ch.Send("2")
	ch.Send("3")
	ch.Send("4")
	ch.Send("5")

	time.Sleep(50 * time.Second)

	ch.Close(shut)
}
