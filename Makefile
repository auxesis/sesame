.DEFAULT_GOAL := run
GOPATH := $(shell pwd)/_vendor

run:
	echo $$GOPATH
	go run -x sesame.go
