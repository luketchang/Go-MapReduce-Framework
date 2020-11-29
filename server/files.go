package server

func (s *Server) getNextFile(path *string) bool {
	if len(s.unprocessed) == 0 {
		return false
	}

	s.popFirstFile(path)
	s.inflight = append(s.inflight, *path)
	return true
}

func (s *Server) popFirstFile(path *string) {
	*path = s.unprocessed[0]
	s.unprocessed = s.unprocessed[1:]
}
