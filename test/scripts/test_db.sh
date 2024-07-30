#!/bin/bash

# Docker container'ını çalıştır
docker run --name kirmac -e POSTGRES_PASSWORD=kirmac123 -e POSTGRES_USER=kirmac -p 5433:5432 -d postgresql:latest
echo "Waiting for postgres to start"
sleep 10

# Yeni veritabanı oluştur
docker exec -it kirmac psql -U kirmac -d postgresql -c "CREATE DATABASE kirmac_site;"
sleep 5
echo "Database created successfully"

# Yeni veritabanında tablo oluştur
docker exec -it kirmac psql -U kirmac -d kirmac_site -c "
CREATE TABLE IF NOT EXISTS properties
(
  ID BIGINT PRIMARY KEY,
  Location VARCHAR(255),
  Price INT,
  Title VARCHAR(255),
  Description TEXT,
  Bedrooms INT,
  Bathrooms INT,
  SquareFeet INT,
  AgentName VARCHAR(255),
  AgentTitle VARCHAR(255),
  ImageURLs TEXT
);"
sleep 5
echo "Table created successfully"