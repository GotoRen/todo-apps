DOCKER=docker
COMPOSE=docker compose
UP=${COMPOSE} up -d
EXEC=${COMPOSE} exec
LOGS=${COMPOSE} logs -f

DB_NAME=sample



# compose
#===============================================================
all: up

build: ## docker build
	${BUILD}

up: ## docker up
	${UP}

logs: ${LOGS} ## docker logs 

.PHONY: exec/ui
exec/ui: ## exec ui container
	${EXEC} ui sh

.PHONY: exec/api
exec/api: ## exec api container
	${EXEC} api sh

.PHONY: exec/db
exec/db: ## exec db container
	${EXEC} db bash

.PHONY: logs/ui
logs/ui: ## logs ui container
	${LOGS} ui

.PHONY: logs/api
logs/api: ## logs api container
	${LOGS} api

.PHONY: logs/db
logs/db: ## logs db container
	${LOGS} db

mysql: ## db(MySQL) container's MySQL access
	${EXEC} db mysql --defaults-extra-file=/home/access.cnf ${DB_NAME}

stop: ## docker stop
	${COMPOSE} stop

down: ## docker down
	${COMPOSE} down

down/data: ## delete persistent volume
	rm -r ./db/data

down/all: ## delete images, network, containers
	${DOCKER} system prune --all

down/vol: ## delete volumes
	${DOCKER} volume prune



# Makefile config
#===============================================================
help: ## display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
