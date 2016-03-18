package stackcli

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestRenamer(t *testing.T) { TestingT(t) }

type RenamerSuite struct{}

var _ = Suite(&RenamerSuite{})

func (s *RenamerSuite) TestMatch(c *C) {
	renamer := NewRenamer(&Project{ProjectName: "test"}, "")
	matches := renamer.match("test/chef/{{project-name}}-cookbooks/{{project-name}}-something")
	c.Assert(len(matches), Equals, 1)
	c.Assert(matches[0].Key, Equals, "{{project-name}}")
	c.Assert(matches[0].Value, Equals, "test")
}
