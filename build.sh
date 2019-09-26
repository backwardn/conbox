#!/bin/bash

export GO111MODULE=on

msg() {
	echo 2>&1 "$0": $@
}

if [ -n "$1" ]; then
	lint=1
else
	lint=
fi

if [ -n "$lint" ]; then
	msg linting
else
	msg NOT linting
fi

if [ -n "$lint" ]; then
	hash gosimple 2>/dev/null && gosimple ./applets
	hash golint 2>/dev/null && golint ./applets
	hash staticcheck 2>/dev/null && staticcheck ./applets
fi

gofmt -s -w ./applets
go fix ./applets/...
go test ./applets/...
go install -v ./applets/...

if [ -n "$lint" ]; then
	hash gosimple 2>/dev/null && gosimple ./common
	hash golint 2>/dev/null && golint ./common
	hash staticcheck 2>/dev/null && staticcheck ./common
fi

gofmt -s -w ./common
go fix ./common
go test ./common
go install -v ./common

if [ -n "$lint" ]; then
	hash gosimple 2>/dev/null && gosimple ./conbox
	hash golint 2>/dev/null && golint ./conbox
	hash staticcheck 2>/dev/null && staticcheck ./conbox
fi

gofmt -s -w ./conbox
go fix ./conbox
go test ./conbox
go install -v ./conbox

