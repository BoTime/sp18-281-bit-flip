package parse

import "strings"

// Takes a stringified comma separated list and
// returns array of separated elements
func SplitCommaSeparated(input string) []string {
	splitInput := strings.Split(input, ",")
	for i := 0; i < len(splitInput); i++ {
		splitInput[i] = strings.TrimSpace(splitInput[i])
	}
	return splitInput
}