package test

import (
	"fmt"
	"ytchan/dftchan"
)

func test() {
	ch, shut := dftchan.Default()
	for i := 0; i < 100; i++ {
		ch.Send(i)
	}

	fmt.Println(ch.History())
	fmt.Println(ch.Pull(100))

	ch.Close(shut)
}
