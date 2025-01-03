package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	route = os.Args
)

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			json.Unmarshal(data, &access)
		case 1:
			json.Unmarshal(data, &values)
		}
	}
}

func compiler(element string) []string {
	json.Unmarshal(api(access.ToDo), &jira)
	var candidate []string

	for i := 0; i < len(jira.Issues); i++ {
		if strings.Contains(jira.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, jira.Issues[i].Fields.Summary)
			candidate = append(candidate, jira.Issues[i].Key)
		}
	}
	return candidate
}

// Search the Jira API
func api(criteria string) []byte {
	result := execute("-c", "curl", "--request", "GET", "--url", access.Cloud+criteria, "--header", "Authorization: Basic "+access.Token, "--header", "Accept: application/json")
	return result
}

// Confirm the current working directory is correct
func rightplace() {
	os.Chdir(access.WordPress)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Switch to the desired branch, and pull any changes
func prepare() {
	var branch string
	if flag == "-p" {
		branch = "main"
	} else {
		branch = "development"
	}
	execute("-e", "git", "switch", branch)
	execute("-e", "git", "pull")
}

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0644))
}

// Enter a record to the log file
func journal(message string) {
	file, err := os.OpenFile(access.Programs+"logs/bowerbird.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		exec.Command(task, args...).CombinedOutput()
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

// Check to see if the current release branch already exists locally
func exists(prefix, tag string) bool {
	found := false
	b, _ := exec.Command("git", "branch").Output()
	if strings.Contains(string(b), prefix+tag) {
		found = true
	}
	return found
}

// Check for edge cases which require the -W flag
func edge() bool {
	found := false
	if strings.Contains(plugin, "roots/wordpress") {
		found = true
	}
	return found
}

// Print a colourized error message
func alert(message string) {
	fmt.Println("\n", bgred, message, halt, reset)
	fmt.Println(bgyellow, "Use -h for more detailed help information ")
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}

// Print program version number
func version() {
	fmt.Println("\n", yellow+"Bowerbird", green+bv, reset)
}

// Print help information for using the program
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -c, --core", reset, "      Install WordPress core updates")
	fmt.Println(green, " -d, --developer", reset, " Install internal developer updates")
	fmt.Println(green, " -h, --help", reset, "      Help information")
	fmt.Println(green, " -p, --premium", reset, "   Install paid subscription updates")
	fmt.Println(green, " -v, --version", reset, "   Display program version")
	fmt.Println(green, " -w, --wpackagist", reset, "Install free wpackagist updates")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("   bowerbird -w")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/bowerbird.git")
	fmt.Println(reset)
}
