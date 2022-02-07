GO ?= go

.PHONY: test
test:
	$(GO) test ./...

.PHONY: build
build:
	$(GO) build ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: commit
commit:
	npx cz

.PHONY: merge
merge:
	gh pr create --fill
	gh pr merge --auto --merge
