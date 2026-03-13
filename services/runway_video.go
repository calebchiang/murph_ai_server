package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type runwayRequest struct {
	Prompt string `json:"prompt"`
	Image  string `json:"image"`
	Model  string `json:"model"`
}

type runwayResponse struct {
	ID string `json:"id"`
}

func GenerateRunwayVideo(imageURL string, prompt string) (string, error) {

	apiKey := os.Getenv("RUNWAY_API_KEY")

	if apiKey == "" {
		return "", errors.New("runway api key not configured")
	}

	reqBody := runwayRequest{
		Prompt: prompt,
		Image:  imageURL,
		Model:  "gen4_turbo",
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(
		"POST",
		"https://api.runwayml.com/v1/image_to_video",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result runwayResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	return result.ID, nil
}
