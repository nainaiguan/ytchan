package test

import (
	"fmt"
	"time"
	"ytchan/api/dft"
	"ytchan/dftchan"
)

func test() {
	ch, shut := dftchan.New(dft.NewDftArgs{
		Size:           10000,
		MaxSendProcess: 100,
	})
	for i := 0; i < 1500; i++ {
		go ch.History()
		go ch.Send(i)
	}

	fmt.Println(ch.History())

	for i := 0; i < 10; i++ {
		fmt.Println(ch.Pull(1))
		time.Sleep(1 * time.Second)
	}
	fmt.Println(ch.History())

	ch.Close(shut)
}
