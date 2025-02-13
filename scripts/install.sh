#!/bin/bash

set -e

# Install binary
cp enix /usr/bin

# Install configuration files
path=/usr/local/share/enix
mkdir -p $path
cp -r style $path
cp -r colors $path
cp -r filetype $path