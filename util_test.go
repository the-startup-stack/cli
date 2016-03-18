package stackcli

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestUtils(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestReverseArray(c *C) {
	files := []string{"a/a", "a/a/a", "a/a/a/a"}
	reversed := traverseDirStructure(files)
	c.Assert(reversed[0], Equals, "a/a/a/a")
	c.Assert(reversed[1], Equals, "a/a/a")
	c.Assert(reversed[2], Equals, "a/a")
}
