package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const LANG = ".go" //the file type being used by your project, used to enforce ledger files
const PROJECT_FOLDER = `C:/projects/Prompt-Tracker`

//Im going to make this specific for my use. I can rewrite these to make them more flexible if wanted
func updateLedgerFiles() {
	//TODO: reread and remove?
	vaultDir := filepath.Join(PROJECT_FOLDER, "Test")
	if err := os.MkdirAll(vaultDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
