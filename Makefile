include scripts/make-rules/common.mk
## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'  ## 可以通过##	name:	info去自动输出
	@echo "$$USAGE_OPTIONS"