#!/bin/bash

if [ -f .env ]; then
    source .env
fi

# two commands are available: up and down
# before using down, think and plan it
# in production, it's just no no
# also, do think about migrate up only
# 'down' command is available/applicable for the first-time development.
echo -n "Enter command (up/down): "
read cmd

cd sql/schema
goose postgres $DB_URL $cmd