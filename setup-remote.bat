@echo off
REM Setup script to initialize Git repository and add remote origin

echo Setting up Git repository...
echo ============================

REM Initialize Git repository
echo Initializing Git repository...
git init

REM Add all files
echo Adding files to Git...
git add .

REM Make initial commit
echo Making initial commit...
git commit -m "Initial commit: Go Projects Toolkit"

REM Add remote origin
echo Adding remote origin...
git remote add origin https://github.com/Neorex80/Go-Pro.git

echo.
echo Git repository initialized and remote origin added successfully!
echo.
echo To push the code to GitHub, run:
echo   git push -u origin main
echo.
