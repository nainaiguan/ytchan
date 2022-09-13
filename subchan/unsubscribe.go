package subchan

func (s *SubChan) Unsubscribe(name string) {
	s.subscriber.Drop(name)
}
