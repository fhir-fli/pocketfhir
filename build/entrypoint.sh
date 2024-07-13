#!/bin/sh

# Start the PocketBase server
/app/pocketfhir serve --http=0.0.0.0:8090 --dir=/app/pb_data --publicDir=/app/pb_public --hooksDir=/app/pb_hooks
