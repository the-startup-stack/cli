package stackcli

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestUtils(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestReverseArray(c *C) {
	files := []string{"a", "b", "c"}
	reversed := reverseArray(files)
	c.Assert(reversed[0], Equals, "c")
	c.Assert(reversed[1], Equals, "b")
	c.Assert(reversed[2], Equals, "a")
}
