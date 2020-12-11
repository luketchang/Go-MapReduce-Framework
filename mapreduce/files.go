package mapreduce

import (
	"log"
	"os"
	"path/filepath"
)

var ConfigFileKeys = []string{
	"mapper",
	"reducer",
	"num-mappers",
	"num-reducers",
	"input-path",
	"intermediate-path",
	"output-path",
}

func (s *Server) stageInputFiles() {
	err := filepath.Walk(s.inputDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			s.unprocessed = append(s.unprocessed, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(MapReduceError{errReadingInputDir, err})
	}
}

func (s *Server) getUnprocessedFilePattern() string {
	if len(s.unprocessed) == 0 {
		return ""
	}

	path := s.popUnprocessedList()
	s.inflight[path] = true
	return path
}

func (s *Server) popUnprocessedList() string {
	path := s.unprocessed[0]
	s.unprocessed = s.unprocessed[1:]
	return path
}

func (s *Server) markFilePatternAsProcessed(remoteIPAddr string, filePattern string) {
	delete(s.inflight, filePattern)
	log.Println(filePattern + " successfully processed by " + s.ipAddressMap[remoteIPAddr])
}

func (s *Server) rescheduleFilePatternJob(remoteIPAddr string, filePattern string) {
	s.unprocessed = append(s.unprocessed, filePattern)
	delete(s.inflight, filePattern)
	log.Println(filePattern + " failed to be processed by " + s.ipAddressMap[remoteIPAddr] + ". rescheduling...")
}

func isEmpty(path string) bool {
	return len(path) == 0
}
