package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	file = get_file()
	url := "https://www.virustotal.com/api/v3/files"

	req, _ := http.NewRequest("POST", url, )

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "multipart/form-data")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func get_file() *os.File {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	return file
}