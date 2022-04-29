package validator

import (
	"testing"

	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	violations := Validate(parser.Makefile{}, &config.Config{})
	assert.Equal(t, 3, len(violations))
}

func TestValidatorTrailingWhitespaceFixture(t *testing.T) {
	violations := Validate(parser.MustParse("../fixtures/trailing_whitespace.make"), &config.Config{})
	_ = violations
	assert.Equal(t, 2, len(violations))
	for i := range violations {
		// assert.Equal(t, "foo", violations[i].FileName)
		assert.Equal(t, "variabletrailingwhitespace", violations[i].Rule)
	}

	assert.Equal(t, violations[0].Violation,
		"Variable \"LAZY_EXPANDED_TRAILING_WHITESPACE_ERROR\" possibly contains unintended trailing whitespace and the comment should placed on the line above.")
	// assert.Equal(t, violations[0].LineNumber, 3)
	assert.Equal(t, violations[0].FileName, "../fixtures/trailing_whitespace.make")

	assert.Equal(t, violations[1].Violation,
		"Variable \"TRAILING_WHITESPACE_ERROR\" possibly contains unintended trailing whitespace and the comment should placed on the line above.")
	// assert.Equal(t, violations[1].LineNumber, 6)
	assert.Equal(t, violations[1].FileName, "../fixtures/trailing_whitespace.make")
}
