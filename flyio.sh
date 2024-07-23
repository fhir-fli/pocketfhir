#!/bin/bash

# Path to your .env file
ENV_FILE="./config/.env.flyio"

# Read the .env file line by line
while IFS= read -r line || [ -n "$line" ]; do
  # Ignore empty lines and lines starting with #
  if [ -z "$line" ] || [ "${line:0:1}" == "#" ]; then
    continue
  fi
  # Set each secret
  flyctl secrets set $line
done < "$ENV_FILE"

# Verify that the environment variables are set
flyctl secrets list

# Deploy the application
flyctl deploy
