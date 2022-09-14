package slowchan

func (d *SlowChan) History() []interface{} {
	return d.sendHistory.Load()
}
