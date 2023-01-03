# 当执行make命令的时候，复制githook scripts到 .git/hooks
COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)