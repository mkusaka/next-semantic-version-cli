#!/bin/bash

set -e -o pipefail

version="${1:-0.0.4}"

if [[ ! "$version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Given version '${version}' does not match to regex" '^[0-9]+\.[0-9]+\.[0-9]+$' >&2
    exit 1
fi

echo "Start downloading next-semantic-version-cli v${version}"

case "$OSTYPE" in
    linux-*)
        os=linux
        ext=tar.gz
    ;;
    darwin*)
        os=Darwin
        ext=tar.gz
    ;;
    *)
        echo "OS '${OSTYPE}' is not supported." >&2
        exit 1
    ;;
esac

machine="$(uname -m)"
case "$machine" in
    x86_64) arch=x86_64 ;;
    i?86) arch=i386 ;;
    aarch64|arm64) arch=arm64 ;;
    *)
        echo "Could not determine arch from machine hardware name '${machine}'"
        exit 1
    ;;
esac

echo "Detected OS=${os} ext=${ext} arch=${arch}"

file="next-semantic-version-cli_${version}_${os}_${arch}.${ext}"
url="https://github.com/mkusaka/next-semantic-version-cli/releases/download/v${version}/${file}"

echo "Downloading ${url} with curl"

curl -L "${url}" | tar xvz next-semantic-version-cli
exe="$(pwd)/next-semantic-version-cli"

echo "Downloaded and unarchived executable: ${exe}"

echo "Done: $("${exe}" -version)"
