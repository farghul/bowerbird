package main

import (
	"fmt"
	"os"
)

// Launch the program and execute as directed by the supplied flag
func main() {
	orders := flag()
	extra = false
	logo()

	switch orders {
	case "-h", "--help":
		help()
	case "-r", "--run":
		credits()
		active = 0
		serialize()
		for _, element := range brands {
			engine(element)
		}
		if active > 0 {
			inform("Pushing to repository")
			push()
		}
	case "-v", "--version":
	case "--zero":
		alert("No flag detected - ")
	default:
		alert("Unknown argument(s) - ")
	}
}

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -h to display help information")
	os.Exit(0)
}

// Print help information for using the program
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	Yellow.Println("\nOperational Flags:")
	Green.Printf("%s", "  -h, --help")
	fmt.Println("		Help Information")
	Green.Printf("%s", "  -r, --run")
	fmt.Println("		Run Program")
	Green.Printf("%s", "  -v, --version")
	fmt.Println("		Display Program Version")
	Yellow.Println("\nExample:")
	fmt.Println("  Adding your path to file if necessary, run:")
	Green.Printf("%s", "    bowerbird -r")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/bowerbird.git")
}

func logo() {
	Orange.Println("▗▄▄▖  ▗▄▖ ▗▖ ▗▖▗▄▄▄▖▗▄▄▖ ▗▄▄▖ ▗▄▄▄▖▗▄▄▖ ▗▄▄▄  ")
	Orange.Println("▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌▐▌ ▐▌  █  ▐▌ ▐▌▐▌  █ ")
	Orange.Println("▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌▐▛▀▀▘▐▛▀▚▖▐▛▀▚▖  █  ▐▛▀▚▖▐▌  █ ")
	Orange.Println("▐▙▄▞▘▝▚▄▞▘▐▙█▟▌▐▙▄▄▖▐▌ ▐▌▐▙▄▞▘▗▄█▄▖▐▌ ▐▌▐▙▄▄▀ ")
	Orange.Println(bv)
}

func credits() {
	fmt.Println("\nAn install tool for WordPress plugins")
	fmt.Println("Created by Byron Stuike")
}
