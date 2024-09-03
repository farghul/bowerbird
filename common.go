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
	route   = os.Args
	hmdr, _ = os.UserHomeDir()
)

func driver() {
	data, err := os.ReadFile("secrets/jira.json")
	inspect(err)
	json.Unmarshal(data, &access)

	search := api(access.Search)
	json.Unmarshal(search, &jira)
}

func compiler(element string) []string {
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
	result := execute("-c", "curl", "--request", "GET", "--url", access.Base+criteria, "--header", "Authorization: Basic "+access.Token, "--header", "Accept: application/json")
	return result
}

// Confirm the current working directory is correct
func rightplace() {
	os.Chdir(hmdr + bitbucket + "blog_gov_bc_ca")
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Switch to the desired branch, and pull any changes
func prepare() {
	tracking("Preparing Branch")
	var branch string
	if flag == "-s" && folder[1] == "events-virtual" {
		branch = "main"
	} else if flag == "-s" {
		branch = "master"
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
	file, err := os.OpenFile("logs/bowerbird.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	fmt.Println("  [program] [flag] [vendor/plugin]:[version] [ticket#]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -v, --version", reset, "	Display Program Version")
	fmt.Println(green, " -s, --subscription", reset, "	Subscription Plugin Update")
	fmt.Println(green, " -d, --developer", reset, "	Internal Developer Plugin Update")
	fmt.Println(green, " -p, --packaged", reset, "	Satis & WPackagist Plugin Update")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println(green, "   bowerbird -p wpackagist-plugin/mailpoet:4.6.1 821")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/bowerbird.git")
	fmt.Println(reset)
}
