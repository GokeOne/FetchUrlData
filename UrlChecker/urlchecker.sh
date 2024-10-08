#!/bin/bash

# Color Definitions
C=$(printf '\033')
RED="${C}[1;31m"
GREEN="${C}[1;32m"
YELLOW="${C}[1;33m"
BLUE="${C}[1;34m"
MAGENTA="${C}[1;35m"
CYAN="${C}[1;36m"
LIGHT_GRAY="${C}[1;37m"
DARK_GRAY="${C}[1;90m"
NC="${C}[0m" # No Color
UNDERLINED="${C}[4m"
ITALIC="${C}[3m"
PARPADEO="${C}[1;5m"

# Usage function
usage() {
    echo -e "${YELLOW}Usage: $0 -d domain.com${NC}"
    echo -e "${YELLOW}Options:${NC}"
    echo -e "  ${GREEN}-d${NC}  Specify the domain to scan"
    echo -e "  ${GREEN}-h${NC}  Show this help message"
    exit 1
}

# Check if gau is installed
if ! command -v gau &> /dev/null; then
	echo "[${RED}ERROR${NC}] gau is not installed. Please install requirements to continue."
	exit 1
fi

# Check if curl is installed
if ! command -v curl &> /dev/null; then
	echo "[${RED}ERROR${NC}] curl is not installed. Please install requirements to continue."
	exit 1
fi

# Parse command-line arguments
while getopts "d:h" opt; do
  case $opt in
    d)
      domain=$OPTARG
      ;;
    h)
      usage
      ;;
    *)
      usage
      ;;
  esac
done

# Check if domain is provided
if [ -z "$domain" ]; then
    echo -e "[${RED}ERROR${NC}] Domain is required."
    usage
fi

# Temp file for storing URLs
temp_file=$(mktemp)

echo "[${GREEN}INFO${NC}] Getting URLs with gau..."
gau "$domain" | tee "$temp_file" >& /dev/null

valid_url="valid_url.txt"
> "$valid_url"

# Check valid URLs
while read -r url; do
	response_code=$(curl -o /dev/null -s -w "%{http_code}" "$url")
	#ffuf -w "$temp_file" -u FUZZ -mc 200 -o ffuf_results.txt
	if [ "$response_code" -eq 200 ]; then
		echo "$url" | tee -a "$valid_url" &> /dev/null
	fi
done < "$temp_file"

# Clean up
rm "$temp_file"
echo "[${GREEN}INFO${NC}] Done. The valid URLs are in $valid_url"
