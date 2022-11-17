#!/usr/bin/bash
go install github.com/cosmtrek/air@latest

mv $(go env GOPATH)/bin/air ~/.air

./home/.air ./server.
