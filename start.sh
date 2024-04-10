#!/bin/sh

# Start Redis in the background
redis-server --daemonize yes

# Start your Go application
./go-gemma
