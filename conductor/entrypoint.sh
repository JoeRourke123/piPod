#!/bin/bash

# Set current working directory to ./conductor only if not already in conductor
cd /conductor || (echo "Failed to change directory to /conductor" && exit 1)

# Check if tidal-dl is installed
if ! command -v tidal-dl &> /dev/null
then
    echo "tidal-dl not found, installing..."

    # Install pip if it is not already installed
    if ! command -v pip &> /dev/null
    then
        echo "pip not found, installing..."
        (apt-get update && apt-get install -y python3-pip) || (echo "Failed to install pip" && exit 1)
    else
        echo "pip is already installed"
    fi

    python3 -m pip install tidal-dl --break-system-packages --upgrade || (echo "Failed to install tidal-dl" && exit 1)
else
    echo "tidal-dl is already installed"
fi

# Build and run the Go project based on the ENV environment variable
if [ "$ENV" == "dev" ]; then
    # Install air
    if ! command -v air &> /dev/null
    then
        echo "air not found, installing..."
        go install github.com/air-verse/air@latest || (echo "Failed to install air" && exit 1)
    else
        echo "air is already installed"
    fi

    echo "Running in development mode with air..."
    air -c .air.toml
elif [ "$ENV" == "prod" ]; then
    echo "Building the Go project..."
    go build -o conductor || (echo "Failed to build the Go project" && exit 1)

    echo "Running the Go project..."
    ./conductor
else
    echo "Unknown environment: $ENV"
    exit 1
fi