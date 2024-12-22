#!/usr/bin/env bash

set -e

#sudo echo "127.0.0.1 slave.a.com" >> /etc/hosts
go test -tags=e2e -race -failfast -count=1 -timeout=30m ./e2e
