package dftchan

func (d *dftChan) Size() int {
	return len(d.data)
}

func (d *dftChan) Capacity() int {
	return d.cap
}
