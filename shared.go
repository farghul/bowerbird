package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Test for an optional flag
func flag() string {
	var passed string

	if len(os.Args) == 1 {
		passed = "--zero"
	} else {
		passed = os.Args[1]
	}
	return passed
}

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	clearout("/data/automation/temp/")
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			err := json.Unmarshal(data, &definitions)
			inspect(err)
		case 1:
			err := json.Unmarshal(data, &jira)
			inspect(err)
		case 2:
			err := json.Unmarshal(data, &token)
			inspect(err)
		}
	}
}

// Compile the results of a Jira API query and save summary and key into a string slice
func compiler(element string) []string {
	if element == "premium" {
		err := json.Unmarshal(api(jira.Review), &query)
		inspect(err)
	} else {
		err := json.Unmarshal(api(jira.ToDo), &query)
		inspect(err)
	}

	var candidate []string

	for i := range query.Issues {
		if strings.Contains(query.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, query.Issues[i].Fields.Summary)
			candidate = append(candidate, query.Issues[i].Key)
		}
	}
	return candidate
}

// Search the Jira API
func api(criteria string) []byte {
	result := execute("-c", "curl", "--request", "GET", "--url", jira.URL+"search?jql="+criteria, "--header", "Authorization: Basic "+token.Jira, "--header", "Accept: application/json")
	return result
}

// Confirm the current working directory is correct
func rightplace() {
	err := os.Chdir(definitions.WordPress)
	inspect(err)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Check for edge cases which require the -W flag
func edge() bool {
	found := false
	if strings.Contains(plugin, "roots/wordpress") {
		found = true
	}
	return found
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-c":
		result, err := osCmd.Output()
		inspect(err)
		return result
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Empty the contents a folder
func clearout(path string) {
	list := ls(path)
	for _, file := range list {
		sweep(path + file)
	}
}

// Remove files or directories
func sweep(cut ...string) {
	inspect(os.Remove(cut[0.]))
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
}
