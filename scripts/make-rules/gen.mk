#gen.run: gen.errcode gen.docgo
.PHONY:	gen.run
gen.run:	gen.clean gen.errcode gen.docgo.doc

.PHONY:gen.clean
gen.clean:
	@rm -rf ./api/client/{clientset,informers,listers}
	@$(FIND) -type f -name '*_generated.go' -delete

.PHONY:	gen.errcode
gen.errcode: gen.errcode.code gen.errcode.doc

.PHONY:	gen.docgo.doc
gen.docgo.doc:
	@echo "===========> Generating missing doc.go for go packages"
	@${ROOT_DIR}/scripts/gendoc.sh

.PHONY:	gen.errcode.code
gen.errcode.code: tools.verify.codegen
	@echo "===========> Generating iam error code go source files"
	@codegen -type=int ${ROOT_DIR}/internal/pkg/code

.PHONY:	gen.errcode.doc
gen.errcode.doc:tools.verify.codegen
	@echo "===========> Generating error code markdown documentation"
	@codegen -type=int -doc \
		-output ${ROOT_DIR}/docs/guide/api/error_code_generated.md ${ROOT_DIR}/internal/pkg/code