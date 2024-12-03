package main

import (
	import "github.com/0xQRx/crtsh-tool/pkg/crtsh"
	"flag"
	"fmt"
	"log"
	// "os"
)

func main() {
	domain := flag.String("domain", "", "Domain to query on crt.sh")
	outputFile := flag.String("outputfile", "", "File to save output (optional)")

	flag.Parse()

	if *domain == "" {
		log.Fatal("Error: --domain is required")
	}

	//fmt.Printf("[DEBUG] Querying crt.sh for domain: %s\n", *domain)

	results, err := crtsh.FetchDomains(*domain)
	if err != nil {
		log.Fatalf("[DEBUG] Error fetching domains: %v\n", err)
	}

	//fmt.Printf("[DEBUG] Found %d domains\n", len(results))

	if *outputFile != "" {
		fmt.Printf("[DEBUG] Writing results to file: %s\n", *outputFile)
		err := crtsh.WriteToFile(*outputFile, results)
		if err != nil {
			log.Fatalf("[DEBUG] Error writing to file: %v\n", err)
		}
		fmt.Printf("[DEBUG] Results successfully written to file: %s\n", *outputFile)
	} else {
		//fmt.Println("[DEBUG] Domains:")
		for _, domain := range results {
			fmt.Println(domain)
		}
	}
}
