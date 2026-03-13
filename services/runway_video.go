package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type runwayRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	PromptImage string `json:"promptImage"`
	Ratio       string `json:"ratio"`
	Duration    int    `json:"duration"`
}

type runwayResponse struct {
	ID string `json:"id"`
}

func GenerateRunwayVideo(imageURL string, prompt string, duration int) (string, error) {

	apiKey := os.Getenv("RUNWAY_API_KEY")

	if apiKey == "" {
		return "", errors.New("runway api key not configured")
	}

	// Veo models only support 4,6,8 seconds
	if duration != 4 && duration != 6 && duration != 8 {
		duration = 8
	}

	reqBody := runwayRequest{
		Model:       "veo3.1_fast",
		Prompt:      prompt,
		PromptImage: imageURL,
		Ratio:       "720:1280",
		Duration:    duration,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	fmt.Println("Runway request payload:", string(jsonBody))

	req, err := http.NewRequest(
		"POST",
		"https://api.dev.runwayml.com/v1/image_to_video",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Runway-Version", "2024-11-06")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Runway status:", resp.StatusCode)
	fmt.Println("Runway response:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", errors.New("runway request failed")
	}

	var result runwayResponse

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}

	if result.ID == "" {
		return "", errors.New("runway returned empty job id")
	}

	return result.ID, nil
}
