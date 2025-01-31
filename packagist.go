package main

// A sequential list of tasks run to complete the program
func packagist(flavour []string) {
	tracking("Updating Composer")
	execute("-v", "composer", "update", "--no-install")
	tracking("Installing updates & commiting changes")
	sift(flavour)
	tracking("Writing to log file")
	journal(ticket + " " + plugin)
	tracking("Pushing to repository")
	push()
}

// Switch to the development branch, and pull any changes
func prepare() {
	execute("-v", "git", "checkout", "development")
	execute("-v", "git", "pull")
}

// Run the appropriate composer require command based on the flag value
func require() {
	if edge() {
		execute("-v", "composer", "require", plugin, "-W", "--no-install")
	} else {
		execute("-v", "composer", "require", plugin, "--no-install")
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
	execute("-v", "git", "add", ".")
	execute("-v", "git", "commit", "-m", ticket, "-m", "Install "+plugin)
}

// Push modified content to the git repository
func push() {
	execute("-v", "git", "push")
}
