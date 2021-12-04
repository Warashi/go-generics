GO ?= go

.PHONY: test
test:
	$(GO) test ./...

.PHONY: build
build:
	$(GO) build ./...

.PHONY: commit
commit:
	npx cz

.PHONY: merge
merge:
	gh pr create
	gh pr merge --auto --merge