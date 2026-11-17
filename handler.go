package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type HookEvent struct {
	Prompt string `json:"prompt"`
	Answer string `json:"last_assistant_message"`
}

//when code is called:
//check if list of files is up to date
//if not update them
//creates ledger files (one to one map with files in targeted project)
//create modification file
//modification file has prompt + answer + diff(wait on diff)
//modification file has link to file modified then diff, repeated for its blast radius (Maybe wait on diff)

//what gets called when there is a hook
func hookHandler() {
	var event HookEvent

	error := json.NewDecoder(os.Stdin).Decode(&event)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	//checks and updates ledger files
	updateLedgerFiles()

	createModificationRecord(event.Prompt)

	file, error := os.OpenFile(
		"C:/projects/Prompt-Tracker/TestFile.md",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0640,
	)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}
	defer file.Close()

	_, error = fmt.Fprintf(file,
		"\n%s\n%s\n\n",
		time.Now(),
		event.Prompt,
	)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	fmt.Println("{}") //tells the hook its done?
}
