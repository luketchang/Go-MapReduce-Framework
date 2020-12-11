package mapreduce

import (
	"io/ioutil"
	"os"
	"path"
)

func Contains(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ClearDirContents(dirPath string) {
	dir, _ := ioutil.ReadDir(dirPath)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dirPath, d.Name()}...))
	}
}
