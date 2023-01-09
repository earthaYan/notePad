# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/earthaYan/notePad
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

# ==============================================================================
include scripts/make-rules/common.mk # 必须在第一个
include scripts/make-rules/golang.mk 
include scripts/make-rules/tools.mk 
include scripts/make-rules/gen.mk 

##	help:	Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'  ## 可以通过##	name:	info去自动输出
	@echo "$$USAGE_OPTIONS"

##	format:	format code by gofmt,goimports,golines
.PHONY:	format
format:	tools.verify.golines tools.verify.goimports
	@echo "===========> Formating codes"
	@$(FIND) -type f -name '*.go'	|	$(XARGS) gofmt -s -w
	@$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	@$(FIND) -type f -name '*.go' | $(XARGS) golines -w --max-len=120 --reformat-tags --shorten-comments --ignore-generated .
	@$(GO) mod edit -fmt

##	gen:	Generate all necessary files, such as error code files.
.PHONY: gen
gen:
	@$(MAKE) gen.run
