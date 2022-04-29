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
		Variables: []parser.Variable{{
			Name:           "LAZY_EXPANDED_TRAILING_WHITESPACE_ERROR",
			Assignment:     "\"$(FOO)\"",
			Comment:        "# Trailing comment adds trailing whitespace error to lazy-expanded variable",
			SimplyExpanded: false}},
	}

	rule := Variabletrailingwhitespace{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, 1, len(ret))
	assert.Equal(t, "avoid trailing whitespace after variable definitions",
		rule.Description())
	assert.Equal(t, "Variable \"LAZY_EXPANDED_TRAILING_WHITESPACE_ERROR\" possibly contains unintended trailing whitespace and the comment should placed on the line above.",
		ret[0].Violation)

	for i := range ret {
		assert.Equal(t, "lazy-expanded-trailing-whitespace.mk", ret[i].FileName)
	}
}

func TestSimplyExpandedTrailingWhitespaceError(t *testing.T) {

	makefile := parser.Makefile{
		FileName: "simply-expanded-trailing-whitespace.mk",
		Variables: []parser.Variable{{
			Name:           "TRAILING_WHITESPACE_ERROR",
			Assignment:     "\"$(FOO)\"",
			Comment:        "# Trailing comment adds trailing whitespace error to simply-expanded variable",
			SimplyExpanded: true}},
	}

	rule := Variabletrailingwhitespace{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, 1, len(ret))
	assert.Equal(t, "Variable \"TRAILING_WHITESPACE_ERROR\" possibly contains unintended trailing whitespace and the comment should placed on the line above.",
		ret[0].Violation)
	for i := range ret {
		assert.Equal(t, "simply-expanded-trailing-whitespace.mk", ret[i].FileName)
		assert.Equal(t, "variabletrailingwhitespace", ret[i].Rule)
	}
}
