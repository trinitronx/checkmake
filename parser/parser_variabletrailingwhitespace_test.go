package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTestTrailingWhitespaceMakefile(t *testing.T) {

	ret, err := Parse("../fixtures/trailing_whitespace.make")

	assert.Equal(t, err, nil)
	assert.Equal(t, ret.FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, len(ret.Rules), 7)
	assert.Equal(t, len(ret.Variables), 3)

	assert.Equal(t, ret.Rules[0].Target, "clean")
	assert.Equal(t, ret.Rules[0].Dependencies, []string{})
	assert.Equal(t, ret.Rules[0].Comment, "## Comments here are OK... this could be help text")
	assert.Equal(t, ret.Rules[0].SpecialTarget, false)
	assert.Equal(t, ret.Rules[0].Body, []string{"rm bar", "rm foo"})
	assert.Equal(t, ret.Rules[0].FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, ret.Rules[0].LineNumber, 8)

	assert.Equal(t, ret.Rules[1].Target, "foo")
	assert.Equal(t, ret.Rules[1].Dependencies, []string{"bar"})
	assert.Equal(t, ret.Rules[1].Comment, "# Regular comment after prerequisite OK")
	assert.Equal(t, ret.Rules[1].SpecialTarget, false)
	assert.Equal(t, ret.Rules[1].Body, []string{"touch foo # Comment after rule OK"})
	assert.Equal(t, ret.Rules[1].FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, ret.Rules[1].LineNumber, 12)

	assert.Equal(t, ret.Rules[2].Target, "bar")
	assert.Equal(t, ret.Rules[2].Dependencies, []string{})
	assert.Equal(t, ret.Rules[2].Comment, "# Regular comment including special char # OK")
	assert.Equal(t, ret.Rules[2].SpecialTarget, false)
	assert.Equal(t, ret.Rules[2].Body, []string{"touch bar # Comment after rule including # char OK"})
	assert.Equal(t, ret.Rules[2].FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, ret.Rules[2].LineNumber, 15)

	assert.Equal(t, ret.Rules[3].Target, "all")
	assert.Equal(t, ret.Rules[3].Dependencies, []string{"foo"})
	assert.Equal(t, ret.Rules[3].Comment, "# Regular comment including variable $(FOO) OK")
	assert.Equal(t, ret.Rules[3].SpecialTarget, false)
	assert.Equal(t, ret.Rules[3].Body, []string{})
	assert.Equal(t, ret.Rules[3].FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, ret.Rules[3].LineNumber, 18)

	assert.Equal(t, ret.Rules[4].Target, "test")
	assert.Equal(t, ret.Rules[4].Dependencies, []string{})
	assert.Equal(t, ret.Rules[4].Comment, "# Regular comment including many special chars ~!@#$%^&*()_+``-=<>,./? OK")
	assert.Equal(t, ret.Rules[4].SpecialTarget, false)
	assert.Equal(t, ret.Rules[4].Body, []string{"@echo lolnah # Comment after rule including many special chars ~!@#$%^&*()_+``-=<>,./?:;'\"[]{}\\\\|\\ - OK"})
	assert.Equal(t, ret.Rules[4].FileName, "../fixtures/trailing_whitespace.make")
	assert.Equal(t, ret.Rules[4].LineNumber, 20)

	assert.Equal(t, ret.Rules[5].Target, ".PHONY")
	assert.Equal(t, ret.Rules[5].Dependencies, []string{"all", "clean", "test"})
	assert.Equal(t, ret.Rules[5].Comment, "# Phony special rule comment OK")
	assert.Equal(t, ret.Rules[5].SpecialTarget, true)
	assert.Equal(t, ret.Rules[5].Body, []string{})
	assert.Equal(t, ret.Rules[5].FileName, "../fixtures/trailing_whitespace.make")
	// assert.Equal(t, ret.Rules[5].LineNumber, 23)

	assert.Equal(t, ret.Rules[6].Target, ".DEFAULT_GOAL")
	assert.Equal(t, ret.Rules[6].Dependencies, []string{"all"})
	assert.Equal(t, ret.Rules[6].Comment, "# .DEFAULT_GOAL: rule comment OK")
	assert.Equal(t, ret.Rules[6].SpecialTarget, true)
	assert.Equal(t, ret.Rules[6].Body, []string{})
	assert.Equal(t, ret.Rules[6].FileName, "../fixtures/trailing_whitespace.make")
	// assert.Equal(t, ret.Rules[6].LineNumber, 25)

}
