# Description

This Bash script scans a specified domain for URLs using the gau tool (Get All URLs), checks their HTTP status using curl, and saves valid URLs (i.e., those returning HTTP status code 200) to a file. It's useful for gathering potentially interesting URLs from a target domain and validating their accessibility.

# How to use

`./urlchecker -d domain.com`


There are two ways to install gau

**First option:**
- `apt install golang-go`
- `go install github.com/lc/gau/v2/cmd/gau@latest`

**Second option**
-   First u need to go to: [releases](https://github.com/lc/gau/releases/)
-   `tar xvf gau_2.2.3_linux_amd64.tar.gz`
-   `mv gau /usr/bin/gau`

**Note**: This script is designed for educational and testing purposes. Always ensure you have permission to scan and gather URLs from the target domain. Unauthorized use may violate legal and ethical guidelines.
