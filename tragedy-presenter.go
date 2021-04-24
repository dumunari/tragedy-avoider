package main

import "fmt"

const(
 colorRed = "\033[31m"
)

func presentTragedyAvoiderResults(tragedies [][]string){
	fmt.Println(colorRed, "===========================================================================================")
	fmt.Println(colorRed, "|| These files matched the tragedy files list. Check them out!!!")
	fmt.Println(colorRed, "|| ----------------------------------------------------------------------------------------")
	for _, fileTragedy := range tragedies[0] {
		fmt.Println(colorRed, fmt.Sprintf("|| %s", fileTragedy))
	}
	fmt.Println(colorRed, "===========================================================================================")
	fmt.Println()
	fmt.Println(colorRed, "===========================================================================================")
	fmt.Println(colorRed, fmt.Sprintf("|| These files matched the tragedy keywords list. Check them out!!!"))
	fmt.Println(colorRed, "|| ----------------------------------------------------------------------------------------")
	for _, keywordTragedy := range tragedies[1] {
		fmt.Println(colorRed, fmt.Sprintf("|| %s", keywordTragedy))
	}

	fmt.Println(colorRed, "===========================================================================================")
}