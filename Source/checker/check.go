package checker

import (
	"fmt"
	"os"
)

func ValidateDomain(domain string) {
	if domain == "" {
		Usage()
		os.Exit(1)
	}
}

func Usage() {
	fmt.Printf("Usage: ./fetchurldata -d <domain>")
}
