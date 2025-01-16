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

// Credentials builds a username:password array
type Credentials []struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Downloads contains the direct download links
type Downloads struct {
	Polylang  string `json:"polylang-pro"`
	AllExport string `json:"wp-all-export-pro"`
	Gravity   string `json:"gravityforms"`
	SearchWP  string `json:"searchwp"`
	Calendar  string `json:"events-calendar-pro"`
	Tickets   string `json:"event-tickets-plus"`
	Virtual   string `json:"events-virtual"`
	Uji       string `json:"uji-countdown-premium"`
}

// ECP structure captures the contents of the composer.json file for Events Calendar Pro
type ECP struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		EventsCalendar string `json:"wpackagist-plugin/the-events-calendar"`
	} `json:"require"`
}

// EVTP structure captures the contents of the composer.json file for Events Tickets Plus
type EVTP struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		EventsCalendar string `json:"wpackagist-plugin/the-events-calendar"`
		EventsTicket   string `json:"wpackagist-plugin/event-tickets"`
	} `json:"require"`
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

// Logins holds the login URLs to access their accounts
type Logins struct {
	Polylang  string `json:"polylang-pro"`
	Allimport string `json:"wp-all-export-pro"`
	Gravity   string `json:"gravityforms"`
	SearchWP  string `json:"searchwp"`
}

// Nouns builds a list of jira tokens and api addresses
type Nouns struct {
	Core      string `json:"core"`
	Prem      string `json:"prem"`
	WordPress string `json:"wordpress"`
	WPac      string `json:"wpac"`
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
	programs string = "/data/scripts/automation/programs/"
	repos    string = "/data/scripts/automation/bitbucket/"
)

var (
	flag      string
	plugin    string
	ticket    string
	query     JQL
	ecp       ECP
	evtp      EVTP
	jira      Jira
	satis     Satis
	ppt       Nouns
	site      Logins
	bitbucket BitBucket
	download  Downloads
	cred      Credentials
	jsons     = []string{programs + "jsons/bitbucket.json", programs + "jsons/credentials.json", programs + "jsons/downloads.json", programs + "jsons/jira.json", programs + "jsons/logins.json", programs + "jsons/nouns.json"}
	// Declare string slices
	folder, number, prem, wpac, core []string
)
