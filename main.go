package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// SlackMappings represents the structure of the slack-mappings.yaml file
type SlackMappings struct {
	Mappings map[string]string `yaml:"mappings"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: slack-mapper <github-username>")
	}

	githubActor := os.Args[1]

	// Read the mapping file
	mappingFile := "slack-mappings.yaml"
	data, err := ioutil.ReadFile(mappingFile)
	if err != nil {
		log.Fatalf("Error reading mapping file %s: %v", mappingFile, err)
	}

	// Parse YAML
	var mappings SlackMappings
	err = yaml.Unmarshal(data, &mappings)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Look up the GitHub actor in mappings
	slackUserID, exists := mappings.Mappings[githubActor]
	if !exists {
		log.Fatalf("GitHub actor '%s' not found in slack mappings", githubActor)
	}

	// Output the Slack user ID
	fmt.Print(slackUserID)
}