package subchan

func (d *subChan) Size() int {
	return len(d.data)
}

func (d *subChan) Capacity() int {
	return d.cap
}
