package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// example.cfg:
// ------------
// mapper ./files/executables/word-count-mapper
// reducer ./files/executables/word-count-reducer
// num-mappers 1
// num-reducers 1
// input-path files/input
// intermediate-path files/intermediate
// output-path files/output

func TestParseConfigCorrectCase(t *testing.T) {
	server := Server{}
	server.initializeFromConfigFile("files/example.cfg")

	assert.EqualValues(t, "./files/executables/word-count-mapper", server.mapperExecutable)
	assert.EqualValues(t, "./files/executables/word-count-reducer", server.reducerExecutable)

	assert.EqualValues(t, 1, server.numMappers)
	assert.EqualValues(t, 1, server.numReducers)

	assert.EqualValues(t, "files/input", server.inputPath)
	assert.EqualValues(t, "files/intermediate", server.intermediatePath)
	assert.EqualValues(t, "files/output", server.outputPath)
}
