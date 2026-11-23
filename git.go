package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runGit(path string, env []string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	
	cmd.Env = append(
		os.Environ(),
		env...
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(
			"git %v: %w: %s",
			args, err, output,
		)
	}

	return strings.TrimSpace(string(output)), nil //trim because git makes newlines
}

func takeSnapshot(repoPath string) (string, error) {

	//create a new path in /tmp/ to store staging area for changes (Index)
	tempIndexDir, err := os.MkdirTemp("", "git-snapshot-*") //MkdirTemp replaces * with random string
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempIndexDir)

	indexPath := filepath.Join(tempIndexDir, "index")
	envVars := []string{
		"GIT_INDEX_FILE=" + indexPath,
	}

	//clear index
	if _, err := runGit(repoPath, envVars, "read-tree", "--empty"); err != nil {
		return "", err
	}

	//add all from this file to index
	if _, err := runGit(repoPath, envVars, "add", "-A"); err != nil {
		return "", err
	}

	//create a tree from index
	treeID, err := runGit(repoPath, envVars, "write-tree")
	if err != nil {
		return "", err
	}

	return treeID, nil
}