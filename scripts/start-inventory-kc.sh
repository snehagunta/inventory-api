#!/bin/bash
set -e
# Function to check if a command is available
source ./scripts/check_docker_podman.sh
NETWORK_CHECK=$(${DOCKER} network ls --filter name=kessel --format json)
if [[ -z "${NETWORK_CHECK}" ]]; then ${DOCKER} network create kessel; fi
${DOCKER} compose --env-file ./scripts/.env -f ./docker-compose-sso.yaml up -d --build
