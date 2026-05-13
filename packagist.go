package main

import "fmt"

// Execute the main set of program functions
func engine(element string) {
	result := compiler(element)
	if len(result) > 0 {
		active++
		rightplace()
		prepare()
		packagist(result)
	} else {
		fmt.Println("No " + element + " update tickets to process.")
	}
}

// Switch to the development branch, and pull any changes
func prepare() {
	execute("git", []string{"checkout", "development"}, ExecOptions{Stream: true})
	execute("git", []string{"pull"}, ExecOptions{Stream: true})
}

// A sequential list of tasks run to complete the program
func packagist(habitat []string) {
	if !extra {
		inform("Updating Composer")
		execute("composer", []string{"update", "--no-install"}, ExecOptions{Stream: true})
		extra = true
	}
	inform("Installing updates & commiting changes")
	sift(habitat)
}

// Iterate through the Args array and assign plugin and ticket values
func sift(box []string) {
	for i := 0; i < len(box); i++ {
		plugin = box[i]
		i++
		ticket = box[i]
		require()
		commit()
	}
}

// Run the appropriate composer require command
func require() {
	if edge() {
		execute("composer", []string{"require", plugin, "-W", "--no-install"}, ExecOptions{Stream: true})
	} else {
		execute("composer", []string{"require", plugin, "--no-install"}, ExecOptions{Stream: true})
	}
}

// Add and commit the update
func commit() {
	execute("git", []string{"add", "."}, ExecOptions{Stream: true})
	execute("git", []string{"commit", "-m", ticket, "-m", "Install " + plugin}, ExecOptions{Stream: true})
}

// Push modified content to the git repository
func push() {
	execute("git", []string{"push"}, ExecOptions{Stream: true})
}
