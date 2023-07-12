package prettify

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadTags parses a JSON file (./tags.json) in map[string][]string format and uses them to apply tags to inputs.
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

// LoadCountryTags populates the country tags from the 'countries' variable with their flag emoji.
func LoadCountryTags() {
	for country, emoji := range countries {
		countryTags[country] = fmt.Sprintf("%s %s", emoji, country)
	}
}

// AddTags looks for instances of tags or their aliases and wraps them appropriately. It might not behave as you expect.
//
//   - It will stop trying to match a tag or its aliases once one has been found.
//   - It will replace country names (e.g. "USA") with a tag with their emoji (e.g. "[[ (flag) USA ]]"
//   - It will look for instances of a tag that look like the first word in a string and wrap them in an alias.
func AddTags(input string) string {
	tagsMatched := map[string]bool{}

	for tag, aliases := range tags {
		if strings.Contains(input, tag) {
			input = strings.Replace(input, tag, fmt.Sprintf("[[%s]]", tag), 1)
			tagsMatched[tag] = true

			continue
		}

		// check for the tag name, capitalized, with a "." 2 characters before it (presumably, at the start of a sentence)
		startWord := strings.Title(tag)
		if startWord != tag {
			idx := strings.Index(input, startWord)

			if idx == 0 || (idx > 2 && input[idx-2] == '.') {
				input = strings.Replace(input, startWord, fmt.Sprintf("[%s]([[%s]])", startWord, tag), 1)

				continue
			}
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

// AddCountryTags adds the tags for countries as described in [AddTags].
func AddCountryTags(input string) string {
	for country, tag := range countryTags {
		input = strings.Replace(input, country, fmt.Sprintf("[[%s]]", tag), 1)
	}

	return input
}
