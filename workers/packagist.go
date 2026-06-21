package workers

import (
	"fmt"

	"github.com/farghul/bowerbird/internal"
)

// Execute the main set of program functions
func Engine(element string) {
	result := compiler(element)
	if len(result) > 0 {
		internal.Active++
		rightplace()
		prepare()
		packagist(result)
	} else {
		fmt.Println("\nNo " + element + " update tickets to process.")
	}
}

// Switch to the development branch, and pull any changes
func prepare() {
	internal.Execute("git", []string{"checkout", "development"}, internal.ExecOptions{Stream: true})
	internal.Execute("git", []string{"pull"}, internal.ExecOptions{Stream: true})
}

// A sequential list of tasks run to complete the program
func packagist(habitat []string) {
	if !internal.Extra {
		internal.Inform("Updating Composer")
		internal.Execute("composer", []string{"update", "--no-install"}, internal.ExecOptions{Stream: true})
		internal.Extra = true
	}
	internal.Inform("Installing updates & commiting changes")
	sift(habitat)
}

// Iterate through the Args array and assign plugin and ticket values
func sift(box []string) {
	for i := 0; i < len(box); i++ {
		internal.Plugin = box[i]
		i++
		internal.Ticket = box[i]
		require()
		commit()
	}
}

// Run the appropriate composer require command
func require() {
	if edge() {
		internal.Execute("composer", []string{"require", internal.Plugin, "-W", "--no-install"}, internal.ExecOptions{Stream: true})
	} else {
		internal.Execute("composer", []string{"require", internal.Plugin, "--no-install"}, internal.ExecOptions{Stream: true})
	}
}

// Add and commit the update
func commit() {
	internal.Execute("git", []string{"add", "."}, internal.ExecOptions{Stream: true})
	internal.Execute("git", []string{"commit", "-m", internal.Ticket, "-m", "Install " + internal.Plugin}, internal.ExecOptions{Stream: true})
}

// Push modified content to the git repository
func Push() {
	internal.Execute("git", []string{"push"}, internal.ExecOptions{Stream: true})
}
