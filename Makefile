# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/earthaYan/notePad
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

# ==============================================================================
include scripts/make-rules/common.mk # 必须在第一个
include scripts/make-rules/golang.mk 
include scripts/make-rules/tools.mk 
include scripts/make-rules/gen.mk 

default:	genSwaggerDoc
##	help:	Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'  ## 可以通过##	name:	info去自动输出
	@echo "$$USAGE_OPTIONS"

##	genSwaggerDoc: generate swagger docs by swaggo/swag
.PHONY: genSwaggerDoc
genSwaggerDoc:
	swag init -g ./cmd/notepad-apiserver/apiserver.go