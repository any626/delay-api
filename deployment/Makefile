include config.mk

run: build up

build:
	docker-compose build

up:
	docker-compose up -d

kill:
	docker rm -f delay-engine 

help:
	@printf "\nUsage:\n"
	@printf "$$HELP_USAGE" | column -t -s $$'\t'
	@echo
