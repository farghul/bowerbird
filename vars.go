package main

// Definitions holds the path to the WordPress repository
type Definitions struct {
	WordPress string `json:"wordpress"`
}

// Jira builds a list of jira api addresses
type Jira struct {
	Review string `json:"review"`
	ToDo   string `json:"todo"`
	URL    string `json:"url"`
}

// JQL holds the extracted data from the JQL queries
type JQL struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
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

// Tokens builds a list of jira tokens and api addresses
type Tokens struct {
	Jira string `json:"jira"`
}

const (
	bv       string = "1.0.0"
	branch   string = "update/"
	reset    string = "\033[0m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgred    string = "\033[41m"
	bgyellow string = "\033[43m"
	halt     string = "program halted "
	tokens   string = "/data/automation/tokens/"
	repo     string = "/data/automation/bitbucket/desso-automation-conf/"
)

var (
	active      int
	query       JQL
	jira        Jira
	extra       bool
	plugin      string
	ticket      string
	token       Tokens
	definitions Definitions
	brands      = []string{"freemius", "premium", "roots", "wpackagist", "wpengine"}
	jsons       = []string{repo + "jsons/definitions.json", repo + "jsons/jira.json", tokens + "tokens.json"}
)
