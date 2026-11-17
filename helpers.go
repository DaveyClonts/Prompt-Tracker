package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const LANG = ".go" //the file type being used by your project, used to enforce ledger files
const PROJECT_FOLDER = `C:/projects/Prompt-Tracker`

//Im going to make this specific for my use. I can rewrite these to make them more flexible if wanted
//Checks if ledger files are up to date, and if not updates them
func updateLedgerFiles() {
	vaultDir := filepath.Join(PROJECT_FOLDER, "Test")

	vaultFileList, error := os.ReadDir(vaultDir)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	targetFileList, error := os.ReadDir(PROJECT_FOLDER)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	fileExists := make([]bool, len(targetFileList))
	for i := 0; i < len(targetFileList); i++ {
		if targetFileList[i].IsDir() || filepath.Ext(targetFileList[i].Name()) != LANG {
			continue
		}
		for j := 0; j < len(vaultFileList); j++ {
			if targetFileList[i].Name() + ".md" == vaultFileList[j].Name() {
				fileExists[i] = true
			}
		}

		//if not found make record in vault
		if fileExists[i] == false {
			file, err := os.OpenFile(filepath.Join(vaultDir, targetFileList[i].Name() + ".md"), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
			if os.IsExist(err) {
				continue
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer file.Close()
		}
	}
}

func createModificationRecord(prompt string) {
	//ledger file gets a link to a modificaiton
	//a modification gets the info

	//create a modification file
	//timestamp and prompt in that modfication file
	//link in ledger to modificaiton file
	//assess the situation

	//TODO: rewrite to use create to name file
	modFile, err := os.OpenFile(
		filepath.Join(PROJECT_FOLDER, "Test"),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0640,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer modFile.Close()

	_, err = fmt.Fprintf(
		modFile,
		`------------------\n%s\n%s\n\`,
		time.Now(),
		prompt,
	)

	//find relevant ledger
	vaultFileList, error := os.ReadDir(vaultDir)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	} //reused code... TODO: find better solution

	//TODO: figure out how find which file was modified
	//might need to force the ai to "fill out a report"
	//wait im wrong just scrape git diff?
	//also could use structured events outputed by codex
	
		
}
