package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)
const (
	TragedyDirEnvVar = "TRAGEDYDIR"
)

func executeTragedyChecks(tragedyfile *Tragedyfile) [][]string{
	var tragedies [][]string
	files := walkThroughDirAndListFiles(os.Getenv(TragedyDirEnvVar))
	tragedies = append(tragedies, examineFilesAndReturnPossibleTragedies(files, tragedyfile.Files))
	tragedies = append(tragedies, examineKeywordsAndReturnPossibleTragedies(files,tragedyfile.Keywords))
	return tragedies
}

func walkThroughDirAndListFiles(dirname string) [] string{
	var files []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(path, ".git") &&
			!strings.Contains(path, ".idea") &&
				!strings.Contains(path, ".jar")  &&
					!strings.Contains(path, ".class") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("An error occurred while walking through %s dir", dirname)
	}

	return files
}

func examineFilesAndReturnPossibleTragedies(files []string, tragedyFileFiles []string) []string{
	var possibleTragedyFiles []string
	for _, file := range files {
		for _, tragedyFile := range tragedyFileFiles {
			if strings.Contains(file, tragedyFile) {
				possibleTragedyFiles = append(possibleTragedyFiles, file)
			}
		}
	}
	return possibleTragedyFiles
}

func examineKeywordsAndReturnPossibleTragedies(files []string, tragedyFileKeywords []string) []string {
	var possibleTragedyKeywords []string
	for _, file := range files {
		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error opening file: %s", file)
		}
		fileLines := strings.Split(string(fileContent), "\n")

		for index, line := range fileLines {
			for _, tragedyKeyword := range tragedyFileKeywords {
				if strings.Contains(line, tragedyKeyword) {
					tragedyKeywordAndLine :=
						fmt.Sprintf("File %s matched keyword %s on line %d", file, tragedyKeyword, index + 1)
					possibleTragedyKeywords = append(possibleTragedyKeywords, tragedyKeywordAndLine)
				}
			}
		}
	}
	return possibleTragedyKeywords
}