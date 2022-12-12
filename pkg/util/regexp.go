package util

import (
	"log"
	"regexp"
)

func RegexpMatch(reStr string, str string) []string {
	re := regexp.MustCompile(reStr)
	submatch := re.FindAllStringSubmatch(str, -1)
	if len(submatch) == 0 {
		return []string{}
	}
	return submatch[0]
}

func ExpectMatchesLen(matches []string, n int) []string {
	if len(matches) != n {
		log.Fatal("expected", n, "matches, got", len(matches))
	}
	return matches
}
