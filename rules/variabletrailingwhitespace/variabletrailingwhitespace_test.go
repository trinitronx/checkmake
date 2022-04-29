package variabletrailingwhitespace

import (
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
)

func TestLazyExpandedTrailingWhitespaceError(t *testing.T) {

	makefile := parser.Makefile{
		FileName: "lazy-expanded-trailing-whitespace.mk",
		Variables: []parser.Variable{parser.Variable{
			Name:           "LAZY_EXPANDED_TRAILING_WHITESPACE_ERROR",
			Assignment:     "\"$(FOO)\" # Trailing comment adds trailing whitespace error to lazy-expanded variable",
			SimplyExpanded: false}},
	}

	rule := Variabletrailingwhitespace{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, 1, len(ret))
	assert.Equal(t, "avoid trailing whitespace after variable definitions",
		rule.Description())
	for i := range ret {
		assert.Equal(t, "lazy-expanded-trailing-whitespace.mk", ret[i].FileName)
	}
}

func TestSimplyExpandedTrailingWhitespaceError(t *testing.T) {

	makefile := parser.Makefile{
		FileName: "simply-expanded-trailing-whitespace.mk",
		Variables: []parser.Variable{parser.Variable{
			Name:           "TRAILING_WHITESPACE_ERROR",
			Assignment:     "\"$(FOO)\" # Trailing comment adds trailing whitespace error to simply-expanded variable",
			SimplyExpanded: true}},
	}

	rule := Variabletrailingwhitespace{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, 1, len(ret))
	for i := range ret {
		assert.Equal(t, "simply-expanded-trailing-whitespace.mk", ret[i].FileName)
		assert.Equal(t, "variabletrailingwhitespace", ret[i].Rule)
	}
}
