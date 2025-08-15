#!/bin/bash

# GitHub Setup Script for Go Projects Toolkit

echo "GitHub Setup for Go Projects Toolkit"
echo "====================================="

# Check if git is installed
if ! command -v git &> /dev/null
then
    echo "Git is not installed. Please install Git and try again."
    exit 1
fi

# Initialize Git repository
echo "Initializing Git repository..."
git init

# Add all files
echo "Adding files to Git..."
git add .

# Make initial commit
echo "Making initial commit..."
git commit -m "Initial commit: Go Projects Toolkit"

echo ""
echo "Repository initialized successfully!"
echo ""
echo "Next steps:"
echo "1. Create a new repository on GitHub (do not initialize with README)"
echo "2. Copy the repository URL"
echo "3. Add the remote origin:"
echo "   git remote add origin <your-repository-url>"
echo "4. Push the code to GitHub:"
echo "   git push -u origin main"
echo ""
echo "Setup complete!"
