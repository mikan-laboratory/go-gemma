#!/bin/sh

# Start Redis in the background with persistence configuration
redis-server --daemonize yes --save 20 1 --dir data

# Start your Go application
./go-gemma
