package internal

// Definitions holds the path to the WordPress repository
type Definitions struct {
	Review    string `json:"review"`
	Search    string `json:"search"`
	Token     string `json:"token"`
	ToDo      string `json:"todo"`
	Basic     string `json:"basic"`
	URL       string `json:"url"`
	WordPress string `json:"wordpress"`
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
	Red      Color  = "\033[31m"
	Green    Color  = "\033[32m"
	Yellow   Color  = "\033[33m"
	BGRed    Color  = "\033[41m"
	BGYellow Color  = "\033[43m"
	Orange   Color  = "\033[38;5;214m"
	BV       string = "1.0.0"
	Halt     string = "program halted "
	Temp     string = "/data/automation/temp/"
	Meta     string = "/data/automation/jsons/bowerbird.json"
)

var (
	Active int
	Query  JQL
	Extra  bool
	Plugin string
	Ticket string
	Defs   Definitions
	Brands = []string{"freemius", "premium", "roots", "wpackagist", "wpengine"}
)
