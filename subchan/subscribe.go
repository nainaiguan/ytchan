package subchan

func (s *subChan) Subscribe(name string, size int) chan interface{} {
	s.subscriber.Add(name, size)
	return s.subscriber.Load(name)
}
