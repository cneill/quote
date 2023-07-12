package prettify

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	replacements = map[string]string{
		"”":  "\"",
		"“":  "\"",
		"’":  "'",
		"—":  "-",
		"\n": " ",
		"  ": " ",
	}
	tags        = map[string][]string{}
	countryTags = map[string]string{}
	wordwrap    = regexp.MustCompile(`([a-zA-Z])- ([a-zA-Z])`)
)

func Prettify(input string) string {
	input = ReplaceGarbage(input)
	input = AddTags(input)

	return fmt.Sprintf("> %s #quote", input)
}

func ReplaceGarbage(input string) string {
	for in, out := range replacements {
		input = strings.ReplaceAll(input, in, out)
	}

	input = wordwrap.ReplaceAllString(input, "$1$2")
	input = strings.TrimSpace(input)

	return input
}
