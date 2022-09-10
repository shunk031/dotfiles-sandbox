DOTPATH := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))

deploy: ## Create symlink to home directory
	@echo '==> Start to deploy dotfiles to home directory.'
	@DOTPATH=$(DOTPATH) go run -tags linux,amd64,server ./cmd/dotfiles/ deploy

init: ## Setup environment settings
	@echo '==> Start to init dotfiles'
	@DOTPATH=$(DOTPATH) go run -tags linux,amd64,server ./cmd/dotfiles/ init

install:
	deploy
	init
