package main

func main() {
	tragedyFileContent := readTragedyfile()
	tragedies := executeTragedyChecks(tragedyFileContent)
	presentTragedyAvoiderResults(tragedies)
}