#!/bin/sh

mockgen -package user -self_package github.com/laster18/poi/api/src/domain/user -source src/domain/user/repository.go -destination src/domain/user/repository_mock.go
mockgen -package room -self_package github.com/laster18/poi/api/src/domain/room -source src/domain/room/repository.go -destination src/domain/room/repository_mock.go
