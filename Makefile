# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOFMT=gofmt
GOLINT=golint
BINARY_NAME=$(shell node -p "require('./package.json').name")
PKG_LINUX=bin/${BINARY_NAME}-linux
VERSION := $(shell node -p "require('./package.json').version")
DESCRIPTION := $(shell node -p "require('./package.json').description")
HOMEPAGE := $(shell node -p "require('./package.json').homepage")
AUTHOR=stephendltg
NODE=v14.17.5
NVM=v0.38.0
DENO=1.13.0

all: deps tool build-app

install: 
	@echo "Installing project ${BINARY_NAME}..."
	. ${NVM_DIR}/nvm.sh && nvm install ${NODE} 
	nvm use ${NODE}
	npm install
	curl -fsSL https://deno.land/x/install/install.sh | sh
	deno upgrade --version ${DENO}

dev:
	$(GORUN) main.go -debug

build-app:
	$(GOBUILD) -v -race main.go

build-deb:
	mkdir -p $(PKG_LINUX)/DEBIAN
	mkdir -p $(PKG_LINUX)/usr/bin/
	echo "Package: $(BINARY_NAME)" > $(PKG_LINUX)/DEBIAN/control
	echo "Version: $(VERSION)" >> $(PKG_LINUX)/DEBIAN/control
	echo "Section: custom" >> $(PKG_LINUX)/DEBIAN/control
	echo "Architecture: all" >> $(PKG_LINUX)/DEBIAN/control
	echo "Essential: no" >> $(PKG_LINUX)/DEBIAN/control
	echo "Depends: libwebkit2gtk-4.0-dev" >> $(PKG_LINUX)/DEBIAN/control
	echo "Maintainer: $(AUTHOR)" >> $(PKG_LINUX)/DEBIAN/control
	echo "Description: $(DESCRIPTION)" >> $(PKG_LINUX)/DEBIAN/control
	echo "Homepage: ${HOMEPAGE}" >> $(PKG_LINUX)/DEBIAN/control
	GOOS=linux $(GOBUILD) -v -ldflags="-X 'main.Title=${BINARY}'" -o $(PKG_LINUX)/usr/bin/$(BINARY_NAME) main.go
	sudo dpkg-deb --build $(PKG_LINUX)
	rm -r $(PKG_LINUX)/*
	rmdir $(PKG_LINUX)

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -ldflags="-X 'main.Title=${BINARY}'" -o bin/$(BINARY_NAME)-linux-amd64 main.go

build-rasp:
	GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -v -ldflags="-X 'main.Title=${BINARY}'" -o bin/$(BINARY_NAME)-rasp main.go

build-darwin:
	mkdir -p bin/$(BINARY_NAME).app/Contents/MacOS
	mkdir -p bin/$(BINARY_NAME).app/Contents/Resources
	cp assets/Info.plist bin/$(BINARY_NAME).app/Contents/Info.plist
	cp assets/icon.icns bin/$(BINARY_NAME).app/Contents/Resources
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -v -ldflags="-X 'main.Title=${BINARY}'" -o bin/$(BINARY_NAME).app/Contents/MacOS/$(BINARY_NAME) main.go

build-win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-H windowsgui -X 'main.Title=${BINARY}'" -v -o bin/$(BINARY_NAME)-win32-amd64.exe main.go

tool:
	$(GOVET) ./...; true
	$(GOFMT) -w main.go

clean:
	go clean -i main.go
	rm -f $(BINARY_NAME)

deps:
	go mod vendor
	# go mod tidy
	go mod verify

nvm:
	curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/${NVM}/install.sh | bash

env:
	@echo "Package ..."
	@echo App: ${BINARY_NAME}
	@echo Version: ${VERSION}
	@echo Description: ${DESCRIPTION}
	@echo Author: ${AUTHOR}
	@echo Homepage: ${HOMEPAGE}
	@echo Node version: ${NODE}
	@echo Deno version: ${DENO}

help:
	@echo "echo package env"
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make clean: remove object files and cached files"
	@echo "make nvm: insall nvm"
	@echo "make pre-install: Pre install nodejs"
	@echo "make deps: get the deployment tools"
