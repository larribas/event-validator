#!/usr/bin/env bash

docker_ip="127.0.0.1"
if boot2docker; then
    docker_ip=$(boot2docker ip)
fi

# Set up infrastructure
redis_container=$(docker run -d -p 6379:6379 redis:2.8)

export EV_TEST_REDIS_HOST=$docker_ip
export EV_TEST_REDIS_PORT=6379

sleep 5

# Launch tests
go test ./...
test_outcome=$?

# Clean up
docker kill $redis_container

# Exit with the status of go test
exit $test_outcome