package slowchan

import (
	"time"
	"ytchan/util/prettylog"
)

func (d *slowChan) slowChanCleanDaemon() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			d.cleanFlag.Clean()
			count := d.sendHistory.Len()
			if count < 1024 {
				prettylog.Infof("no need to clean history")
			} else if count < 8192 {
				tmp := make([]interface{}, count)
				copy(tmp, d.sendHistory.h[count/4:])
				d.sendHistory.h = tmp
				prettylog.Infof("clean complete: %s", count)
			} else {
				tmp := make([]interface{}, count)
				copy(tmp, d.sendHistory.h[count/2:])
				d.sendHistory.h = tmp
				prettylog.Infof("clean complete: %d", count)
			}
			d.cleanFlag.Done()

			time.Sleep(3 * time.Second)
		}
	}
}
