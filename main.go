package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-c", "--core":
			director("roots")
		case "-h", "--help":
			help()
		case "-w", "--wpackagist":
			director("wpackagist")
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	} else {
		alert("No flag detected -")
	}
}

func director(variety string) {
	serialize()
	flavour := compiler(variety)
	if len(flavour) > 0 {
		rightplace()
		prepare()
		packagist(flavour)
	} else {
		journal("No update tickets to process.")
	}
}
