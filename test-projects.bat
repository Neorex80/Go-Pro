@echo off
echo Testing Go Projects Toolkit

REM Test CLI tools
echo Testing CLI tools...

echo Testing file-renamer...
cd cli-tools/file-renamer
go build .
if %errorlevel% == 0 (
    echo ✓ file-renamer builds successfully
    del file-renamer.exe
) else (
    echo ✗ file-renamer failed to build
)
cd ..\..

echo Testing password-generator...
cd cli-tools/password-generator
go build .
if %errorlevel% == 0 (
    echo ✓ password-generator builds successfully
    del password-generator.exe
) else (
    echo ✗ password-generator failed to build
)
cd ..\..

echo Testing todo-cli...
cd cli-tools/todo-cli
go build .
if %errorlevel% == 0 (
    echo ✓ todo-cli builds successfully
    del todo-cli.exe
) else (
    echo ✗ todo-cli failed to build
)
cd ..\..

REM Test web APIs
echo Testing web APIs...

echo Testing simple-rest-api...
cd web-apis/simple-rest-api
go mod tidy
go build .
if %errorlevel% == 0 (
    echo ✓ simple-rest-api builds successfully
    del simple-rest-api.exe
) else (
    echo ✗ simple-rest-api failed to build
)
cd ..\..

echo Testing weather-api-client...
cd web-apis/weather-api-client
go build .
if %errorlevel% == 0 (
    echo ✓ weather-api-client builds successfully
    del weather-api-client.exe
) else (
    echo ✗ weather-api-client failed to build
)
cd ..\..

REM Test automation scripts
echo Testing automation scripts...

echo Testing file-organizer...
cd automation-scripts/file-organizer
go build .
if %errorlevel% == 0 (
    echo ✓ file-organizer builds successfully
    del file-organizer.exe
) else (
    echo ✗ file-organizer failed to build
)
cd ..\..

echo Testing system-monitor...
cd automation-scripts/system-monitor
go build .
if %errorlevel% == 0 (
    echo ✓ system-monitor builds successfully
    del system-monitor.exe
) else (
    echo ✗ system-monitor failed to build
)
cd ..\..

echo Testing complete!
