#!/usr/bin/env bash

docker run --rm \
	-v $(pwd):/src \
	-v /var/run/docker.sock:/var/run/docker.sock \
	centurylink/golang-builder \
	splorenzoarribas/event_validator:latest