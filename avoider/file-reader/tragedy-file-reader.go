package file_reader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

const (
	tragedyfileEnvVar   = "TRAGEDYFILE"
	userHomeEnvVar      = "HOME"
	colorBlue           = "\033[34m"
	tragedyfileFilename = "/Tragedyfile.yaml"
)

type Tragedyfile struct {
	Files []string `yaml:",flow"`
	Keywords []string `yaml:",flow"`
}

func ReadTragedyfile() *Tragedyfile {
	tragedyEnvVar := os.Getenv(tragedyfileEnvVar)
	userHomeDir := os.Getenv(userHomeEnvVar)
	var tragedyFileLocation string

	if tragedyEnvVar != "" {
		tragedyFileLocation = tragedyEnvVar
		directoryOrFileExists(tragedyFileLocation)
		fmt.Println(colorBlue, "===========================================================================================")
		fmt.Println(colorBlue, fmt.Sprintf("|| Using local Tragedyfile: %s", tragedyEnvVar))
		fmt.Println(colorBlue, "===========================================================================================")
		fmt.Println()
	} else {
		tragedyFileLocation = userHomeDir + tragedyfileFilename
		directoryOrFileExists(tragedyFileLocation)
		fmt.Println(colorBlue, "===========================================================================================")
		fmt.Println(colorBlue, fmt.Sprintf("|| Using global Tragedyfile: %s/%s", userHomeDir, "Tragedyfile.yaml"))
		fmt.Println(colorBlue, "===========================================================================================")
		fmt.Println()
	}

	var tragedyFileContent Tragedyfile
	return tragedyFileContent.getConf(tragedyFileLocation)
}

func directoryOrFileExists(directoryOrFileName string) {
	fileInfo, err := os.Stat(directoryOrFileName)
	if os.IsNotExist(err) || fileInfo.Size() == 1 || fileInfo.Size() == 0  {
		log.Fatalf("Tragedyfile %s could not be found. Please, check file existence and/or PATH correctness.", directoryOrFileName)
	}
}

func (c *Tragedyfile) getConf(tragedyfile string) *Tragedyfile {
	yamlFile, err := ioutil.ReadFile(tragedyfile)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error reading the Tragedyfile: %s ", tragedyfile))
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Tragedyfile.yaml unmarshal error: %v", err))
	}
	return c
}
