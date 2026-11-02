package main

import (
	"fmt"
	"os"
	"strings"
)

const LANG = ".go" //the file type being used by your project
const PROJECT_FOLDER = `C:/projects/Prompt-Tracker`

//Im going to make this specific for my use. I can rewrite these to make them more flexible if wanted
func checkLedgerFiles() {
	vaultFileList, error:= os.ReadDir(PROJECT_FOLDER + "Test")
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	projectFileList, error:= os.ReadDir(PROJECT_FOLDER)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	currentLedgerFiles := []os.DirEntry{}

	for _, file := range projectFileList {
		if (strings.Contains(file.Name(), LANG)) {
			currentLedgerFiles = append(currentLedgerFiles, file)
		}
	}

	fileExists := []bool{}

	//TODO: make missing files?
	for i := 0; i < len(vaultFileList); i++ {
		for j := 0; j < len(currentLedgerFiles); j++ {
			if currentLedgerFiles[i].Name() == vaultFileList[j].Name() {
				//found file
				fileExists[i] = true
			}
		}
	}
}