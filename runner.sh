#!/bin/bash

find . -type d -not -path '*/\.*' | fzf --preview 'echo "Execute main.go in: {}"' | xargs -I {} bash -c 'cd "{}" && go run main.go'
