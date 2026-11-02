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

	//check vault if ledger files is up to date
	


	file, error:= os.OpenFile(
		"C:/projects/Prompt-Tracker/TestFile.md",
		os.O_APPEND|os.O_WRONLY,
		0640,
	)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}
	defer file.Close()

	

	fmt.Fprintf(file,
		"%s\n%s\n\n%s",
		time.Now(),
		event.Prompt,
		event.Answer,
	)

	fmt.Println("{}") //tells the hook its done?
}