package utils

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func GetUrls(domain string) []string {
	var urls []string

	//Exec gau to get URLs
	cmd := exec.Command("gau", domain)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("[ERROR] Failed to run gau for domain %s: %v\n", domain, err)
	}

	// Read gau output and save URLs in list
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("[ERROR] Error reading gau output: %v\n", err)
	}

	return urls
}

func CheckValidUrls(urls []string) []string {
	var validUrls []string
	for _, url := range urls {
		status := getStatus(url)
		if status == 200 {
			fmt.Printf("[INFO] Valid URL: %s\n", url)
			validUrls = append(validUrls, url)
		} else {
			fmt.Printf("[ERROR] Invalid URL (status %d): %s\n", status, url)
		}
	}
	return validUrls
}

func getStatus(url string) int {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch URL %s: %v\n", url, err)
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func SearchSensitivePatterns(validUrls []string) {
	patterns := []string{
		"(?i)pass", "(?i)password", "(?i)contraseña", "(?i)contra", "(?i)cred", "(?i)credential",
		"(?i)credentials", "(?i)user", "(?i)username", "(?i)usuario", "(?i)cliente", "(?i)hash",
	}

	patternFile := "pattern_matches.txt"
	file, err := os.Create(patternFile)
	if err != nil {
		log.Fatalf("[ERROR] Could not create patther matches file : %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, url := range validUrls {
		fmt.Printf("[INFO] Scanning URL: %s\n", url)
		content, err := fetchContent(url)
		if err != nil {
			fmt.Printf("[ERROR] Failed to fetch content form %s: %v\n", url, err)
			continue
		}

		for _, pattern := range patterns {
			matched, err := regexp.MatchString(pattern, content)
			if err != nil {
				log.Printf("[ERROR] Regex error: %v\n", err)
				continue
			}
			if matched {
				fmt.Printf("[INFO] Sensitive pattern found in: %s | Pattern: %s\n", url, pattern)
				writer.WriteString(fmt.Sprintf("URL: %s | Pattern: %s\n", url, pattern))
			}
		}
	}

	writer.Flush()

}

func fetchContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}

	//Read body response
	var contentBuilder strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return contentBuilder.String(), nil
}
