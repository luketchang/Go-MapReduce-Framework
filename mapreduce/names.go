package mapreduce

import (
	"fmt"
	"strconv"
	"strings"
)

const BaseWidth int = 5

func ChangeExtension(fileName string, newExt string) string {
	lastDotIndex := strings.LastIndex(fileName, ".")
	newName := fileName[:lastDotIndex+1] + newExt
	return newName
}

func GetPaddedNumber(num int) string {
	formattingString := "%0" + strconv.Itoa(BaseWidth) + "d"
	return fmt.Sprintf(formattingString, num)
}
