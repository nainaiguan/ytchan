package subchan

import "context"

func (s *subChan) Close(cancelFunc context.CancelFunc) {
	cancelFunc()

	s.closeFlag.Close()

	for {
		if s.sendProcess.Load() == 0 {
			break
		}
	}

	close(s.data)
}
