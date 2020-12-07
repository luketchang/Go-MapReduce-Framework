package mapreduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgsValid(t *testing.T) {
	s := Server{}
	s.parseArgumentList([]string{
		"mr",
		"--mapper",
		"./mapper/mrm",
		"--reducer",
		"./reducer/mrr",
		"--config",
		"./test_files/config_test.cfg",
	})

	assert.Equal(t, s.mapper, "./mapper/mrm")
	assert.Equal(t, s.reducer, "./reducer/mrr")
}
