package mapreduce

import (
	"fmt"
	"os"
	"path/filepath"
)

func Contains(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ClearDirContents(dirPath string) error {
	files, err := filepath.Glob(filepath.Join(dirPath, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func PrintStageMessage(message string) {
	fmt.Println("")
	fmt.Printf("___________________%s___________________", message)
	fmt.Println("")
	fmt.Println("")
}
