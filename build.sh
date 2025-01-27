#!/bin/bash

APP_NAME="my-todo-cli"
VERSION="1.0.0"
PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

mkdir -p release

for PLATFORM in "${PLATFORMS[@]}"; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_NAME="$APP_NAME-$VERSION-$GOOS-$GOARCH"

  if [ "$GOOS" = "windows" ]; then
    BIN_NAME="$BIN_NAME.exe"
  fi

  echo "Building for $GOOS/$GOARCH..."
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "release/$BIN_NAME"
done
