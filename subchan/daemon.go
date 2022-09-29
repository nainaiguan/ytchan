package subchan

import (
	"time"
	"ytChan/util/prettylog"
)

func (d *SubChan) subChanCleanDaemon() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			d.cleanFlag.Clean()
			count := len(d.sendHistory.h)
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
				prettylog.Infof("clean complete: %s", count)
			}
			d.cleanFlag.Done()

			time.Sleep(3 * time.Second)
		}
	}
}
