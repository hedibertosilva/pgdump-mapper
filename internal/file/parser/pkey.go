package parser

import "regexp"

func PKey(line string) string {
	re := regexp.MustCompile(`ADD CONSTRAINT (\w+) PRIMARY KEY \((\w+)\);`)
	match := re.FindStringSubmatch(line)
	if len(match) == 3 {
		// It's expected 3 elements:
		// [0] Original line
		// [1] PKey name
		// [2] Pkey column
		return match[2]
	}
	return ""
}
