#!/bin/bash 
go test -coverprofile=coverage.out ./...
#:CoverageLoad
#:CoverageShow
