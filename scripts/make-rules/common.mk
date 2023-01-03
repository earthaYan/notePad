# 当执行make命令的时候，复制githook scripts
COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)