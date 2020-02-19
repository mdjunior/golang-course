# Run it in the problem's folder as `make -f ../makefile`

export GOPATH=$(CURDIR)

TESTS = $(shell find . -name '*.in')

.PHONY: all format new $(TESTS)

all: format $(TESTS)

format:
	go fmt main.go

$(TESTS):
	go run main.go < $@ | diff -w $(@:.in=.out) -
