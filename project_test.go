package stackcli

import (
	. "gopkg.in/check.v1"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProject(t *testing.T) { TestingT(t) }

type ProjectSuite struct{}

var _ = Suite(&ProjectSuite{})

func CleanupTest() {
	rmCmd := exec.Command("rm", "-rf", "test/chef")
	rmCmd.Run()
}

func (s *ProjectSuite) TestDirectoryStructure(c *C) {
	CleanupTest()

	project := &Project{
		HomeDir:       `irrelevant`,
		ProjectName:   `test`,
		DirectoryName: `test/chef`,
		TempDir:       `test/fixtures/`,
	}
	project.CreateProjectDir()
	project.CopyAndRenameFiles()

	fileList := []string{}

	err := filepath.Walk(project.DirectoryName, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	c.Assert(len(fileList), Equals, 6)
	c.Assert(fileList[1], Equals, "test/chef/test-cookbooks")
	c.Assert(fileList[3], Equals, "test/chef/test-cookbooks/test-something")
	c.Assert(fileList[5], Equals, "test/chef/test-file")
}
