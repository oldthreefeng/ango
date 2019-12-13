LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Buildstamp=$(shell date -u '+%Y-%m-%d %H:%M:%S %Z')"
LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Githash=$(shell git rev-parse --short HEAD)"
LDFLAGS += -X "github.com/oldthreefeng/ango/cmd.Goversion=$(shell go version)"
BRANCH := $(shell git symbolic-ref HEAD 2>/dev/null | cut -d"/" -f 3)
BUILD := $(shell git rev-parse --short HEAD)
VERSION = $(BRANCH)-$(BUILD)

BASEPATH := $(shell pwd)
CGO_ENABLED = 0
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOMOD = $(GOCMD) mod

NAME := ango
DIRNAME := bin
GOBIN := /usr/local/go/bin/
SRCFILE= main.go
SOFTWARENAME=$(NAME)-$(VERSION)

PLATFORMS := linux darwin

.PHONY: test
test:
	$(GOTEST) -v ./...
.PHONY: run
run:
	$(GOBUILD) -ldflags '$(LDFLAGS)'  -o $(NAME) $(SRCFILE) 
	./$(NAME)
.PHONY: deps
deps:
	$(GOMOD) tidy
	$(GOMOD) download

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
	GOOS=$@ GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on GOPROXY=https://goproxy.cn $(GOBUILD) -ldflags '$(LDFLAGS)'  -o $(NAME) $(SRCFILE)
	cp -f $(NAME) $(DIRNAME)
	cp -f $(NAME) $(GOBIN)
	tar czvf $(BUILDDIR)/$(SOFTWARENAME)-$@-amd64.tar.gz $(DIRNAME)
.PHONY: clean
clean:
	-rm -rf $(NAME)
	-rm -rf $(DIRNAME)
	-rm -rf $(BUILDDIR)
