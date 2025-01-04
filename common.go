package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	reset     string = "\033[0m"
	green     string = "\033[32m"
	yellow    string = "\033[33m"
	red       string = "\033[41m"
	bv        string = "2.0"
	relbranch string = "release/"
	upbranch  string = "update/DESSO-"
	halt      string = "program halted "
	zero      string = "Not enough arguments supplied -"
	bitbucket string = "/Documents/bitbucket/"
)

var (
	ecp            ECP
	evtp           EVTP
	satis          Satis
	flag           = verify()
	reader         = bufio.NewReader(os.Stdin)
	hmdr, _        = os.UserHomeDir()
	passed         = os.Args
	inputs         = len(passed)
	release        string
	plugin         string
	ticket         string
	number, folder []string
)

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
	if flag == "-s" {
		branch = "main"
	} else {
		branch = "development"
	}
	execute("git", "switch", branch)
	execute("git", "pull")
}

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0644))
}

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Run standard terminal commands and display the output
func execute(task string, args ...string) {
	osCmd := exec.Command(task, args...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err := osCmd.Run()
	inspect(err)
}

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
func verify() string {
	var f string
	if inputs < 2 {
		f = "--zero"
	} else {
		f = passed[1]
	}
	return f
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

// WIP - Dynamically update the require field in the composer.json file
// func monitor() {
// 	// grep := capture("grep Version readme.txt | grep " + number[1] + " readme.txt | grep higher")
// 	grep := capture("grep", "newer", "readme.txt")
// 	// fmt.Println(bytes.IndexRune(grep, '*'))
// 	requires := strings.Split(string(grep), " ")
// 	evtp.Require.EventsCalendar = `>` + strings.Trim(requires[4], "\n")
// }

// Decide whether an update or release branch is needed, and make it so
func checkout(prefix string) {
	suffix := ""
	if flag == "-r" {
		suffix = release
	} else {
		suffix = ticket
	}

	if exists(prefix, suffix) {
		execute("git", "switch", prefix+suffix)
	} else {
		execute("git", "checkout", "-b", prefix+suffix)
	}
}

// Add and commit the update
func commit() {
	execute("git", "add", ".")
	execute("git", "commit", "-m", plugin+" (DESSO-"+ticket+")")
}

// Push modified content to the git repository
func push() {
	switch flag {
	case "-r":
		execute("git", "push", "--set-upstream", "origin", relbranch+release)
	case "-s":
		execute("git", "push", "--set-upstream", "origin", upbranch+ticket)
	default:
		execute("git", "push")
	}
}

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
