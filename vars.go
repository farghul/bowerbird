package main

const (
	reset     string = "\033[0m"
	green     string = "\033[32m"
	yellow    string = "\033[33m"
	bgred     string = "\033[41m"
	bgyellow  string = "\033[43m"
	bv        string = "2.0"
	relbranch string = "release/"
	bitbucket string = "/BitBucket/"
	upbranch  string = "update/DESSO-"
	halt      string = "program halted "
	zero      string = "Not enough arguments supplied -"
)

var (
	ecp    ECP
	evtp   EVTP
	satis  Satis
	jira   Ticket
	access Atlassian
	flag   string
	plugin string
	ticket string
	// Declare string slices
	number, folder, free, paid, dev []string
)

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
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

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	Dev    string `json:"dev"`
	Free   string `json:"free"`
	Paid   string `json:"paid"`
	Base   string `json:"base"`
	Token  string `json:"token"`
	Search string `json:"search"`
}

// Ticket holds the extracted data from the JQL queries
type Ticket struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
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
			Updated     string        `json:"updated"`
			Summary     string        `json:"summary"`
			FixVersions []interface{} `json:"fixVersions"`
			Created     string        `json:"created"`
		} `json:"fields"`
	} `json:"issues"`
}
