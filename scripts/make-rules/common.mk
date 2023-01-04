SHELL := /bin/bash
# Linux command settings
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs --no-run-if-empty

# 当执行make命令的时候，复制githook scripts到 .git/hooks
COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)