#!/bin/bash

# Set variables
projectId="fhirfli-413723"
location="us-central1"
repository="containers"
projectName="pocketfhir"
gcloud config set project $projectId

# only needed the first time
# gcloud auth login

# Define registry location
registryLocation="$location-docker.pkg.dev/$projectId/$repository/$projectName"

# Build the Docker image
docker build -t $projectName -f build/Dockerfile .

# Tag the Docker image
docker tag $projectName $registryLocation

# Push the tagged image to the Artifact Registry
docker push $registryLocation

# Deploy on Google Cloud Run
gcloud run deploy $projectName --image $registryLocation --platform managed --region $location --allow-unauthenticated --port 8080
