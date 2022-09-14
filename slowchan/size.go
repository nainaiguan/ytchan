package slowchan

func (d *SlowChan) Size() int {
	return len(d.data)
}

func (d *SlowChan) Capacity() int {
	return d.cap
}
