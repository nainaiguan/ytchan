package dftchan

func (d *DftChan) Pull(size int) []interface{} {
	d.pullProcess.Add()

	ret := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		m := <-d.data
		ret = append(ret, m)
	}

	d.pullProcess.Done()

	return ret
}
