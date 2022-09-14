package slowchan

import "time"

func (d *SlowChan) slowChanCleanDaemon() {
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
