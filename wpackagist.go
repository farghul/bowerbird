package main

// A sequential list of tasks run to complete the program
func wpackagist() {
	tracking("Composer Update")
	execute("-e", "composer", "update", "--no-install")
	tracking("Plugin Update")
	sift()
	tracking("Git Push")
	push()
}

// Add the previously tested plugins to the composer-prod.json file
func released() {
	release = solicit("Enter the current release number: ")
	checkout(relbranch)
	wpackagist()
}

// Run the appropriate composer require command based on the flag value
func require() {
	if flag == "-r" {
		if edge() {
			execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "-W", "--no-install")
		} else {
			execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "--no-install")
		}
	} else {
		if edge() {
			execute("-e", "composer", "require", plugin, "-W", "--no-install")
		} else {
			execute("-e", "composer", "require", plugin, "--no-install")
		}
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sift() {
	for i := 2; i < inputs; i++ {
		plugin = free[i]
		i++
		ticket = free[i]
		require()
		commit()
	}
}

// Decide whether an update or release branch is needed, and make it so
func checkout(prefix string) {
	suffix := ""
	if flag == "-r" {
		suffix = release
	} else {
		suffix = ticket
	}

	if exists(prefix, suffix) {
		execute("-e", "git", "switch", prefix+suffix)
	} else {
		execute("-e", "git", "checkout", "-b", prefix+suffix)
	}
}

// Add and commit the update
func commit() {
	execute("-e", "git", "add", ".")
	execute("-e", "git", "commit", "-m", plugin+" ("+ticket+")")
}

// Push modified content to the git repository
func push() {
	switch flag {
	case "-r":
		execute("-e", "git", "push", "--set-upstream", "origin", relbranch+release)
	case "-s":
		execute("-e", "git", "push", "--set-upstream", "origin", upbranch+ticket)
	default:
		execute("-e", "git", "push")
	}
}
