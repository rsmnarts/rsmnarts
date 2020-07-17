package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func UpdateGist(idGist string, payload interface{}) {
	baseUrlGit := "https://api.github.com"

	reqBody, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/gists/%s", baseUrlGit, idGist), bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("ACCESS_TOKEN")))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
}
