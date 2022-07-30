.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: clean
clean:
	rm bin/main

.PHONY: cicheckconfig
cicheckconfig:
	circleci config validate

.PHONY: cilocalrun
cilocalrun:
	circleci local execute build