package main

import (
	"flag"
	"fmt"

	"github.com/farghul/bowerbird/internal"
	"github.com/farghul/bowerbird/workers"
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
			internal.Active = 0
			workers.Serialize()

			for _, element := range internal.Brands {
				workers.Engine(element)
			}
			if internal.Active > 0 {
				internal.Inform("Pushing to repository")
				workers.Push()
			}
		case *versionshort, *versionlong:
			internal.Orange.Println(internal.BV)
		}
	} else {
		internal.Alert("No arguments found - ")
	}
}

// Print help information for using the program
func help() {
	internal.Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	internal.Yellow.Println("\n Flags:")
	internal.Green.Printf("%s", "  -h, -help")
	fmt.Println("	   Help Information")
	internal.Green.Printf("%s", "  -u, -update")
	fmt.Println("	   Run main program")
	internal.Green.Printf("%s", "  -v, -version")
	fmt.Println("	   Display Program Version")
	internal.Yellow.Println("\nRun:")
	fmt.Println("  From the folder containing your compiled executable, run:")
	internal.Green.Printf("%s", "    bowerbird -u")
	internal.Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	internal.Green.Println("    https://github.com/farghul/bowerbird.git")
}

func logo() {
	internal.Orange.Println("▗▄▄▖  ▗▄▖ ▗▖ ▗▖▗▄▄▄▖▗▄▄▖ ▗▄▄▖ ▗▄▄▄▖▗▄▄▖ ▗▄▄▄  ")
	internal.Orange.Println("▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌▐▌ ▐▌  █  ▐▌ ▐▌▐▌  █ ")
	internal.Orange.Println("▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌▐▛▀▀▘▐▛▀▚▖▐▛▀▚▖  █  ▐▛▀▚▖▐▌  █ ")
	internal.Orange.Println("▐▙▄▞▘▝▚▄▞▘▐▙█▟▌▐▙▄▄▖▐▌ ▐▌▐▙▄▞▘▗▄█▄▖▐▌ ▐▌▐▙▄▄▀ ")
	internal.Orange.Println(internal.BV)
}

func credits() {
	fmt.Println("\nAn install tool for WordPress plugins")
	fmt.Println("Created by Byron Stuike")
}
