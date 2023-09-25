package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to unmarshal JSON data
type CloudIPRanges struct {
	Prefixes []struct {
		IPv4Prefix string `json:"ipv4Prefix"`
		IPv6Prefix string `json:"ipv6Prefix"`
	} `json:"prefixes"`
}

func main() {
	// URL of the JSON file
	url := "https://www.gstatic.com/ipranges/cloud.json"

	// Fetch data from the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Decode JSON data
	var cloudIPRanges CloudIPRanges
	err = json.NewDecoder(resp.Body).Decode(&cloudIPRanges)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Initialize slices for IPv4 and IPv6 prefixes
	ipv4Prefixes := make([]string, 0)
	ipv6Prefixes := make([]string, 0)

	// Summarize the prefixes and add them to the respective slices
	for _, prefix := range cloudIPRanges.Prefixes {
		// Add IPv4 prefix to the IPv4 slice
		if prefix.IPv4Prefix != "" {
			ipv4Prefixes = append(ipv4Prefixes, prefix.IPv4Prefix)
		}
		// Add IPv6 prefix to the IPv6 slice
		if prefix.IPv6Prefix != "" {
			ipv6Prefixes = append(ipv6Prefixes, prefix.IPv6Prefix)
		}
	}

	// Print summarized IPv4 prefixes
	fmt.Println("Summarized IPv4 Prefixes:")
	for _, ipv4Prefix := range ipv4Prefixes {
		fmt.Println(ipv4Prefix)
	}

	// Print summarized IPv6 prefixes
	fmt.Println("\nSummarized IPv6 Prefixes:")
	for _, ipv6Prefix := range ipv6Prefixes {
		fmt.Println(ipv6Prefix)
	}

}
