package stackcli

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type DirectoryRenamer struct {
	Vals     map[string]string
	FileName string
}

type Match struct {
	Key   string
	Value string
}

func NewDirectoryRenamer(project *Project, filename string) *DirectoryRenamer {
	return &DirectoryRenamer{
		FileName: filename,
		Vals: map[string]string{
			"project-name": project.ProjectName,
		},
	}
}

func (r *DirectoryRenamer) execute() string {
	matches := r.match(r.FileName)
	return r.renameMatches(r.FileName, matches)
}

func (r *DirectoryRenamer) renameMatches(origFileName string, matches []Match) string {
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

func (r *DirectoryRenamer) match(path string) []Match {
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

func (r *DirectoryRenamer) getMatch(path string) {}
