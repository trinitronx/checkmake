package phonydeclared

import (
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
)

func TestAllTargetsArePhony(t *testing.T) {

	makefile := parser.Makefile{
		FileName:  "phony-declared-all-phony.mk",
		Variables: []parser.Variable{{}},
		Rules: []parser.Rule{
			{Target: "all"},
			{Target: "clean"},
			{
				Target:        ".PHONY",
				Dependencies:  []string{"all", "clean"},
				Comment:       "",
				SpecialTarget: true,
				Body:          []string{},
				FileName:      "phony-declared-missing-one-phony.mk",
				LineNumber:    -1,
			},
		},
	}

	rule := Phonydeclared{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, len(ret), 0)

}

func TestMissingOnePhonyTarget(t *testing.T) {

	makefile := parser.Makefile{
		FileName:  "phony-declared-missing-one-phony.mk",
		Variables: []parser.Variable{{}},
		Rules: []parser.Rule{
			{Target: "all"},
			{Target: "clean"},
			{
				Target:        ".PHONY",
				Dependencies:  []string{"all"},
				Comment:       "",
				SpecialTarget: true,
				Body:          []string{},
				FileName:      "phony-declared-missing-one-phony.mk",
				LineNumber:    -1,
			},
		}}

	rule := Phonydeclared{}

	ret := rule.Run(makefile, rules.RuleConfig{})

	assert.Equal(t, len(ret), 1)

	for i := range ret {
		assert.Equal(t, "phony-declared-missing-one-phony.mk", ret[i].FileName)
	}
}
