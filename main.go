package main

import (
	"flag"
	"fmt"
	"os"
)

// Launch the program and execute as directed by the supplied flag
func main() {
	helpshort := flag.Bool("h", false, "Display help information")
	helplong := flag.Bool("help", false, "Display help information")
	updateshort := flag.Bool("u", false, "Run main program")
	updatelong := flag.Bool("update", false, "Run main program")
	versionshort := flag.Bool("v", false, "Display Program Version")
	versionlong := flag.Bool("version", false, "Display Program Version")
	flag.Parse()

	if flag.NFlag() > 0 {
		switch {
		case *helpshort, *helplong:
			help()
		case *updateshort, *updatelong:
			logo()
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
		case *versionshort, *versionlong:
			Orange.Println(bv)
		}
	} else {
		alert("No arguments found - ")
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
	Yellow.Println("\n Flags:")
	Green.Printf("%s", "  -h, -help")
	fmt.Println("	   Help Information")
	Green.Printf("%s", "  -u, -update")
	fmt.Println("	   Run main program")
	Green.Printf("%s", "  -v, -version")
	fmt.Println("	   Display Program Version")
	Yellow.Println("\nRun:")
	fmt.Println("  From the folder containing your compiled executable, run:")
	Green.Printf("%s", "    bowerbird -u")
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
