#!/usr/bin/env bash

export PROJECT_ID=$1
export BASE64ENCODED_GCP_PROVIDER_CREDS=$(base64 service-account.json | tr -d "\n")
sed "s/BASE64ENCODED_GCP_PROVIDER_CREDS/$BASE64ENCODED_GCP_PROVIDER_CREDS/g;s/PROJECT_ID/$PROJECT_ID/g" gcp-provider.yaml | kubectl apply -f -
