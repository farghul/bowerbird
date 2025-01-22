package main

// BitBucket builds a list of BitBucket tokens and api addresses
type BitBucket struct {
	Email     string `json:"email"`
	Reviewer1 string `json:"reviewer1"`
	Reviewer2 string `json:"reviewer2"`
	Token     string `json:"token"`
	URL       string `json:"url"`
	UUID      string `json:"uuid"`
}

// Jira builds a list of jira tokens and api addresses
type Jira struct {
	Token string `json:"token"`
	ToDo  string `json:"todo"`
	URL   string `json:"url"`
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

// Nouns builds a list of jira tokens and api addresses
type Nouns struct {
	WordPress string `json:"wordpress"`
}

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
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
	zero     string = "Not enough arguments supplied -"
	assets   string = "/data/scripts/automation/assets/"
	repos    string = "/data/scripts/automation/bitbucket/"
)

var (
	query     JQL
	jira      Jira
	ppt       Nouns
	plugin    string
	ticket    string
	bitbucket BitBucket
	jsons     = []string{assets + "jsons/bitbucket.json", assets + "jsons/jira.json", assets + "jsons/nouns.json"}
)
