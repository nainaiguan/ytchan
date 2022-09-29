package dftchan

func (d *dftChan) History() []interface{} {
	return d.sendHistory.Load()
}
