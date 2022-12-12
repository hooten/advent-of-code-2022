package util

import (
	"log"
	"regexp"
	"strings"
)

func RegexpMatch(reStr string, str string) ([]string, bool) {
	re := regexp.MustCompile(reStr)
	submatch := re.FindAllStringSubmatch(str, -1)
	if len(submatch) == 0 {
		return nil, false
	}
	return submatch[0], true
}

func MustRegexpMatch(reStr string, str string) []string {
	match, ok := RegexpMatch(reStr, str)
	if !ok {
		log.Fatal("bad match", reStr, str)
	}
	return match
}

func RegexpMatchMultiLine(reStr string, str string) [][]string {
	reLines := Map(func(s string) string {
		return "^" + s + "$"
	}, Filter(func(s string) bool {
		return s != ""
	}, strings.Split(reStr, "\n")))
	lines := Filter(func(s string) bool {
		return s != ""
	}, strings.Split(str, "\n"))
	var matches [][]string
	for i, reLine := range reLines {
		match, _ := RegexpMatch(reLine, lines[i])
		matches = append(matches, match)
	}
	return matches
}
