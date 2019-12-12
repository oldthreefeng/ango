LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Buildstamp=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Githash=$(shell git rev-parse --short HEAD)"
LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Goversion=$(shell go version)"
BRANCH := $(shell git symbolic-ref HEAD 2>/dev/null | cut -d"/" -f 3)
BUILD := $(shell git rev-parse --short HEAD)
VERSION = $(BRANCH)-$(BUILD)

BASEPATH := $(shell pwd)
CGO_ENABLED = 0
GOCMD = go
GOBUILD = $(GOCMD) build

NAME := ango
DIRNAME := angodir
GOBIN := /usr/local/go/bin/
SRCFILE= main.go
SOFTWARENAME=$(NAME)-$(VERSION)

PLATFORMS := linux darwin

.PHONY: release
release: linux darwin

BUILDDIR:=$(BASEPATH)/../build

.PHONY:Asset
Asset:
	@[ -d $(BUILDDIR) ] || mkdir -p $(BUILDDIR)
	@[ -d $(DIRNAME) ] || mkdir -p $(DIRNAME)

.PHONY: $(PLATFORMS)
$(PLATFORMS): Asset
	@echo "编译" $@
	GOOS=$@ GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -x -o $(NAME) $(SRCFILE)
	cp -f $(NAME) $(DIRNAME)
	cp -f $(NAME) $(GOBIN)
	tar czvf $(BUILDDIR)/$(SOFTWARENAME)-$@-amd64.tar.gz $(DIRNAME)
.PHONY: clean
clean:
	-rm -rf $(NAME)
	-rm -rf $(DIRNAME)
	-rm -rf $(BUILDDIR)
