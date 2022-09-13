package subchan

func (s *SubChan) History() []interface{} {
	return s.sendHistory.Load()
}
