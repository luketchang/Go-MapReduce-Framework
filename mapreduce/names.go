package mapreduce

import "strings"

func ChangeExtension(fileName string, newExt string) string {
	lastDotIndex := strings.LastIndex(fileName, ".")
	newName := fileName[:lastDotIndex+1] + newExt
	return newName
}
