package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type HookEvent struct {
	Prompt string `json:"prompt"`
}

//what gets called when there is a hook.
//prints prompt to obsidian file in desired format
func hookHandler() {
	var event HookEvent

	error := json.NewDecoder(os.Stdin).Decode(&event)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}

	file, err := os.OpenFile(
		"C:/projects/Prompt-Tracker/TestFile.md",
		os.O_APPEND|os.O_WRONLY,
		0640,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file,
		"This is a test print!\n %s",
		event.Prompt + "\n",
	)

	fmt.Println("{}") //tells the hook its done?
}