package main

// BitBucket builds a list of BitBucket tokens and api addresses
type BitBucket struct {
	Email string `json:"email"`
	Token string `json:"token"`
	URL   string `json:"url"`
	UUID  string `json:"uuid"`
}

// Jira builds a list of jira tokens and api addresses
type Jira struct {
	Review string `json:"review"`
	Token  string `json:"token"`
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
	assets   string = "/data/automation/assets/"
	repos    string = "/data/automation/bitbucket/"
)

var (
	active     int
	query      JQL
	jira       Jira
	extra      bool
	plugin     string
	ticket     string
	bitbucket  BitBucket
	ppt        map[string]string
	variations = []string{"freemius", "premium", "roots", "wpackagist", "wpengine"}
	jsons      = []string{assets + "bitbucket.json", assets + "jira.json", assets + "nouns.json"}
)
