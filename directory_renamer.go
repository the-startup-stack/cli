package stackcli

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type DirectoryRenamer struct {
	Files []string
	Vals  map[string]string
}

type Match struct {
	Key   string
	Value string
}

func NewDirectoryRenamer(project *Project, fileList []string) *DirectoryRenamer {
	return &DirectoryRenamer{
		Files: fileList,
		Vals: map[string]string{
			"project-name": project.ProjectName,
		},
	}
}

func (r *DirectoryRenamer) execute() {
	reversedFileList := reverseArray(r.Files)

	for _, file := range reversedFileList {
		matches := r.match(file)
		fmt.Println("File: ", file)
		r.renameMatches(file, matches)
	}
}

func (r *DirectoryRenamer) renameMatches(origFileName string, matches []Match) {
	newFileName := origFileName
	for _, match := range matches {
		newFileName = strings.Replace(newFileName, match.Key, match.Value, -1)
		fmt.Println(newFileName)
	}
	err := os.Rename(origFileName, newFileName)
	if err != nil {
		fmt.Println("Error renaming")
	}
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
