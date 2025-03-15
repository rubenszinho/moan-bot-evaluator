#!/bin/bash

echo "Building the Discord Moan Bot..."
go mod tidy

go build -o bin/discord-moan-bot cmd/main.go

echo "Build complete. Run './bin/discord-moan-bot' to start the bot."
