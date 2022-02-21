# Code generated by go.einride.tech/sage. DO NOT EDIT.
# To learn more, see .sage/sagefile.go and https://github.com/einride/sage.

.DEFAULT_GOAL := all

sagefile := .sage/bin/sagefile

.PHONY: $(sagefile)
$(sagefile):
	@cd .sage && go mod tidy && go run .

.PHONY: sage
sage: $(sagefile)

.PHONY: update-sage
update-sage:
	@cd .sage && go get -d go.einride.tech/sage@latest && go mod tidy && go run .

.PHONY: clean-sage
clean-sage:
	@git clean -fdx .sage/tools .sage/bin .sage/build

.PHONY: all
all: $(sagefile)
	@$(sagefile) All

.PHONY: buf-lint
buf-lint: $(sagefile)
	@$(sagefile) BufLint

.PHONY: buf-push
buf-push: $(sagefile)
	@$(sagefile) BufPush

.PHONY: clang-format-proto
clang-format-proto: $(sagefile)
	@$(sagefile) ClangFormatProto
