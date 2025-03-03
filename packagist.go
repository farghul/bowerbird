package main

// Execute the main set of program functions
func engine(element string) {
	result := compiler(element)
	if len(result) > 0 {
		active++
		rightplace()
		prepare()
		packagist(result)
	} else {
		journal("No " + element + " update tickets to process.")
	}
}

// Switch to the development branch, and pull any changes
func prepare() {
	execute("-v", "git", "checkout", "development")
	execute("-v", "git", "pull")
}

// A sequential list of tasks run to complete the program
func packagist(flavour []string) {
	if !extra {
		tracking("Updating Composer")
		execute("-v", "composer", "update", "--no-install")
		extra = true
	}
	tracking("Installing updates & commiting changes")
	sift(flavour)
	tracking("Writing to log file")
	journal(ticket + " " + plugin)
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
		execute("-v", "composer", "require", plugin, "-W", "--no-install")
	} else {
		execute("-v", "composer", "require", plugin, "--no-install")
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
