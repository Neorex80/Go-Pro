@echo off
REM GitHub Setup Script for Go Projects Toolkit using GitHub CLI

echo GitHub Setup for Go Projects Toolkit (using GitHub CLI)
echo =======================================================

REM Check if gh is installed
gh --version >nul 2>&1
if %errorlevel% neq 0 (
    echo GitHub CLI (gh) is not installed. Please install it and try again.
    echo Visit https://cli.github.com/ for installation instructions.
    exit /b 1
)

REM Check if git is installed
git --version >nul 2>&1
if %errorlevel% neq 0 (
    echo Git is not installed. Please install Git and try again.
    exit /b 1
)

REM Initialize Git repository
echo Initializing Git repository...
git init

REM Add all files
echo Adding files to Git...
git add .

REM Make initial commit
echo Making initial commit...
git commit -m "Initial commit: Go Projects Toolkit"

REM Ask for repository name
echo.
echo Creating GitHub repository...
set /p repo_name="Enter repository name (or press Enter for 'go-projects-toolkit'): "
if "%repo_name%"=="" set repo_name=go-projects-toolkit

REM Ask for repository visibility
echo Choose repository visibility:
echo 1) Public (default)
echo 2) Private
set /p visibility="Enter choice (1 or 2): "
if "%visibility%"=="2" (
    set visibility_flag=--private
) else (
    set visibility_flag=--public
)

REM Create GitHub repository
echo Creating GitHub repository: %repo_name%
gh repo create %repo_name% %visibility_flag% --clone=false

REM Get GitHub username
for /f "tokens=6" %%a in ('gh auth status --show-hostname ^| findstr "Logged in to"') do set github_user=%%a

REM Add remote origin
echo Adding remote origin...
git remote add origin https://github.com/%github_user%/%repo_name%.git

REM Push to GitHub
echo Pushing code to GitHub...
git push -u origin main

echo.
echo Repository created and code pushed successfully!
echo Repository URL: https://github.com/%github_user%/%repo_name%
echo.
echo Setup complete!
