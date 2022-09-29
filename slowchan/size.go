package slowchan

func (d *slowChan) Size() int {
	return len(d.data)
}

func (d *slowChan) Capacity() int {
	return d.cap
}
