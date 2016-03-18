package stackcli

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestDirectoryRenamer(t *testing.T) { TestingT(t) }

type DirectoryRenamerSuite struct{}

var _ = Suite(&DirectoryRenamerSuite{})

func (s *DirectoryRenamerSuite) TestMatch(c *C) {
	filelist := []string{}
	renamer := NewDirectoryRenamer(&Project{ProjectName: "test"}, filelist)
	matches := renamer.match("test/chef/{{project-name}}-cookbooks/{{project-name}}-something")
	c.Assert(len(matches), Equals, 1)
	c.Assert(matches[0].Key, Equals, "{{project-name}}")
	c.Assert(matches[0].Value, Equals, "test")
}
