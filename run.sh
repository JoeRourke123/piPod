#!/bin/bash

PIPE_NAME="pipod-pipe"
PIPE_PATH="/tmp/$PIPE_NAME"

# Function to clean up the pipe and exit
cleanup() {
    echo "Cleaning up..."
    rm -f "$PIPE_PATH"
    exit 0
}

# Set up the pipe if it doesn't already exist
if [[ ! -p "$PIPE_PATH" ]]; then
    echo "Setting up the pipe..."
    mkfifo "$PIPE_PATH"
fi

# Start the pipe listener
(while true; do
    eval "$(cat $PIPE_PATH)"
done) &
PIPE_PID=$!
trap cleanup SIGINT SIGTERM

# Run docker-compose up
docker compose up -d

# Wait for the script to be killed
wait $PIPE_PID