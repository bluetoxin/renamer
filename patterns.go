package main

import (
	"regexp"
	"strconv"
)

func RenameDefault(fileName string) string {
	re := regexp.MustCompile(`_(\d+)\.txt`)
	matches := re.FindStringSubmatch(fileName)
	if len(matches) != 2 {
		return fileName
	}

	fileIndex, _ := strconv.Atoi(matches[1])
	newName := "(" + strconv.Itoa(fileIndex) + " of 4).txt"
	return newName
}

func RenameCustom(fileName string) string {
	re := regexp.MustCompile(`(.*?)_(\d+)\.txt`)
	matches := re.FindStringSubmatch(fileName)
	if len(matches) != 3 {
		return fileName
	}

	prefix := matches[1]
	fileIndex, _ := strconv.Atoi(matches[2])
	newName := strconv.Itoa(fileIndex) + " - " + prefix + ".txt"
	return newName
}
