#!/bin/bash

TARGET_DIR="/usr/local/bin"
BIN_NAME="my-todo-cli"  # Name of the binary inside the .tar.gz
VERSION="1.0.0"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architectures
case "$ARCH" in
  "x86_64") ARCH="amd64" ;;
  "arm64") ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Download URL
DOWNLOAD_URL="https://github.com/Shivam583-hue/kanban-board-cli-/releases/download/v1.0.0/my-todo-cli-$VERSION-$OS-$ARCH.tar.gz"

echo "Installing $BIN_NAME..."
curl -L $DOWNLOAD_URL | tar -xz -C $TARGET_DIR $BIN_NAME
chmod +x "$TARGET_DIR/$BIN_NAME"
echo "Done! Run '$BIN_NAME' to start."
