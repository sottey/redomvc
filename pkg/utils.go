package pkg

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

func CheckSingleDomain(domain string) {
	delay := time.Duration(viper.GetInt("delay")) * time.Millisecond
	time.Sleep(delay)
	checkDomain(domain)
}

func CheckFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup
	sem := make(chan struct{}, viper.GetInt("workers"))
	delay := time.Duration(viper.GetInt("delay")) * time.Millisecond
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			sem <- struct{}{}
			time.Sleep(delay)
			checkDomain(d)
			<-sem
		}(domain)
	}
	wg.Wait()
}

func checkDomain(domain string) {
	verbose := viper.GetBool("verbose")
	if verbose {
		fmt.Printf("Checking: '%s'\n", domain)
	}

	payload := CheckAvailabilityRequest{DomainNames: []string{domain}}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", viper.GetString("api"), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	user := viper.GetString("username")
	token := viper.GetString("token")
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, token)))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("User-Agent", "domaincheck/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error for %s: %v\n", domain, err)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if verbose {
		fmt.Printf("Raw response: %s\n", string(data))
	}

	if resp.StatusCode != 200 {
		fmt.Printf("HTTP %d for %s\nBody: %s\n", resp.StatusCode, domain, string(data))
		return
	}

	var result CheckAvailabilityResponse
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Printf("JSON error for %s: %v\nRaw: %s\n", domain, err, string(data))
		return
	}

	if len(result.Results) > 0 && result.Results[0].Purchasable {
		price := result.Results[0].PurchasePrice
		renewalPrice := result.Results[0].RenewalPrice
		fmt.Printf("%s is available for $%.2f (Renewal: $%.2f)\n", result.Results[0].DomainName, price, renewalPrice)
	} else {
		fmt.Printf("%s is NOT available.\n", domain)
	}
}
