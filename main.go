package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// cookie jar to store cookies between requests
	jar, _ := cookiejar.New(nil)

	// HTTP client with the cookie jar
	client := &http.Client{Jar: jar}

	loginURL := "https://incident-management-api.amalitech-dev.net/auth/login"
	loginData := map[string]string{
		"email":    "daniel.mensah@amalitech.com",
		"password": "mnTrPezkA9HkORre",
	}

	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	data := bytes.NewBuffer(loginJSON)
	fmt.Println(data)
	fmt.Println("---------------------------------------------------------")

	// login request with a JSON body
	loginResp, err := client.Post(loginURL, "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		fmt.Println("Login request failed:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(loginResp.Body)
	fmt.Println("---------------------------------------------------------")

	// Print the response status code and headers
	fmt.Printf("Login Status: %s\n", loginResp.Status)
	fmt.Println("Login Response Headers:")
	for key, values := range loginResp.Header {
		fmt.Printf("%s: %s\n", key, values)
	}

	//************************************************************************
	//************************************************************************

	if loginResp.Status == "201 Created" {
		fmt.Println("---------------------------------------------------------")

		incidentEndpoint := "https://incident-management-api.amalitech-dev.net/incident"

		fmt.Println("---------------------------------------------------------")

		//request to the incident endpoint
		req, err := http.NewRequest("POST", incidentEndpoint, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Make the request and capture the response
		loginResp, err = client.Do(req)
		if err != nil {
			fmt.Println("Request to incident endpoint failed:", err)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(loginResp.Body)

		// Print the response status code and headers
		fmt.Printf("Incidient Endpoint Status: %s\n", loginResp.Status)
		fmt.Println("Incident Endpoint Response Headers:")
		for key, values := range loginResp.Header {
			fmt.Printf("%s: %s\n", key, values)
		}
	}

}
