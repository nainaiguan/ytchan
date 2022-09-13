package subchan

import "time"

func (d *SubChan) reconcile() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			if len(d.data) != 0 {
				message := <-d.data
				for _, c := range d.subscriber.m {
					c <- message
				}
			}
		}
	}
}

func (d *SubChan) subChanCleanDaemon() {
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
