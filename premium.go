package main

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// A sequential list of tasks run to complete the program
func quarterback() {
	prepare()
	tracking("Creating new branch")
	checkout(branch)
	tracking("Running update script")
	script()
	correct()
	tracking("Commiting changes")
	commit()
	tracking("Tagging to Satis")
	tags()
	tracking("Pushing to repository")
	push()
}

// Premium directs the preliminary actions to determine if the program can continue
func premium() {
	assign(prem[0], prem[1])
	os.Chdir(access.Root + folder[1])
	learn()

	if folder[1] == "gravityforms" {
		login(values.Credentials[2].Username, values.Credentials[2].Password, values.Downloads.Gravity, values.Links.Gravity)
	} else {
		directdownload("values.Downloads." + folder[1])
	}

	satis.Version, ecp.Version, evtp.Version = number[1], number[1], number[1]

	if strings.Contains(folder[1], "event") {
		if ecp.Name+":"+ecp.Version == plugin || evtp.Name+":"+evtp.Version == plugin {
			quarterback()
		}
	} else if satis.Name+":"+satis.Version == plugin {
		quarterback()
	} else {
		alert("Plugin name does not match composer.json entry - program halted")
	}
}

// Split the supplied arguments and assign them to variables
func assign(p, t string) {
	plugin, ticket = p, t
	number = strings.Split(plugin, ":")
	folder = strings.Split(number[0], "/")
}

// Read the composer.json file and store the results in a structure
func learn() {
	current, _ := os.ReadFile("composer.json")
	err := json.Unmarshal(current, &satis)
	inspect(err)
	err = json.Unmarshal(current, &ecp)
	inspect(err)
	err = json.Unmarshal(current, &evtp)
	inspect(err)
}

func login(username, password, download, login string) {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	inspect(err)
	client := http.Client{Jar: jar}
	client.PostForm(login, url.Values{
		"password": {password},
		"username": {username},
	})

	// execute("-e", "curl", download, "-o", "~/Downloads/"+filename)
	directdownload(download)
}

func directdownload(value string) {
	execute("-e", "curl", value, "-o", "~/Downloads/"+folder[1])
}

// Create an update branch if necessary
func checkout(prefix string) {
	if exists(prefix, ticket) {
		execute("-e", "git", "switch", prefix+ticket)
	} else {
		execute("-e", "git", "checkout", "-b", prefix+ticket)
	}
}

// Run the update script on downloaded content
func script() {
	execute("-e", "sh", "-c", "scripts/update.sh ~/Downloads/"+folder[1]+"/")
}

// Convert the structure back into json and overwrite the composer.json file
func correct() {
	var updated []byte
	if strings.Contains(ecp.Name, "calendar") {
		updated, _ = json.MarshalIndent(ecp, "", "    ")
	} else if strings.Contains(evtp.Name, "tickets") || strings.Contains(evtp.Name, "virtual") {
		updated, _ = json.MarshalIndent(evtp, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(satis, "", "    ")
	}
	document("composer.json", updated)
}

// Tag the version so Satis can package it
func tags() {
	execute("-e", "git", "tag", "v"+satis.Version)
	execute("-e", "git", "push", "origin", "--tags")
}
