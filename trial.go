package main

import (
	// "bytes"
	"encoding/json"
	// "mime/multipart"
	"strings"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	payload, err := get_file("sample.txt")
	if err != nil {
		fmt.Println("Error getting file:", err)
		return
	}

	fmt.Println(payload) //debug

	// for _, fileName := range files {
        // fileContent, err := os.ReadFile(fileName)
        // if err != nil {
            // fmt.Println("Error reading file:", err)
            // continue
        // }

        // payload := strings.NewReader(fmt.Sprintf("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\nContent-Type: text/plain\r\n\r\n%s\r\n-----011000010111000001101001--", fileName, fileContent))

	url := "https://www.virustotal.com/api/v3/files"

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", "f6335ba146578f6a07ddf8e11af1966423366278fda02b750262b4833bcbfedc")
	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }

    defer res.Body.Close()
    body, _ := io.ReadAll(res.Body)

    fmt.Println(string(body))
	
	var file_id string = get_id(body)

	fmt.Println(file_id)

}

func get_file(FileName string) (*strings.Reader, error) {
	fileContent, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
		
	}

	payload := strings.NewReader(fmt.Sprintf("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\nContent-Type: text/plain\r\n\r\n%s\r\n-----011000010111000001101001--", FileName, fileContent))

	fmt.Println("File Opened Successfully")
	return payload, nil
}

func get_id (body []byte) string {
	var id string = ""
	fmt.Println(string(body))
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return ""
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if id, ok := data["id"].(string); ok {
			fmt.Println(id)
			
		} else {
			fmt.Println("ID not found in response")
		}
	} else {
		fmt.Println("Data not found in response")
	}
	return id
}