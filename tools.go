package main

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

func execute(task string, args []string, opts ExecOptions) ([]byte, error) {
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

// Run a terminal command using flags to customize the output
func executor(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-c":
		result, err := osCmd.Output()
		inspect(err)
		return result
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}

// Empty the contents a folder
func clearout(path string) {
	list := ls(path)
	for _, file := range list {
		sweep(path + file)
	}
}

// Remove files or directories
func sweep(cut ...string) {
	inspect(os.RemoveAll(cut[0.]))
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
}
