package presenter

import "fmt"

const (
 	colorRed = "\033[31m"
 	colorGreen = "\033[32m"
 	fileTragedyKind = "files"
 	keywordTragedyKind = "keywords"
)

func PresentTragedyAvoiderResults(tragedies [][]string){
	presentTragediesCheck(tragedies[0], fileTragedyKind)
	presentTragediesCheck(tragedies[1], keywordTragedyKind)
}

func presentTragediesCheck(tragedies []string, tragedyKind string) {
	if len(tragedies) == 0 {
		presentSafeFromTragedyResult(tragedyKind)
	} else {
		presentTragedyResult(tragedies, tragedyKind)
	}
}

func presentTragedyResult(tragedies []string, tragedyKind string) {
	fmt.Println(colorRed, "===========================================================================================")
	fmt.Println(colorRed, fmt.Sprintf("|| These files matched the tragedy %s list. Check them out!!!", tragedyKind))
	fmt.Println(colorRed, "|| ----------------------------------------------------------------------------------------")
	for _, tragedy := range tragedies {
		fmt.Println(colorRed, fmt.Sprintf("|| %s", tragedy))
	}
	fmt.Println(colorRed, "===========================================================================================")
}

func presentSafeFromTragedyResult(tragedyKind string){
	fmt.Println(colorGreen, "===========================================================================================")
	fmt.Println(colorGreen, fmt.Sprintf("|| We couldn't find any %s tragedy. Congrats!!!", tragedyKind))
	fmt.Println(colorGreen, "|| We suggest you check your Tragedyfile for more accurate tragedy checks, just in case :D")
	fmt.Println(colorGreen, "===========================================================================================")
}