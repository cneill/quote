package prettify

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadTags() error {
	// TODO: make this configurable
	file := "./tags.json"

	contents, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read file %q contents: %w", file, err)
	}

	if err := json.Unmarshal(contents, &tags); err != nil {
		return fmt.Errorf("failed to unmarshal file %q contents: %w", file, err)
	}

	LoadCountryTags()

	return nil
}

func LoadCountryTags() {
	for country, emoji := range countries {
		countryTags[country] = fmt.Sprintf("%s %s", emoji, country)
	}
}

func AddTags(input string) string {
	tagsMatched := map[string]bool{}

	for tag, aliases := range tags {
		if strings.Contains(input, tag) {
			input = strings.Replace(input, tag, fmt.Sprintf("[[%s]]", tag), 1)
			tagsMatched[tag] = true

			continue
		}

		// check for the tag name capitalized at the start of a sentence
		startWord := strings.Title(tag)
		idx := strings.Index(input, startWord)

		if idx == 0 || (idx > 2 && input[idx-2] == '.') {
			input = strings.Replace(input, startWord, fmt.Sprintf("[%s]([[%s]])", startWord, tag), 1)

			continue
		}

		// handle aliases for the tag
	aliases:
		for _, alias := range aliases {
			if strings.Contains(input, alias) {
				input = strings.Replace(input, alias, fmt.Sprintf("[%s]([[%s]])", alias, tag), 1)

				break aliases
			}
		}
	}

	input = AddCountryTags(input)

	input = strings.ReplaceAll(input, "[[[[", "[[")
	input = strings.ReplaceAll(input, "]]]]", "]]")

	return input
}

func AddCountryTags(input string) string {
	for country, tag := range countryTags {
		if strings.Contains(input, country) {
			input = strings.Replace(input, country, fmt.Sprintf("[[%s]]", tag), 1)
		}
	}

	return input
}
