package subchan

func (s *subChan) Unsubscribe(name string) {
	s.subscriber.Drop(name)
}
