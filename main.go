package main

import (
	"fmt"
	"time"
	dftchan2 "ytChan/dftchan"
)

func main() {
	ch, shut := dftchan2.Default()

	for i := 0; i < 102; i++ {
		ch.Send(1)
		ch.Send(1)
		fmt.Println(ch.Pull(1))
	}

	time.Sleep(40 * time.Second)
	fmt.Println(cap(ch.History()))
	ch.Close(shut)
}
