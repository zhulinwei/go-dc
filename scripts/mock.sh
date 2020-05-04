#!/usr/bin/env bash

CGO_ENABLED=0 mockgen -destination pkg/dao/mock/user_mock.go -source pkg/dao/user.go
CGO_ENABLED=0 mockgen -destination pkg/service/mock/user_mock.go -source pkg/service/user.go
