package stackcli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
)

type Project struct {
	HomeDir       string
	ProjectName   string
	DirectoryName string
	TempDir       string
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
	newDirName := fmt.Sprintf("%s/%s/", p.TempDir, "chef-repo-template-master")
	cpCmd := exec.Command("cp", "-rf", newDirName, p.DirectoryName)
	err := cpCmd.Run()
	if err != nil {
		panic("Could not copy files")
	}
}

func (p *Project) CopyAndRenameFiles() {
	p.CopyFiles()

	p.iterateDir(p.DirectoryName)
}

func (p *Project) iterateDir(startDir string) {
	dir := fmt.Sprintf("%s/", startDir)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filename := fmt.Sprintf("%s%s", dir, file.Name())
		renamer := NewRenamer(p, filename)
		newName := renamer.execute()

		if file.IsDir() {
			p.iterateDir(newName)
		}
	}
}

func (p *Project) Create() {
	p.CreateProjectDir()
	p.CreateTempDir()
	p.DownloadAndExtractZip()
	p.CopyAndRenameFiles()
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
