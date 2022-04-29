# All Special Built-In Target Names for GNU Make
# Reference: https://www.gnu.org/software/make/manual/html_node/Special-Targets.html#Special-Targets

.PHONY: all noparallel clean one two three

.DEFAULT: all
all: my-precious intermediate ignoreerrors foo noparallel foo.o second-expansion

.SUFFIXES: .o
foo.c:
	touch foo.c
.o: # This is technically not a Special Built-In Target but an OLD style SUFFIX rule
	echo $(CC) -c $(CFLAGS) $(CPPFLAGS) -o $@ $<
	touch $@

.PRECIOUS: my-precious
my-precious:
	touch my-precious

.INTERMEDIATE: intermediate
intermediate: intermediate-secondary
	touch intermediate

.SECONDARY: intermediate-secondary
intermediate-secondary:
	touch intermediate-secondary


.SECONDEXPANSION:
foovar := foo
barvar := bar
bar:
	touch bar
second-expansion: $(foovar) $$(barvar)
	echo $^

.DELETE_ON_ERROR:

.IGNORE: ignoreerrors src foo
ignoreerrors:
	exit 1

.LOW_RESOLUTION_TIME: foo
src:
	touch src
	exit 1

foo: src
	cp -p src foo
	exit 1

.SILENT:

.EXPORT_ALL_VARIABLES:

.NOTPARALLEL: noparallel
noparallel: one two three
one:
	sleep 0.3
	echo 1
two:
	echo 2
three:
	sleep 0.1
	echo 3

.ONESHELL:

.POSIX:

clean:
	rm -f foo
	rm -f foo.bar
	rm -f bar
	rm -f foo.c
	rm -f foo.h
	rm -f foo.o
	rm -f src
	rm -f my-precious
	rm -f intermediate
	rm -f intermediate-secondary
