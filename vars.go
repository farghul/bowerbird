package main

// Definitions holds the path to the WordPress repository
type Definitions struct {
	WordPress string `json:"wordpress"`
}

// Jira builds a list of Jira API addresses
type Jira struct {
	Review string `json:"review"`
	Search string `json:"search"`
	Token  string `json:"token"`
	ToDo   string `json:"todo"`
	Basic  string `json:"basic"`
	URL    string `json:"url"`
}

// JQL holds the extracted data from the JQL queries
type JQL struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

type Color string

const (
	Reset           = "\033[0m"
	Black    Color  = "\033[30m"
	Red      Color  = "\033[31m"
	Green    Color  = "\033[32m"
	Yellow   Color  = "\033[33m"
	Blue     Color  = "\033[34m"
	Magenta  Color  = "\033[35m"
	Cyan     Color  = "\033[36m"
	White    Color  = "\033[37m"
	BGRed    Color  = "\033[41m"
	BGYellow Color  = "\033[43m"
	Orange   Color  = "\033[38;5;214m"
	bv       string = "1.0.0"
	halt     string = "program halted "
	meta     string = "/data/automation/jsons/"
	temp     string = "/data/automation/temp/"
)

var (
	active int
	query  JQL
	jira   Jira
	extra  bool
	plugin string
	ticket string
	defs   Definitions
	brands = []string{"freemius", "premium", "roots", "wpackagist", "wpengine"}
	jsons  = []string{meta + "definitions.json", meta + "jira.json"}
)
