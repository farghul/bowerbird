package main

// A sequential list of tasks run to complete the program
func packagist(flavour []string) {
	tracking("Updating Composer")
	execute("-e", "composer", "update", "--no-install")
	tracking("Installing updates & commiting changes")
	sift(flavour)
	tracking("Writing to log file")
	journal(ticket + " " + plugin)
	tracking("Pushing to repository")
	push()
}

// Run the appropriate composer require command based on the flag value
func require() {
	if edge() {
		execute("-e", "composer", "require", plugin, "-W", "--no-install")
	} else {
		execute("-e", "composer", "require", plugin, "--no-install")
	}
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

// Add and commit the update
func commit() {
	execute("-e", "git", "add", ".")
	execute("-e", "git", "commit", "-m", ticket+" install "+plugin)
}

// Push modified content to the git repository
func push() {
	execute("-e", "git", "push")
}
