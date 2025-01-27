#!/bin/bash

TARGET_DIR="/usr/local/bin"
BIN_NAME="kanban-board-cli"  # Match your actual binary name

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architectures
case "$ARCH" in
  "x86_64") ARCH="amd64" ;;
  "arm64") ARCH="arm64" ;;  # For Apple Silicon
esac

# Updated URL
DOWNLOAD_URL="https://github.com/Shivam583-hue/kanban-board-cli-/releases/download/v1.0.0/kanban-board-cli-$OS-$ARCH.tar.gz"

echo "Installing $BIN_NAME..."
curl -L $DOWNLOAD_URL | tar -xz -C $TARGET_DIR $BIN_NAME
chmod +x "$TARGET_DIR/$BIN_NAME"
echo "Done! Run '$BIN_NAME' to start."
