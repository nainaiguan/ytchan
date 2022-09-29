package subchan

func (s *subChan) History() []interface{} {
	return s.sendHistory.Load()
}
