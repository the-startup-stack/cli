package stackcli

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
)

type Project struct {
	HomeDir       string
	ProjectName   string
	DirectoryName string
	TempDir       string
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) CreateTempDir() {
	os.MkdirAll(p.TempDir, 0755)
}

func (p *Project) DownloadAndExtractZip() {
	zipFile := fmt.Sprintf("%v/chef.zip", p.TempDir)
	downloadFile(zipFile, "https://github.com/the-startup-stack/chef-repo-template/archive/master.zip")
	unzip(zipFile, p.TempDir)
}

func (p *Project) Create() {
	p.CreateTempDir()
	p.DownloadAndExtractZip()
}

func NewProject(projectName string, directoryName string) *Project {
	usr, err := user.Current()

	if err != nil {
		panic("Could not get current user")
	}

	return &Project{
		HomeDir:       usr.HomeDir,
		ProjectName:   projectName,
		DirectoryName: directoryName,
		TempDir:       fmt.Sprintf("%v/.the-startup-stack", usr.HomeDir),
	}
}
