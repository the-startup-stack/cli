package stackcli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
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

func (p *Project) CopyFiles() {
	cpCmd := exec.Command("cp", "-rf", p.TempDir, p.DirectoryName)
	err := cpCmd.Run()
	if err != nil {
		panic("Could not copy files")
	}
}
func (p *Project) CopyAndRenameFiles() {
	p.CopyFiles()

	fileList := []string{}

	err := filepath.Walk(p.DirectoryName, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	renamer := NewDirectoryRenamer(p, fileList)
	renamer.execute()
}

func (p *Project) Create() {
	p.CreateProjectDir()
	p.CreateTempDir()
	p.DownloadAndExtractZip()
}

func (p *Project) CreateProjectDir() {
	err := os.MkdirAll(p.DirectoryName, 0755)
	if err != nil {
		panic("Could not create project directory")
	}
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
