#!/bin/bash

cat base-docker-compose.yaml > docker-compose.yaml

for ((i=0; i<$1; i++)); do
    echo "

  client_$i:
    build:
      context: ./services/client
      dockerfile: Dockerfile
    container_name: client_$i
    depends_on:
      - server
    environment:
      - AGENCY_ID=$i
      - SERVER_HOST=server
      - SERVER_PORT=5678" >> docker-compose.yaml
done