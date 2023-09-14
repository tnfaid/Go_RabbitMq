package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// Disable hostname verification (not recommended for production)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	//create http request
	jsonPayloadNewImai := []byte(`{"transaction_id":"cfa3d778-556c-4d88-b8b3-c79705bfd4de","incentive_rules_param":"{arp_flag=[AP], channel_id=[i4]}","bonus":"{business_id=[00054986]}","main_trigger":"business_id|00048265","incentive_rules_id":"INCTR_IMEI003","msisdn":"6281231990215","event_topic_trigger":"cjx-dom-in"}`)
	url := "https://localhost:11000/pes/v1/incentive-tracking/save"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayloadNewImai))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "MTI3MTgxZmI2MjI3MWViZmMyYzIzNWU3OGMyNjk0MTcyNjAzNTUwMTZkNTFkZGFjNzY1Nzg3NGFkNzQ1OGNiNg==")

	//send the http request
	// Create an HTTP client with the custom TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	//handle response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code %d\n", resp.StatusCode)
		// Handle the error or return
	}

	//Read and parse the response body
	//Example read as string
	var responseBody string
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(resp.Body)
	responseBody = buffer.String()

	fmt.Println("Response:", responseBody)
}
