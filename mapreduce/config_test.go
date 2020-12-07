package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// configtest.cfg:
// ------------
// mapper ./test_files/executables/word-count-mapper
// reducer ./test_files/executables/word-count-reducer
// num-mappers 1
// num-reducers 1
// input-path test_files/input
// intermediate-path test_files/intermediate
// output-path test_files/output

func TestParseConfigCorrectCase(t *testing.T) {
	server := Server{}
	server.initializeFromConfigFile("test_files/config_test.cfg")

	assert.EqualValues(t, "./test_files/executables/word-count-mapper", server.mapperExecutable)
	assert.EqualValues(t, "./test_files/executables/word-count-reducer", server.reducerExecutable)

	assert.EqualValues(t, 1, server.numMappers)
	assert.EqualValues(t, 1, server.numReducers)

	assert.EqualValues(t, "test_files/input", server.inputPath)
	assert.EqualValues(t, "test_files/intermediate", server.intermediatePath)
	assert.EqualValues(t, "test_files/output", server.outputPath)
}
