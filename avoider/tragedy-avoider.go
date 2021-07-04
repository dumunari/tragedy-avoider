package main

import (
	"os"
	"tragedy-avoider/avoider/checker"
	"tragedy-avoider/avoider/file-reader"
	"tragedy-avoider/avoider/presenter"
)

func main() {
	tragedyFileContent := file_reader.ReadTragedyfile()
	tragedies := checker.ExecuteTragedyChecks(tragedyFileContent)
	os.Exit(presenter.PresentTragedyAvoiderResults(tragedies))
}