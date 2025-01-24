package main

// Launch the program and execute the appropriate code
func main() {
	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-v", "--version":
		version()
	case "--zero":
		serialize()
		result := compiler("wpackagist")
		if len(result) > 0 {
			rightplace()
			prepare()
			packagist(result)
		} else {
			journal("No WPackagist update tickets to process.")
		}
	default:
		alert("Unknown flag -")
		help()
	}
}
