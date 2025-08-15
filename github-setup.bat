@echo off
REM GitHub Setup Script for Go Projects Toolkit

echo GitHub Setup for Go Projects Toolkit
echo =====================================

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

echo.
echo Repository initialized successfully!
echo.
echo Next steps:
echo 1. Create a new repository on GitHub (do not initialize with README)
echo 2. Copy the repository URL
echo 3. Add the remote origin:
echo    git remote add origin ^<your-repository-url^>
echo 4. Push the code to GitHub:
echo    git push -u origin main
echo.
echo Setup complete!
