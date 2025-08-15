#!/bin/bash

# Script to set Git user identity and commit changes

echo "Setting Git user identity..."
echo "=========================="

# Set Git user identity
echo "Setting Git user email..."
git config user.email "neore@rex.com"

echo "Setting Git user name..."
git config user.name "Neorex80"

# Commit changes
echo "Committing changes..."
git add .
git commit -m "Initial commit: Go Projects Toolkit"

echo ""
echo "Git user identity set and changes committed successfully!"
echo ""
