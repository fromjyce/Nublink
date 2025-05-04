#!/bin/bash

CERT_DIR="$HOME/.nublink"
mkdir -p "$CERT_DIR"

openssl req -x509 -newkey rsa:4096 -keyout "$CERT_DIR/key.pem" -out "$CERT_DIR/cert.pem" \
  -days 365 -nodes -subj "/CN=localhost"

echo "Certificates generated in $CERT_DIR"