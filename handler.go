package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type HookEvent struct {
	Prompt string `json:"prompt"`
	Answer string `json:last_assistant_message`
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
		"%s\n%s\n\n%s",
		time.Now(),
		event.Prompt,
		event.Answer,
	)

	fmt.Println("{}") //tells the hook its done?
}