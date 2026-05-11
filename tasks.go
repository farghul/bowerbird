package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

// Test for an optional flag
func flag() string {
	var passed string
	if len(os.Args) == 1 {
		passed = "--zero"
	} else {
		passed = os.Args[1]
	}
	return passed
}

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	clearout(temp)
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			err := json.Unmarshal(data, &defs)
			inspect(err)
		case 1:
			err := json.Unmarshal(data, &jira)
			inspect(err)
		}
	}
}

// Compile the results of a Jira API query and save summary and key into a string slice
func compiler(element string) []string {
	var data []byte
	var err error
	if element == "premium" {
		data, err = api(jira.Basic + jira.Review)
		inspect(err)
	} else {
		data, err = api(jira.Basic + jira.ToDo)
		inspect(err)
	}
	err = json.Unmarshal(data, &query)
	inspect(err)

	var candidate []string
	for i := range query.Issues {
		if strings.Contains(query.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, query.Issues[i].Fields.Summary)
			candidate = append(candidate, query.Issues[i].Key)
		}
	}
	return candidate
}

func api(criteria string) ([]byte, error) {
	baseURL := jira.URL + "search/jql?jql="

	fullURL := baseURL + criteria

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Basic "+jira.Token)
	req.Header.Set("Accept", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Confirm the current working directory is correct
func rightplace() {
	err := os.Chdir(defs.WordPress)
	inspect(err)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Check for edge cases which require the -W flag
func edge() bool {
	found := false
	if strings.Contains(plugin, "roots/wordpress") {
		found = true
	}
	return found
}
