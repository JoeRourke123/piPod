#!/bin/bash

# Check if we are in the 'ui' directory, if not move into it
cd /ui || (echo "Failed to change directory to /ui" && exit 1)

# Check if node_modules directory exists to avoid reinstalling dependencies
if [ ! -d "node_modules" ]; then
  echo "Installing dependencies..."
  yarn install
else
  echo "Dependencies already installed."
fi

# Check if the build directory exists to avoid rebuilding
if [ ! -d "build" ]; then
  echo "Building the project..."
  yarn build
else
  echo "Build already exists."
fi

# Check the ENV environment variable and run the appropriate command
if [ "$ENV" == "dev" ]; then
  echo "Running in development mode..."
  yarn start
elif [ "$ENV" == "prod" ]; then
  echo "Running in production mode..."
  if ! command -v serve &> /dev/null; then
    echo "Installing serve..."
    npm install -g serve
  fi
  serve -s build
fi