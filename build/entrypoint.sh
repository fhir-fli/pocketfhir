#!/bin/sh

# Start the PocketBase server
/app/pocketfhir serve &

# Wait for PocketBase to be ready
while ! nc -z 127.0.0.1 8090; do
  echo "Waiting for PocketBase to be ready..."
  sleep 2
done

echo "PocketBase is ready, starting Caddy..."

# Start the Caddy server
caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
