package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
	Err    error
}

var apiPatterns = []string{
	"https://%s.service-now.com/api/now/table/incident",
	"https://%s.my.salesforce.com/services/data/",
	"https://%s.atlassian.net/rest/api/2/project",
	"https://%s.azurewebsites.net/api/health",
	"https://%s.zendesk.com/api/v2/users.json",
	"https://%s.okta.com",
	"https://%s.slack.com",
	"https://%s.gitbook.io"
	"https://%s.box.com",
	"https://%s.hubspotpage.com",
	"https://%s.myshopify.com",
	"https://%s.statuspage.io",
	"https://vpn.%s.com",
	// (personalize sua lista aqui)
}

// ---------------- Executor ----------------

func check(client *http.Client, url string) Result {
	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Err: err}
	}
	defer resp.Body.Close()

	return Result{
		URL:    url,
		Status: resp.StatusCode,
	}
}

// ---------------- Classificação ----------------

func classify(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "🟢"
	case status >= 300 && status < 400:
		return "🟡"
	case status >= 400:
		return "🔴"
	default:
		return "⚪"
	}
}

// ---------------- Main ----------------

func main() {
	var company string
	fmt.Print("Digite a empresa: ")
	fmt.Scanln(&company)
	company = strings.ToLower(company)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup
	workers := 20

	// Workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				results <- check(client, url)
			}
		}()
	}

	// Producer
	go func() {
		for _, pattern := range apiPatterns {
			jobs <- fmt.Sprintf(pattern, company)
		}
		close(jobs)
	}()

	// Closer
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consumer
	for r := range results {
		if r.Err != nil {
			fmt.Printf("❌ %s (erro)\n", r.URL)
			continue
		}
		fmt.Printf("%s %s [%d]\n", classify(r.Status), r.URL, r.Status)
	}

	fmt.Println("Scan finalizado.")
}
