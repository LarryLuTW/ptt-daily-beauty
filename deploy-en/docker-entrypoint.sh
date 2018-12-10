#!/bin/bash

gcloud auth activate-service-account --key-file /src/deploy-en/credential.json
gcloud config set project daily-beauty-209105

exec "$@"