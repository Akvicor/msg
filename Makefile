PROJECT=msg

# 系统工具
MKDIR=/usr/bin/mkdir
CP=/usr/bin/cp
MV=/usr/bin/mv
RM=/usr/bin/rm
MAKE=/usr/bin/make
SED=/usr/bin/sed
DATE=/usr/bin/date
ARCH=/usr/bin/arch
# 其他工具
GIT=/usr/bin/git
GO=/usr/bin/go
GOFMT=/usr/bin/gofmt
SCC=/kayuki/vivc/bin/scc

.PHONY: build count

# 编译项目
build:
	$(MAKE) -C frontend build
	$(MAKE) -C backend build

# 统计代码行数
count:
	@$(SCC) ./

