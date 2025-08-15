#!/bin/bash

echo "Testing Go Projects Toolkit"

# Test CLI tools
echo "Testing CLI tools..."

echo "Testing file-renamer..."
cd cli-tools/file-renamer
go build .
if [ $? -eq 0 ]; then
    echo "✓ file-renamer builds successfully"
    rm -f file-renamer
else
    echo "✗ file-renamer failed to build"
fi
cd ../..

echo "Testing password-generator..."
cd cli-tools/password-generator
go build .
if [ $? -eq 0 ]; then
    echo "✓ password-generator builds successfully"
    rm -f password-generator
else
    echo "✗ password-generator failed to build"
fi
cd ../..

echo "Testing todo-cli..."
cd cli-tools/todo-cli
go build .
if [ $? -eq 0 ]; then
    echo "✓ todo-cli builds successfully"
    rm -f todo-cli
else
    echo "✗ todo-cli failed to build"
fi
cd ../..

# Test web APIs
echo "Testing web APIs..."

echo "Testing simple-rest-api..."
cd web-apis/simple-rest-api
go mod tidy
go build .
if [ $? -eq 0 ]; then
    echo "✓ simple-rest-api builds successfully"
    rm -f simple-rest-api
else
    echo "✗ simple-rest-api failed to build"
fi
cd ../..

echo "Testing weather-api-client..."
cd web-apis/weather-api-client
go build .
if [ $? -eq 0 ]; then
    echo "✓ weather-api-client builds successfully"
    rm -f weather-api-client
else
    echo "✗ weather-api-client failed to build"
fi
cd ../..

# Test automation scripts
echo "Testing automation scripts..."

echo "Testing file-organizer..."
cd automation-scripts/file-organizer
go build .
if [ $? -eq 0 ]; then
    echo "✓ file-organizer builds successfully"
    rm -f file-organizer
else
    echo "✗ file-organizer failed to build"
fi
cd ../..

echo "Testing system-monitor..."
cd automation-scripts/system-monitor
go build .
if [ $? -eq 0 ]; then
    echo "✓ system-monitor builds successfully"
    rm -f system-monitor
else
    echo "✗ system-monitor failed to build"
fi
cd ../..

echo "Testing complete!"
