package slowchan

import "time"

func (s *SlowChan) Start(bufsize int) chan interface{} {
	c := make(chan interface{}, bufsize)
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				time.Sleep(s.step)

				c <- <-s.data
			}
		}
	}()
	return c
}
