package subchan

func (s *SubChan) Subscribe(name string, size int) chan interface{} {
	s.subscriber.Add(name, size)
	return s.subscriber.Load(name)
}
