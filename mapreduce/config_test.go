package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// example.cfg:
// ------------
// mapper ./mapreduce/files/executables/word-count-mapper
// reducer ./mapreduce/files/executables/word-count-reducer
// num-mappers 1
// num-reducers 1
// input-path ./mapreduce/files/input
// intermediate-path ./mapreduce/files/intermediate
// output-path ./mapreduce/files/output

func TestParseConfigCorrectCase(t *testing.T) {
	server := Server{}
	server.initializeFromConfigFile("files/example.cfg")

	assert.EqualValues(t, "./mapreduce/files/executables/word-count-mapper", server.mapperExecutable)
	assert.EqualValues(t, "./mapreduce/files/executables/word-count-reducer", server.reducerExecutable)

	assert.EqualValues(t, 1, server.numMappers)
	assert.EqualValues(t, 1, server.numReducers)

	assert.EqualValues(t, "./mapreduce/files/input", server.inputPath)
	assert.EqualValues(t, "./mapreduce/files/intermediate", server.intermediatePath)
	assert.EqualValues(t, "./mapreduce/files/output", server.outputPath)
}
