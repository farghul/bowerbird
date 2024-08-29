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

func driver() {
	data, err := os.ReadFile("secrets/jira.json")
	inspect(err)
	json.Unmarshal(data, &jira)

	search := api(jira.Search)
	json.Unmarshal(search, &desso)
}

func compiler(element string) []string {
	var candidate []string
	for i := 0; i < len(desso.Issues); i++ {
		if strings.Contains(desso.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, desso.Issues[i].Fields.Summary)
			candidate = append(candidate, desso.Issues[i].Key)
		}
		// if strings.Contains(desso.Issues[i].Fields.Summary, "wpackagist") {
		// 	free = append(free, desso.Issues[i].Fields.Summary)
		// 	free = append(free, desso.Issues[i].Key)
		// }

		// if strings.Contains(desso.Issues[i].Fields.Summary, "premium") {
		// 	paid = append(paid, desso.Issues[i].Fields.Summary)
		// 	paid = append(paid, desso.Issues[i].Key)
		// }

		// if strings.Contains(desso.Issues[i].Fields.Summary, "bcgov") {
		// 	dev = append(dev, desso.Issues[i].Fields.Summary)
		// 	dev = append(dev, desso.Issues[i].Key)
		// }
	}
	return candidate
}

// Grab the ticket information from Jira in order to extract the DESSO-XXXX identifier
func api(criteria string) []byte {
	result := execute("-c", "curl", "--request", "GET", "--url", jira.Base+criteria, "--header", "Authorization: Basic "+jira.Token, "--header", "Accept: application/json")
	return result
}

// Confirm the current working directory is correct
func changedir() {
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

// Record a message to the log file
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

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Run standard terminal commands and display the output
// func execute(task string, args ...string) {
// 	osCmd := exec.Command(task, args...)
// 	osCmd.Stdout = os.Stdout
// 	osCmd.Stderr = os.Stderr
// 	err := osCmd.Run()
// 	inspect(err)
// }

// Run a terminal command, then capture and return the output as a byte
// func capture(task string, args ...string) []byte {
// 	lpath, err := exec.LookPath(task)
// 	inspect(err)
// 	osCmd, _ := exec.Command(lpath, args...).CombinedOutput()
// 	return osCmd
// }

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Test for the minimum number of arguments
// func verify() string {
// 	var f string
// 	if inputs < 2 {
// 		f = "--zero"
// 	} else {
// 		f = passed[1]
// 	}
// 	return f
// }

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

// WIP - Dynamically update the require field in the composer.json file
// func monitor() {
// grep := capture("grep Version readme.txt | grep " + number[1] + " readme.txt | grep higher")
// grep := capture("grep", "newer", "readme.txt")
// fmt.Println(bytes.IndexRune(grep, '*'))
// requires := strings.Split(string(grep), " ")
// evtp.Require.EventsCalendar = `>` + strings.Trim(requires[4], "\n")
// }

// Print a colourized error message
func alert(message string) {
	fmt.Println(red, message, halt, reset)
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}

// Print program version number
func version() {
	fmt.Println(yellow+"Bowerbird", green+bv, reset)
}

// Print help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag] [vendor/plugin]:[version] [ticket#]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -v, --version", reset, "	Display Program Version")
	fmt.Println(green, " -s, --subscription", reset, "	Subscription Plugin Update")
	fmt.Println(green, " -r, --release", reset, "	Production Release Plugin Update")
	fmt.Println(green, " -p, --packaged", reset, "	Satis & WPackagist Plugin Update")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println(green, "   bowerbird -p wpackagist-plugin/mailpoet:4.6.1 821")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/bowerbird.git")
	fmt.Println(reset)
}
