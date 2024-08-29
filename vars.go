package main

import (
	"bufio"
	"os"
)

const (
	reset     string = "\033[0m"
	green     string = "\033[32m"
	yellow    string = "\033[33m"
	red       string = "\033[41m"
	bv        string = "2.0"
	relbranch string = "release/"
	upbranch  string = "update/DESSO-"
	halt      string = "program halted "
	zero      string = "Not enough arguments supplied -"
	bitbucket string = "/BitBucket/"
)

var (
	ecp     ECP
	evtp    EVTP
	satis   Satis
	desso   Ticket
	jira    Atlassian
	flag    string
	release string
	plugin  string
	ticket  string
	route   = os.Args[1]
	hmdr, _ = os.UserHomeDir()
	inputs  = len(desso.Issues)
	reader  = bufio.NewReader(os.Stdin)
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
	Team    string `json:"team"`
	Base    string `json:"base"`
	Path    string `json:"path"`
	Token   string `json:"token"`
	Issue   string `json:"issue"`
	Split   string `json:"split"`
	Search  string `json:"search"`
	Project string `json:"project"`
}

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
