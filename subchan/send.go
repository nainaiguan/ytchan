package subchan

import (
	"errors"
	"ytChan/util/prettylog"
)

func (d *subChan) Send(message interface{}) error {
	if d.closeFlag.Load() == 1 {
		err := errors.New("the chan is already closed")
		prettylog.Errorf("subChan.Send Error, err: %s", err)
		return err
	}

	if d.maxSendProcess <= d.sendProcess.Load() {
		err := errors.New("too much sendProcess")
		prettylog.Errorf("subChan.Send Error, err: %s", err)
		return err
	}

	if d.cap <= d.Size() {
		prettylog.Infof("out of the chan range")
		return nil
	}

	d.cleanFlag.Load()
	defer d.cleanFlag.Free()

	d.sendProcess.Add()
	d.data <- message
	d.sendHistory.Add(message)
	for _, c := range d.subscriber.m {
		c <- <-d.data
	}
	d.sendProcess.Done()

	return nil
}
