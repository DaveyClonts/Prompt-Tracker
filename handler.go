package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const DIR_PATH = `C:/projects/Prompt-Tracker`

type HookEvent struct {
	HookEventName	string `json:"hook_event_name"`
	Prompt 			string `json:"prompt"`
	Answer 			string `json:"last_assistant_message"`
}

//when code is called:
//creates dir for data if none exists
	//snapshots and vault will go there
//check if list of files is up to date
//if not update them
//creates ledger files (one to one map with files in targeted project)
//create modification file
//modification file has prompt + answer + diff(wait on diff)
//modification file has link to file modified then diff, repeated for its blast radius (Maybe wait on diff)

//what gets called when there is a hook
func hookHandler() {

	//check if directory exists
	_, err := os.Stat(filepath.Join(DIR_PATH, "tracker"))
	if os.IsNotExist(err) {
		//if not create it
		vaultDir, err := os.Mkdir("tracker", 0755)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	} else if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var event HookEvent
	error := json.NewDecoder(os.Stdin).Decode(&event)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	//checks and updates ledger files
	updateLedgerFiles()

	//TODO: make this only run when you have both snapshots
	createModificationRecord(event)

	//TEMP FOR HOOK TEST
	// file, error := os.OpenFile(
	// 	"C:/projects/Prompt-Tracker/TestFile.md",
	// 	os.O_CREATE|os.O_APPEND|os.O_WRONLY,
	// 	0640,
	// )
	// if error != nil {
	// 	fmt.Fprintln(os.Stderr, error)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// _, error = fmt.Fprintf(file,
	// 	"\n%s\n%s\n\n",
	// 	time.Now(),
	// 	event.Prompt,
	// )
	// if error != nil {
	// 	fmt.Fprintln(os.Stderr, error)
	// 	os.Exit(1)
	// }
	//TEMP END

	fmt.Println("{}") //tells the hook its done?
}
