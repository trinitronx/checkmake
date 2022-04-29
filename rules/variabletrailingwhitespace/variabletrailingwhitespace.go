// Package variabletrailingwhitespace implements the ruleset for making sure a variable
// that is followed by a comment is flagged as a trailing whitespace error
// Ideally a Makefile variable shouldn't include a trailing comment because
// it introduces whitespace between the variable definition and comment
// Managing Projects with GNU Make, 3rd Edition by Robert Mecklenburg
// states the following:
//
// The value of a variable consists of all the words to the right of the
// assignment symbol with leading space trimmed. Trailing spaces are not trimmed
// This can occasionally cause trouble, for instance, if the trailing whitespace
// is included in the variable and subsequently used in a command script
// Source:
//   https://www.oreilly.com/library/view/managing-projects-with/0596006101/ch03.html
//
package variabletrailingwhitespace

import (
	"fmt"
	"regexp"

	//"strings"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
)

func init() {
	rules.RegisterRule(&Variabletrailingwhitespace{})
}

// Variabletrailingwhitespace is an empty struct on which to call the rule functions
type Variabletrailingwhitespace struct {
}

var (
	vT                         = "Variable %q possibly contains unintended trailing whitespace and the comment should placed on the line above."
	whitespaceCommentRegexp, _ = regexp.Compile(`\s+#.*$`)
)

// Name returns the name of the rule
func (r *Variabletrailingwhitespace) Name() string {
	return "variabletrailingwhitespace"
}

// Description returns the description of the rule
func (r *Variabletrailingwhitespace) Description() string {
	return "avoid trailing whitespace after variable definitions"
}

// Run executes the rule logic
func (r *Variabletrailingwhitespace) Run(makefile parser.Makefile, config rules.RuleConfig) rules.RuleViolationList {
	ret := rules.RuleViolationList{}

	for _, variable := range makefile.Variables {
		if whitespaceCommentRegexp.MatchString(variable.Assignment) &&
			!variable.SpecialVariable {
			ret = append(ret, rules.RuleViolation{
				Rule:       r.Name(),
				Violation:  fmt.Sprintf(vT, variable.Name),
				FileName:   makefile.FileName,
				LineNumber: variable.LineNumber,
			})
		}
	}

	return ret
}
