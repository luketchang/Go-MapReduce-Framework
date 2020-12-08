package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStageFiles(t *testing.T) {
	s := Server{}
	s.inputDir = "test_files/input"
	s.stageInputFiles()

	assert.Equal(t, s.unprocessed, []string{
		"test_files/input/00001.input",
		"test_files/input/00002.input",
		"test_files/input/00003.input"})
}

func TestGetNextFileValid(t *testing.T) {
	s := Server{}
	s.inputDir = "test_files/input/"
	s.stageInputFiles()

	file := s.getNextFile()
	assert.Equal(t, file, "test_files/input/00001.input")
	assert.Equal(t, s.inflight[0], "test_files/input/00001.input")
	assert.Equal(t, s.unprocessed[0], "test_files/input/00002.input")
}
