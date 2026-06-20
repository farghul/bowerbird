package workers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/farghul/bowerbird/internal"
)

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func Serialize() {
	internal.Clearout(internal.Temp)
	for index, element := range internal.Jsons {
		data, err := os.ReadFile(element)
		internal.Inspect(err)
		switch index {
		case 0:
			err := json.Unmarshal(data, &internal.Defs)
			internal.Inspect(err)
		case 1:
			err := json.Unmarshal(data, &internal.Jira)
			internal.Inspect(err)
		}
	}
}

// Compile the results of a Jira API query and save summary and key into a string slice
func Compiler(element string) []string {
	var data []byte
	var err error
	if element == "premium" {
		data, err = API(internal.Jira.Basic + internal.Jira.Review)
		internal.Inspect(err)
	} else {
		data, err = API(internal.Jira.Basic + internal.Jira.ToDo)
		internal.Inspect(err)
	}
	err = json.Unmarshal(data, &internal.Query)
	internal.Inspect(err)

	var candidate []string
	for i := range internal.Query.Issues {
		if strings.Contains(internal.Query.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, internal.Query.Issues[i].Fields.Summary)
			candidate = append(candidate, internal.Query.Issues[i].Key)
		}
	}
	return candidate
}

func API(criteria string) ([]byte, error) {
	baseURL := internal.Jira.URL + "search/jql?jql="

	fullURL := baseURL + criteria

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Basic "+internal.Jira.Token)
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
func Rightplace() {
	err := os.Chdir(internal.Defs.WordPress)
	internal.Inspect(err)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		internal.Alert("This is not the correct folder,")
	}
}

// Check for edge cases which require the -W flag
func Edge() bool {
	found := false
	if strings.Contains(internal.Plugin, "roots/wordpress") {
		found = true
	}
	return found
}
