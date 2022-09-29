# ytChan: An easy message queue plugin.

## Introduction
### Include three kinds of queue: dftchan, slowchan, subchan
All queues support normal sending and pulling of messages of any type. 
The history function is built into the queue, and the history pool is maintained by a daemon thread. 
All queues support custom maximum concurrency. 
In theory, the maximum concurrency you set must be within the tolerance of the official golang channel.

- ### dftchan
```
package main

import (
	"fmt"
	"github.com/ytChan/dftchan"
)

func main() {
    ch, shut := dftchan.Default()

    fmt.Println(ch.Size())
    fmt.Println(ch.Capacity())

    _ = ch.Send("1")
    _ = ch.Send("2")

    fmt.Println(ch.History())

    fmt.Println(ch.Pull(1))
    fmt.Println(ch.Pull(2))

    ch.Close(shut)
}
```

- ### subchan
Support subscription function, each message will be pushed to all subscribed users.
```
package main

import (
	"fmt"
	"sync"
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
	for i := 0; i < 1000; i++ {
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
		
	ch.Close(shut)
}
```
- ### slowchan
Support buffering function, users can set the message push interval by themselves.
```
package main

import (
	"fmt"
	"time"
	"ytChan/slowchan"
)

func main() {
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

```