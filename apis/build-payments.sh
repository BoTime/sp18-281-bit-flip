#!/bin/bash

USERNAME="bbamsch"
CONTAINER="cmpe281-payment"
TAG="latest"
DOCKERFILE="Dockerfile-Payments"

POSITIONAL=()
while [[ $# -gt 0 ]]
do
  key="$1"

  # Parse Arguments
  case $key in
    -u|--username)
      USERNAME="$2"
      shift
      shift
      ;;
    -c|--container)
      CONTAINER="$2"
      shift
      shift
      ;;
    -t|--tag)
      TAG="$2"
      shift
      shift
      ;;
    -d|--dockerfile)
      DOCKERFILE="$2"
      shift
      shift
      ;;
    -g|--gopath)
      GOPATH="$2"
      shift
      shift
      ;;
    *) # positional arg
      POSITIONAL+=("$1")
      shift
      ;;
  esac
done

docker build -t "$USERNAME/$CONTAINER:$TAG" --file "$DOCKERFILE" .
