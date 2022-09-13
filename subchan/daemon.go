package subchan

import "time"

func (d *SubChan) Reconcile() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)

			if len(d.data) != 0 {
				message := <-d.data
				for _, c := range d.subscriber {
					c <- message
				}
			}
		}
	}
}

func (d *SubChan) SubChanCleanDaemon() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			d.cleanFlag.Clean()
			l := len(d.sendHistory.h)
			tmp := make([]interface{}, l)
			copy(tmp, d.sendHistory.h)
			d.sendHistory.h = tmp
			d.cleanFlag.Done()

			time.Sleep(30 * time.Second)
		}
	}
}
