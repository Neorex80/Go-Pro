# Weather API Client

A client that fetches and displays weather data from a public API.

## Features

- Fetches current weather data for a specified city
- Displays temperature, weather condition, humidity, and wind speed
- Simple command-line interface

## Setup

1. Navigate to the project directory:
   ```bash
   cd web-apis/weather-api-client
   ```

2. Get your free API key from [WeatherAPI.com](https://www.weatherapi.com/)
3. Replace `YOUR_API_KEY_HERE` in `main.go` with your actual API key

## Usage

```bash
go run . <city>
```

### Examples

```bash
# Get weather for London
go run . London

# Get weather for New York
go run . "New York"
```

## Building

To build the executable:

```bash
go build -o weather-client
```

Then run it:

```bash
./weather-client <city>
```

## Notes

- The free tier of WeatherAPI allows for 1,000,000 calls per month
- The API key in the code is a placeholder and needs to be replaced with your own key
- Error handling is included for network issues and invalid responses
