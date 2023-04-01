#!/bin/bash

find . -maxdepth 1 -type d -not -path '*/\.*' | sort | fzf --preview 'echo "Execute main.go in: {}"' | xargs -I {} bash -c 'cd "{}" && go run main.go'
