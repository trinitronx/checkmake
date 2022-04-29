# Regular comment OK

LAZY_EXPANDED_TRAILING_WHITESPACE_ERROR = "$(FOO)" # Trailing comment adds trailing whitespace error to lazy-expanded variable
# No comment below OK
FOO := "foo"
TRAILING_WHITESPACE_ERROR := "$(FOO)" # Trailing comment adds trailing whitespace error to simply-expanded variable

clean: ## Comments here are OK... this could be help text
	rm bar
	rm foo

foo: bar # Regular comment after prerequisite OK
	touch foo # Comment after rule OK

bar: # Regular comment including special char # OK
	touch bar # Comment after rule including # char OK

all: foo # Regular comment including variable $(FOO) OK

test: # Regular comment including many special chars ~!@#$%^&*()_+``-=<>,./? OK
	@echo lolnah # Comment after rule including many special chars ~!@#$%^&*()_+``-=<>,./?:;'"[]{}\\|\ - OK

.PHONY: all clean test # Phony special rule comment OK

.DEFAULT_GOAL: all # .DEFAULT_GOAL: rule comment OK
