# URL Shortener with Handler Options
## This Go program serves as a URL shortener with multiple handler options based on different data sources.

## Overview
### The program utilizes Go's net/http package to create a web server that responds with redirections based on short URLs provided in different formats: JSON, YAML, and BoltDB.

## Dependencies
github.com/boltdb/bolt: A pure Go key/value store database.
## Installation
Ensure you have Go installed. If not, follow instructions on golang.org.
Fetch the necessary dependencies using Go Modules:
```bash
go mod tidy
```
## Usage
The program supports the following command-line flags:

- json: Path to a JSON file containing short URL mappings.
- yaml: Path to a YAML file containing short URL mappings.
- bolt: Path to a BoltDB file containing short URL mappings.


## Example
To start the server with JSON file-based URL mappings:
```bash 
go run main.go -json=path/to/your/json/file.json
```
Replace path/to/your/json/file.json with the actual path to your JSON file.

# Handlers
Map Handler: Handles basic / route and forwards other routes to specified short URLs.
JSON Handler: Reads short URL mappings from a JSON file.
YAML Handler: Reads short URL mappings from a YAML file.
BoltDB Handler: Reads short URL mappings from a BoltDB database file.

# Files
main.go: Contains the main server logic and routing setup.
handlers.go: Defines functions for different types of URL handlers (JSON, YAML, BoltDB).
db.go: Implements BoltDB-specific functionality.

# Notes
Ensure the server is run with appropriate file paths specified using the command-line flags (-json, -yaml, -bolt).
Each handler option (JSON, YAML, BoltDB) provides flexibility in storing and retrieving short URL mappings.
