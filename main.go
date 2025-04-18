package main

import (
	"fmt"
	"log"
	"os"
)

// Launch the program and execute the appropriate code
func main() {
	orders := flag()
	extra = false

	switch orders {
	case "-h", "--help":
		help()
	case "-r", "--run":
		active = 0
		serialize()
		for _, element := range variations {
			engine(element)
		}
		if active > 0 {
			tracking("Pushing to repository")
			push()
		}
	case "-v", "--version":
		version()
	case "--zero":
		alert("No flag detected -")
	default:
		alert("Unknown argument(s) -")
		help()
	}
}

// Enter a record to the log file
func journal(message string) {
	file, err := os.OpenFile("/data/automation/logs/bowerbird.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	fmt.Println(message)
}

// Print a colourized error message
func alert(message string) {
	fmt.Println("\n", bgred, message, halt, reset)
	fmt.Println(bgyellow, "Use -h for more detailed help information ")
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}

// Print program version number
func version() {
	fmt.Println("\n", yellow+"Bowerbird", green+bv, reset)
}

// Print help information for using the program
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag]")
	fmt.Println(yellow, "\nOperational Flags:")
	fmt.Println(green, " -h, --help", reset, "      Help information")
	fmt.Println(green, " -r, --run", reset, "       Run Program")
	fmt.Println(green, " -v, --version", reset, "   Display program version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("   bowerbird -r")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/bowerbird.git")
	fmt.Println(reset)
}
