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
	BBA       string `json:"bba"`
	BitBucket string `json:"bitbucket"`
	Jira      string `json:"jira"`
	JQA       string `json:"jqa"`
	Core      string `json:"core"`
	Prem      string `json:"prem"`
	Repos     string `json:"repos"`
	Reviewer1 string `json:"reviewer1"`
	Reviewer2 string `json:"reviewer2"`
	ToDo      string `json:"todo"`
	WordPress string `json:"wordpress"`
	WPac      string `json:"wpac"`
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
	Logins struct {
		Polylang  string `json:"polylang-pro"`
		Allimport string `json:"wp-all-export-pro"`
		Gravity   string `json:"gravityforms"`
		SearchWP  string `json:"searchwp"`
	} `json:"logins"`
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
	bv       string = "1.0.0"
	branch   string = "update/"
	reset    string = "\033[0m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgred    string = "\033[41m"
	bgyellow string = "\033[43m"
	halt     string = "program halted "
	zero     string = "Not enough arguments supplied -"
	assets   string = "/data/scripts/automation/programs/"
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
	folder, number, prem, wpac, core []string
)
