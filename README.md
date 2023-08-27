# GO-WEATHER

## Description

GO-WEATHER is a command-line tool that provides weather forecasts for a specified city. It fetches weather data from an external API and displays the current weather conditions and hourly forecasts.

## Features

- Fetches weather data using an external API.
- Displays current weather conditions and hourly forecasts.
- Supports customization of the number of forecasted hours.
- Uses color coding to highlight weather conditions.

## Installation

You can install GO-WEATHER using the following command:
```sh
go install
```
On mac-os you can install it globally using commands:
```sh
go build
mv go-weather /usr/local/bin/
```
## Usage
To use go-weather, simply run the command:
```sh
go-weather
```
If you want to specify city name and number of forecasted hours, you can run command:
```sh
go-weather [city] [number of houers(1-24)]
```
For example:
```sh
go-weather London 8
```
You can also run app using commands:
```sh
go run main.go
```
Or
```sh
go run main.go [city] [number of houers(1-24)]
```