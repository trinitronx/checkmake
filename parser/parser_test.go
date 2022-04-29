package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSimpleMakefile(t *testing.T) {

	ret, err := Parse("../fixtures/simple.make")

	assert.Equal(t, err, nil)
	assert.Equal(t, ret.FileName, "../fixtures/simple.make")
	assert.Equal(t, len(ret.Rules), 7)
	assert.Equal(t, len(ret.Variables), 2)
	assert.Equal(t, ret.Rules[0].Target, "clean")
	assert.Equal(t, ret.Rules[0].Comment, "")
	assert.Equal(t, ret.Rules[0].SpecialTarget, false)
	assert.Equal(t, ret.Rules[0].Body, []string{"rm bar", "rm foo"})
	assert.Equal(t, ret.Rules[0].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[0].LineNumber, 6)

	assert.Equal(t, ret.Rules[1].Target, "foo")
	assert.Equal(t, ret.Rules[1].Body, []string{"touch foo"})
	assert.Equal(t, ret.Rules[1].Dependencies, []string{"bar"})
	assert.Equal(t, ret.Rules[1].Comment, "")
	assert.Equal(t, ret.Rules[1].SpecialTarget, false)
	assert.Equal(t, ret.Rules[1].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[1].LineNumber, 10)

	assert.Equal(t, ret.Rules[2].Target, "bar")
	assert.Equal(t, ret.Rules[2].Body, []string{"touch bar"})
	assert.Equal(t, ret.Rules[2].Comment, "")
	assert.Equal(t, ret.Rules[2].SpecialTarget, false)
	assert.Equal(t, ret.Rules[2].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[2].LineNumber, 13)

	assert.Equal(t, ret.Rules[3].Target, "all")
	assert.Equal(t, ret.Rules[3].Dependencies, []string{"foo"})
	assert.Equal(t, ret.Rules[3].Comment, "")
	assert.Equal(t, ret.Rules[3].SpecialTarget, false)
	assert.Equal(t, ret.Rules[3].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[3].LineNumber, 16)

	assert.Equal(t, ret.Rules[4].Target, "test")
	assert.Equal(t, ret.Rules[4].Dependencies, []string{})
	assert.Equal(t, ret.Rules[4].Comment, "")
	assert.Equal(t, ret.Rules[4].SpecialTarget, false)
	assert.Equal(t, ret.Rules[4].Body, []string{}) // Body begins w/ spaces instead of tab char "  @echo lolnah"
	assert.Equal(t, ret.Rules[4].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[4].LineNumber, 18)

	assert.Equal(t, ret.Rules[5].Target, ".PHONY")
	assert.Equal(t, ret.Rules[5].Dependencies, []string{"all", "clean", "test"})
	assert.Equal(t, ret.Rules[5].Comment, "")
	assert.Equal(t, ret.Rules[5].SpecialTarget, true)
	assert.Equal(t, ret.Rules[5].Body, []string{})
	assert.Equal(t, ret.Rules[5].FileName, "../fixtures/simple.make")
	assert.Equal(t, ret.Rules[5].LineNumber, 22)

	assert.Equal(t, ret.Rules[6].Target, ".DEFAULT_GOAL")
	assert.Equal(t, ret.Rules[6].Dependencies, []string{"all"})
	assert.Equal(t, ret.Rules[6].Comment, "")
	assert.Equal(t, ret.Rules[6].SpecialTarget, true)
	assert.Equal(t, ret.Rules[6].Body, []string{})
	assert.Equal(t, ret.Rules[6].FileName, "../fixtures/simple.make")
	// assert.Equal(t, ret.Rules[6].LineNumber, 23)

	assert.Equal(t, ret.Variables[0].Name, "expanded")
	assert.Equal(t, ret.Variables[0].Assignment, "\"$(simple)\"")
	assert.Equal(t, ret.Variables[0].SimplyExpanded, false)
	assert.Equal(t, ret.Variables[0].SpecialVariable, false)
	assert.Equal(t, ret.Variables[0].FileName, "../fixtures/simple.make")
	// assert.Equal(t, ret.Variables[0].LineNumber, 3)

	assert.Equal(t, ret.Variables[1].Name, "simple")
	assert.Equal(t, ret.Variables[1].Assignment, "\"foo\"")
	assert.Equal(t, ret.Variables[1].SimplyExpanded, true)
	assert.Equal(t, ret.Variables[1].SpecialVariable, false)
	assert.Equal(t, ret.Variables[1].FileName, "../fixtures/simple.make")
	// assert.Equal(t, ret.Variables[1].LineNumber, 4)
}
