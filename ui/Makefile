YARN=yarn
INSTALL=$(YARN) install
RUN=$(YARN) run
CACHE=$(YARN) cache
GLOBAL=$(YARN) global



install: ## module install
	$(INSTALL)

run: ## yarn run
	$(RUN) start

list: ## package list
	$(YARN) list --depth=0

upgrade: ## all packages upgrade
	$(GLOBAL) upgrade

clean: ## cache clean
	$(CACHE) clean

cache: ## cache list
	$(CACHE) list

cache/dir: ## cache list (directory)
	$(CACHE) dir

config: ## config list
	$(YARN) config list

env: ## Display the environment variable
	$(RUN) env



# Makefile config
#===============================================================
help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
