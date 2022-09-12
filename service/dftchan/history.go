package dftchan

func (d *DftChan) History() []interface{} {
	return d.sendHistory.Load()
}
