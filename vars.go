package main

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
	Dev       string `json:"dev"`
	Core      string `json:"core"`
	WPac      string `json:"wpac"`
	Prem      string `json:"prem"`
	ToDo      string `json:"todo"`
	Token     string `json:"token"`
	Repos     string `json:"repos"`
	Cloud     string `json:"cloud"` // jira API production environment
	Programs  string `json:"programs"`
	WordPress string `json:"wordpress"`
}

// Bundle holds the information necessary to login and download premium plugin updates
type Bundle struct {
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

// Pack holds the extracted data from the JQL queries
type Pack struct {
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
	assets   string = "/data/scripts/jira-automation/programs/"
)

var (
	ecp    ECP
	evtp   EVTP
	satis  Satis
	jira   Pack
	flag   string
	plugin string
	ticket string
	values Bundle
	access Atlassian
	jsons  = []string{assets + "jsons/env.json", assets + "jsons/bundle.json"}
	// Declare string slices
	folder, number, dev, prem, wpac, core []string
)
