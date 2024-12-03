package crtsh

import (
	"fmt"
	"io"
	"net/http"
	//"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const crtshURL = "https://crt.sh/?q="
const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.199 Safari/537.36"

// FetchDomains queries crt.sh and parses the results to return a slice of domains.
func FetchDomains(domain string) ([]string, error) {
	url := fmt.Sprintf("%s%%.%s", crtshURL, domain)

	//fmt.Printf("[DEBUG] Fetching URL: %s\n", url)

	// Create a custom HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("[DEBUG] Failed to create HTTP request: %v", err)
	}

	// Add custom headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Connection", "keep-alive")

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[DEBUG] Failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[DEBUG] Non-200 response: %d", resp.StatusCode)
	}

	//fmt.Printf("[DEBUG] HTTP request successful, status code: %d\n", resp.StatusCode)

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[DEBUG] Failed to read response body: %v", err)
	}
	//fmt.Printf("[DEBUG] Response body length: %d bytes\n", len(bodyBytes))

	// Save the response for manual inspection
	// err = os.WriteFile("response.html", bodyBytes, 0644)
	// if err != nil {
	// 	fmt.Printf("[DEBUG] Failed to save response body: %v\n", err)
	// } else {
	// 	fmt.Println("[DEBUG] Response body saved to response.html")
	// }

	// Parse the response HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(bodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("[DEBUG] Failed to parse HTML: %v", err)
	}

	// Try to locate the results table dynamically
	var resultsTable *goquery.Selection
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if table.Find("tr").Length() > 1 {
			//fmt.Printf("[DEBUG] Table %d has %d rows\n", i, table.Find("tr").Length())
			resultsTable = table
		}
	})

	if resultsTable == nil {
		return nil, fmt.Errorf("[DEBUG] Results table not found in HTML")
	}

	// Parse the domains from the results table
	var results []string
	resultsTable.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 { // Skip header row
			return
		}

		// Look for the sixth column (adjust if needed)
		td := row.Find("td").Eq(5)
		if td.Length() > 0 {
			// Extract domains while handling <br> tags
			var domainStrings []string
			td.Contents().Each(func(i int, selection *goquery.Selection) {
				if goquery.NodeName(selection) == "br" {
					domainStrings = append(domainStrings, "")
				} else {
					text := strings.TrimSpace(selection.Text())
					if len(domainStrings) == 0 || domainStrings[len(domainStrings)-1] != "" {
						domainStrings = append(domainStrings, text)
					} else {
						domainStrings[len(domainStrings)-1] = text
					}
				}
			})
			for _, domain := range domainStrings {
				if domain != "" {
					results = append(results, domain)
				}
			}
		}
	})

	// Deduplicate the results
	uniqueResults := removeDuplicates(results)

	//fmt.Printf("[DEBUG] Total unique domains found: %d\n", len(uniqueResults))
	return uniqueResults, nil
}

// removeDuplicates removes duplicate entries from a slice of strings.
func removeDuplicates(items []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
