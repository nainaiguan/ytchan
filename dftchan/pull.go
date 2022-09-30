package dftchan

import (
	"ytchan/util/prettylog"
)

func (d *dftChan) Pull(size int) []interface{} {
	if size > d.Size() {
		prettylog.Infof("dftchan.Pull Info: %s", "the chan is empty now")
		return nil
	}

	d.pullProcess.Add()

	ret := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		m := <-d.data
		ret = append(ret, m)
	}

	d.pullProcess.Done()

	return ret
}
