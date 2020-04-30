#!/bin/bash

reflex -r '(\.go$|go\.mod|\.tmpl)' -s go run main.go
