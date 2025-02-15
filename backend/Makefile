PROJECT=msg

# 系统工具
MKDIR=mkdir
CP=cp
MV=mv
RM=rm
MAKE=make
SED=sed
DATE=date
ARCH=arch
# 其他工具
GIT=git
GO=go
GOFMT=gofmt

# 编译信息
BRANCH=$(shell $(GIT) rev-parse --abbrev-ref HEAD)
VERSION=$(shell $(GIT) describe --tags --always | $(SED) 's/^v//')
COMMIT=$(shell $(GIT) rev-parse --verify HEAD)
BUILD_TIME=$(shell $(DATE) +"%Y-%m-%d %H:%M:%S %z")
GOARCH=$(shell $(ARCH) | $(SED) 's/x86_64/amd64/' | $(SED) 's/aarch64/arm64/')
GO_FLAGS=-ldflags "-s -w -X '$(PROJECT)/cmd/def.Branch=$(BRANCH)' -X '$(PROJECT)/cmd/def.Version=$(VERSION)' -X '$(PROJECT)/cmd/def.Commit=$(COMMIT)' -X '$(PROJECT)/cmd/def.BuildTime=$(BUILD_TIME)'"

.PHONY: format build release example migrate dev count

# 格式化项目
format:
	@$(GO) mod tidy
	@$(GOFMT) -s -w .
	@echo "Version: $(VERSION)-$(BRANCH)_$(COMMIT) ($(BUILD_TIME))"

# 编译项目
build:
	$(MAKE) format
	CGO_ENABLED=0 $(GO) build $(GO_FLAGS) -o build/$(PROJECT) main.go

# 编译为目标架构发行版
release:
	$(MAKE) format
	CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) $(GO) build $(GO_FLAGS) -o build/$(PROJECT)_$(GOARCH) main.go

# 生成默认配置文件
example:
	$(MAKE) build
	- $(MKDIR) build/data
	cd build && ./$(PROJECT) example -c > ./data/config.toml

# 创建、初始化数据库
migrate:
	$(MAKE) build
	- $(MKDIR) build/data
	-# $(RM) build/data/$(PROJECT).db
	cd build && ./$(PROJECT) migrate

# 调试运行
dev:
	$(MAKE) build
	cd build && ./$(PROJECT) migrate && ./$(PROJECT) server

