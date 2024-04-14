#!/bin/bash

BASE="$(pwd)"
BACKEND="$BASE/backend"
FRONTEND="$BASE/frontend"

# generate database queries

echo "Generating database queries"
sqlc generate

# start frontend

echo "Starting frontend"
cd $FRONTEND && pnpm run dev:wails
