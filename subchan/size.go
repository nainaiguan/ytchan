package subchan

func (d *SubChan) Size() int {
	return len(d.data)
}

func (d *SubChan) Capacity() int {
	return d.cap
}
