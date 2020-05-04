#!/usr/bin/env bash

CGO_ENABLED=0 go test github.com/zhulinwei/go-dc/pkg/service
CGO_ENABLED=0 go test github.com/zhulinwei/go-dc/pkg/controller
