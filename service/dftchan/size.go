package dftchan

func (d *DftChan) Size() int {
	return len(d.data)
}

func (d *DftChan) Capacity() int {
	return 0
}
