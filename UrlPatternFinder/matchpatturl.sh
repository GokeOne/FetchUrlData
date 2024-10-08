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

if ! command -v curl &> /dev/null; then
	echo "[${RED}ERROR${NC}] curl is not installed. Plase install rquirements to continue.."
	exit 1
fi


search_patterns() {
	local valid_url="$1"
	local pattern_file="pattern_matches.txt"
	> "$pattern_file"

	patterns=("pass" "password" "contraseña" "contra" "cred" "credential" "credentials" "user" "username" "usuario" "cliente")

	while read -r url; do
		#download content and delete null bytes
		content=$(curl -s "$url" | tr -d '\0')

		for pattern in "${patterns[@]}"; do
			if echo "$content" | grep -Eqi "$pattern"; then
				echo "[${CYAN}INFO${NC}] Sensitive pattern found in: $url"
				echo "${CYAN}URL: ${NC}$url | ${CYAN}Pattern: ${NC}$pattern" >> "$pattern_file"
			fi
		done
	done < "$valid_url"
}

if [ "$#" -eq 1 ]; then
	valid_url="$1"
else
	read -p "Enter the path to the file containing URLs: " valid_url
fi

if [ ! -f "$valid_url" ] || [ ! -s "$valid_url" ]; then
	echo "[${RED}ERROR${NC}] The file $valid_url does not exist or is empty."
	exit 1
fi

search_patterns "$valid_url"
