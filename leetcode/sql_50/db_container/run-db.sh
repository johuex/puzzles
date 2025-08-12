#!/bin/bash

# NOTE: don't launch, it's just example of commands

# build image
docker build -t postgres-17-alpine-puzzles .

# create volume for daat reusage
docker volume create pgdata-puzzles

# only first launch
docker run -d \
 --name postgres-puzzles \
 -e POSTGRES_USER=postgres \
 -e POSTGRES_PASSWORD=123456 \
 -e POSTGRES_DB_NAME=demo \
 -e POSTGRES_SCHEMA_NAME=public \
 -v pgdata-puzzles:/var/lib/postgresql/data \
 -p 5432:5432 \
 postgres-17-alpine-puzzles


# if container exist
docker start postgres-puzzles
docker stop postgres-puzzles