package slowchan

func (d *slowChan) History() []interface{} {
	return d.sendHistory.Load()
}
