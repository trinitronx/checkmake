package minphony

import (
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/stretchr/testify/assert"
)

var mpRunTests = []struct {
	mf parser.Makefile
	vl rules.RuleViolationList
}{
	{
		mf: parser.Makefile{
			FileName: "green-eggs.mk",
			Rules: parser.RuleList{
				{Target: "green-eggs"},
				{Target: "ham"},
			},
			Variables: parser.VariableList{
				{Name: "PHONY", Assignment: "green-eggs ham"},
			},
		},
		vl: rules.RuleViolationList{
			rules.RuleViolation{
				Rule:       "minphony",
				Violation:  "Missing required phony target \"kleen\"",
				FileName: "green-eggs.mk",
				LineNumber: -1,
			},
			rules.RuleViolation{
				Rule:       "minphony",
				Violation:  "Missing required phony target \"awl\"",
				FileName: "green-eggs.mk",
				LineNumber: -1,
			},
			rules.RuleViolation{
				Rule:       "minphony",
				Violation:  "Missing required phony target \"toast\"",
				FileName: "green-eggs.mk",
				LineNumber: -1,
			},
		},
	},
	{
		mf: parser.Makefile{
			FileName: "kleen.mk",
			Rules: parser.RuleList{
				{Target: "awl"},
				{Target: "distkleen"},
				{Target: "kleen"},
			},
			Variables: parser.VariableList{
				{Name: "PHONY", Assignment: "awl kleen distkleen"},
			},
		},
		vl: rules.RuleViolationList{
			rules.RuleViolation{
				Rule:       "minphony",
				Violation:  "Missing required phony target \"toast\"",
				FileName:   "kleen.mk",
				LineNumber: -1,
			},
		},
	},
}

func TestMinPhony_new(t *testing.T) {
	mp := &MinPhony{required: []string{"oh", "hai"}}

	assert.Equal(t, []string{"oh", "hai"}, mp.required)
	assert.Equal(t, "minphony", mp.Name())
	assert.Equal(t, "Minimum required phony targets must be present", mp.Description())
}

func TestMinPhony_Run(t *testing.T) {
	mp := &MinPhony{required: []string{"kleen", "awl", "toast"}}

	for _, test := range mpRunTests {
		assert.Equal(t, test.vl, mp.Run(test.mf, rules.RuleConfig{}))
	}
}
