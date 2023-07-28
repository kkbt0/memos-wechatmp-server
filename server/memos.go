package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RequestMemosData struct {
	Content        string   `json:"content"`
	Visibility     string   `json:"visibility"`
	ResourceIDList []int    `json:"resourceIdList"`
	RelationList   []string `json:"relationList"`
}

// PRIVATE PUBLIC
func CreateMemo(url string, content string, visibility string, resourceIdList []int) error {

	payload := RequestMemosData{
		Content:        content,
		Visibility:     visibility,
		ResourceIDList: resourceIdList,
		RelationList:   []string{},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Body)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

type RequestResourceData struct {
	Filename        string `json:"filename"`
	ExternalLink    string `json:"externalLink"`
	Type            string `json:"type"`
	DownloadToLocal bool   `json:"downloadToLocal"`
}

type ResponseResourceData struct {
	Data struct {
		ID            int    `json:"id"`
		CreatorID     int    `json:"creatorId"`
		CreatedTs     int    `json:"createdTs"`
		UpdatedTs     int    `json:"updatedTs"`
		Filename      string `json:"filename"`
		ExternalLink  string `json:"externalLink"`
		Type          string `json:"type"`
		Size          int    `json:"size"`
		PublicID      string `json:"publicId"`
		LinkedMemoAmt int    `json:"linkedMemoAmount"`
	} `json:"data"`
}

func CreateResourceByLink(url string, externalLink string) (int, error) {
	url = strings.ReplaceAll(url, "memo?", "resource?")
	data := RequestResourceData{
		Filename:        "",
		ExternalLink:    externalLink,
		Type:            "",
		DownloadToLocal: true,
	}
	// Convert the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return 0, err
	}

	// Create a request with the specified method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Parse the response JSON into the ResponseData struct
	var responseData ResponseResourceData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return 0, err
	}

	// Extract the ID from the response data
	fmt.Println("Resource:", responseData)
	id := responseData.Data.ID
	return id, nil
}
