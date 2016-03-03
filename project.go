package stackcli

import (
	"fmt"
	"os/user"
)

type Project struct {
	HomeDir string
}

func (p *Project) CloneGit() {

}

func (p *Project) createTempDir() {
	fmt.Println("Creating tmp directory .the-startup-stack at %v", p.HomeDir)

	projectDir := `.the-startup-stack`
	dir := fmt.Sprintf("%v/%v", p.HomeDir, projectDir)
	fmt.Println("dir: %v", dir)
}

func (p *Project) Create() {
	p.createTempDir()
}

func NewProject() *Project {
	usr, err := user.Current()
	if err != nil {
		panic("Could not get current user")
	}

	return &Project{
		HomeDir: usr.HomeDir,
	}
}
