#!/bin/sh
set -e

echo "==> Loading configuration and downloading dependencies..."
go mod download

echo "==> Compiling go application..."
go build main.go

echo "==> Container listening..."
exec tail -f /dev/null
