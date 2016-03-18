package stackcli

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Renamer struct {
	Vals     map[string]string
	FileName string
}

type Match struct {
	Key   string
	Value string
}

func NewRenamer(project *Project, filename string) *Renamer {
	return &Renamer{
		FileName: filename,
		Vals: map[string]string{
			"project-name": project.ProjectName,
		},
	}
}

func (r *Renamer) execute() string {
	matches := r.match(r.FileName)
	return r.renameMatches(r.FileName, matches)
}

func (r *Renamer) renameMatches(origFileName string, matches []Match) string {
	newFileName := origFileName
	for _, match := range matches {
		newFileName = strings.Replace(newFileName, match.Key, match.Value, -1)
	}

	err := os.Rename(origFileName, newFileName)
	if err != nil {
		panic(err)
	}

	return newFileName
}

func (r *Renamer) match(path string) []Match {
	matches := []Match{}

	for key, value := range r.Vals {
		newKey := fmt.Sprintf("{{%s}}", key)
		matchTest := fmt.Sprintf("\\{\\{%s\\}\\}", key)
		isMatch, _ := regexp.MatchString(matchTest, path)
		if isMatch {
			matches = append(matches, Match{
				Key:   newKey,
				Value: value,
			})
		}
	}

	return matches
}

func (r *Renamer) getMatch(path string) {}
