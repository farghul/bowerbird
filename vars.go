package main

const (
	bv       string = "2.0"
	branch   string = "update/"
	reset    string = "\033[0m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgred    string = "\033[41m"
	bgyellow string = "\033[43m"
	halt     string = "program halted "
	zero     string = "Not enough arguments supplied -"
)

var (
	ecp   ECP
	evtp  EVTP
	satis Satis
	// vendor Links
	// creds  Login
	jira   Ticket
	access Atlassian
	// update Download
	values Composite
	flag   string
	plugin string
	ticket string
	jsons  = []string{"jsons/access.json", "jsons/composite.json"}
	// Declare string slices
	number, folder, wpac, prem, dev []string
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
	Dev   string `json:"dev"`
	WPac  string `json:"wpac"`
	Prem  string `json:"prem"`
	Repo  string `json:"repo"`
	Root  string `json:"root"`
	Base  string `json:"base"`
	Token string `json:"token"`
	ToDo  string `json:"todo"`
}

type Links struct {
	AllImport string `json:"allimport"`
	Gravity   string `json:"gravity"`
	PolyLang  string `json:"polylang"`
	SeachWP   string `json:"searchwp"`
}

type Login struct {
	Credentials []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

type Download struct {
	AllImport string `json:"allimport"`
	Gravity   string `json:"gravity"`
	PolyLang  string `json:"polylang"`
	SeachWP   string `json:"searchwp"`
}

type Composite struct {
	Credentials []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
	Downloads struct {
		Polylang  string `json:"polylang-pro"`
		AllExport string `json:"wp-all-export-pro"`
		Gravity   string `json:"gravityforms"`
		SearchWP  string `json:"searchwp"`
		Calendar  string `json:"events-calendar-pro"`
		Tickets   string `json:"event-tickets-plus"`
		Virtual   string `json:"events-virtual"`
		Uji       string `json:"uji-countdown-premium"`
	} `json:"downloads"`
	Links struct {
		Polylang  string `json:"polylang-pro"`
		Allimport string `json:"wp-all-export-pro"`
		Gravity   string `json:"gravityforms"`
		Searchwp  string `json:"searchwp"`
	} `json:"links"`
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
