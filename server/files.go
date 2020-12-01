package server

import (
	"log"
	"os"
	"path/filepath"
)

func (s *Server) stageInputFiles() {
	err := filepath.Walk(s.inputPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			s.unprocessed = append(s.unprocessed, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(MapReduceError{errReadingInputDir, err.Error()})
	}
}

func (s *Server) getNextFile() string {
	if len(s.unprocessed) == 0 {
		return ""
	}

	path := s.popFirstFile()
	s.inflight = append(s.inflight, path)
	return path
}

func (s *Server) popFirstFile() string {
	path := s.unprocessed[0]
	s.unprocessed = s.unprocessed[1:]
	return path
}

func isEmpty(path string) bool {
	return len(path) == 0
}
