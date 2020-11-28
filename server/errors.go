package server

import "errors"

var (
	errBadArgs = errors.New("bad arguments")
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
