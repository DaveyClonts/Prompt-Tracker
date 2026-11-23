package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const LANG = ".go" //the file type being used by your project, used to enforce ledger files

//Im going to make this specific for my use. I can rewrite these to make them more flexible if wanted
//Checks if ledger files are up to date, and if not updates them
//TODO: make it rename when file name changes
func updateLedgerFiles() {
	vaultDir := filepath.Join(DIR_PATH, "Test")

	vaultFileList, error := os.ReadDir(vaultDir)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	targetFileList, error := os.ReadDir(DIR_PATH)
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

func createModificationRecord(event HookEvent) {
	//ledger file gets a link to a modificaiton
	//a modification gets the info

	//create a modification file
	//timestamp and prompt in that modfication file
	//create a snapshot (tree ID) of project before AI changes
	//AI modifies
	//take second snapshot
	//run a diff against second snapshot to original snapshot
	//store diff in modification file
	//link in ledger to modificaiton file
	//assess the situation


	//TODO: check if this mod already exists 
	modName := "Mod: " + time.Now().Format("2004-04-30 15:04:05")
	modFile, err := os.OpenFile(
		filepath.Join(DIR_PATH, "Test", modName),
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
		event.Prompt,
	)


	//TODO: DIFF LOGIC

	//TODO: figure out specifics on which indvidual files were modified so that diff can be stored

	//TODO: need to redo some code so that this info gets stored
	//prob move out....
	if (event.HookEventName == "UserPromptSubmit") {
		beforeSnapshot, err := takeSnapshot(DIR_PATH)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}



	if (event.HookEventName == "Stop") {
		afterSnapshot, err := takeSnapshot(DIR_PATH)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func trackChanges() {
	
}
