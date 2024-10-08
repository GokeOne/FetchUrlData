# URL Scanner & Sensitive Data Finder

This is a simple command-line tool built in Go that scans a specified domain for URLs using the `gau` (Get All URLs) tool, checks their HTTP status, and searches for sensitive patterns such as passwords, credentials, and usernames within the content of valid URLs (HTTP status 200).

## Features
- Retrieves URLs from a given domain using `gau`.
- Filters out URLs that return HTTP status 200.
- Scans the content of valid URLs for sensitive patterns such as:
  - Passwords
  - Usernames
  - Credentials
- Outputs any sensitive data findings into a file `pattern_matches.txt`.

## Requirements

Before running this program, make sure you have the following tools installed:

1. [Go](https://golang.org/doc/install) (version 1.16+ recommended)
2. [curl](https://curl.se/download.html)
3. [gau](https://github.com/lc/gau) - Get All URLs, (If it is Windows, put it in the path)

## Installation

1. Clone this repository:

`git clone https://github.com/GokeOne/FetchUrlData`
`cd FetchUrlData`

2. Ensure all required tools are installed:
`go install github.com/lc/gau/v2/cmd/gau@latest`

3. Build the Go project(optional)

`go build -o FetchUrlData`


# Usage

You can run the program using go run, or if you have built the binary, execute the binary directly.

1. Via `go run`

`go run main.go -d <domain.com>`

2. Running the compiled binary

`binary` -d <domain.com>


# How it works

1. Fetching URLs: The program uses gau to retrieve all discoverable URLs related to the provided domain
2. HTTP Status Check: For each URL  retrieved, the program checks its HTTP status using curl and only keeps the URLs that return a status code of 200
3. Sensitive Pattern Search: The program scans the content of valid URLs for sensitive data patterns such as password, user, credentials, etc.
4. Output: If any sensitive patterns are found, the results are saved in pattern_matches.txt

# Pattern List
The following patterns are searched in the content of each valid URL:

- (?i)pass
- (?i)password
- (?i)contraseña
- (?i)contra
- (?i)cred
- (?i)credential
- (?i)credentials
- (?i)user
- (?i)username
- (?i)usuario
- (?i)cliente

This list can be expanded by modifying the patterns array in the utils/funcs.go file.


# Contributing

Feel free to open issues or submit pull requests if you want to improve or expand the functionality of this project.
