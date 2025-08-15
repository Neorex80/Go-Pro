#!/bin/bash

# GitHub Setup Script for Go Projects Toolkit using GitHub CLI

echo "GitHub Setup for Go Projects Toolkit (using GitHub CLI)"
echo "======================================================="

# Check if gh is installed
if ! command -v gh &> /dev/null
then
    echo "GitHub CLI (gh) is not installed. Please install it and try again."
    echo "Visit https://cli.github.com/ for installation instructions."
    exit 1
fi

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

# Ask for repository name
echo ""
echo "Creating GitHub repository..."
read -p "Enter repository name (or press Enter for 'go-projects-toolkit'): " repo_name
if [ -z "$repo_name" ]; then
    repo_name="go-projects-toolkit"
fi

# Ask for repository visibility
echo "Choose repository visibility:"
echo "1) Public (default)"
echo "2) Private"
read -p "Enter choice (1 or 2): " visibility
if [ "$visibility" = "2" ]; then
    visibility_flag="--private"
else
    visibility_flag="--public"
fi

# Create GitHub repository
echo "Creating GitHub repository: $repo_name"
gh repo create "$repo_name" $visibility_flag --clone=false

# Add remote origin
echo "Adding remote origin..."
git remote add origin "https://github.com/$(gh auth status --show-hostname | grep 'Logged in to' | cut -d' ' -f6)/$repo_name.git"

# Push to GitHub
echo "Pushing code to GitHub..."
git push -u origin main

echo ""
echo "Repository created and code pushed successfully!"
echo "Repository URL: https://github.com/$(gh auth status --show-hostname | grep 'Logged in to' | cut -d' ' -f6)/$repo_name"
echo ""
echo "Setup complete!"
