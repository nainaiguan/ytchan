package slowchan

import "context"

func (d *SlowChan) Close(cancelFunc context.CancelFunc) {
	cancelFunc()
	d.closeFlag.Close()

	for {
		if d.sendProcess.Load() == 0 {
			break
		}
	}

	close(d.data)
}
