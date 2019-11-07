SHELL:=/bin/bash

export DELAY_ENGINE_CONTAINER=delay-engine
export DELAY_ENGINE_IMAGE=delay-engine-image
export DELAY_ENGINE_PORT=8080

define HELP_USAGE
  make or make run \t Builds the docker images and runs the containers. \n
  make build \t Builds the docker images. \n
  make up \t Brings up the docker containers. \n
  make kill \t Kills the docker containers. \n
endef
export HELP_USAGE
