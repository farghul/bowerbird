package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type ExecOptions struct {
	Stream bool
	Env    []string
	Dir    string
}

func Execute(task string, args []string, opts ExecOptions) ([]byte, error) {
	cmd := exec.Command(task, args...)
	cmd.Env = append(os.Environ(), opts.Env...)
	cmd.Dir = opts.Dir

	if opts.Stream {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return nil, cmd.Run()
	}

	return cmd.CombinedOutput()
}

// Check for errors, print the result if found
func Inspect(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Provide and highlight an informational message
func Inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func Alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(Halt)
	Inform("Use -h to display help information")
	os.Exit(0)
}

// Empty the contents a folder
func Clearout(path string) {
	list := ls(path)
	for _, file := range list {
		sweep(path + file)
	}
}

// Remove files or directories
func sweep(cut ...string) {
	Inspect(os.RemoveAll(cut[0.]))
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	Inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	Inspect(err)
	return outcome
}

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}
