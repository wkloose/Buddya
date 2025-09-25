package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/wkloose/Buddya/models"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func CallGeminiAPI(prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY belum diset di environment")
	}

	url := "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions"

	reqBody := ChatRequest{
		Model: "gemini-2.0-flash",
		Messages: []ChatMessage{
			{Role: "system", Content: "You are an AI assistant that provides cultural event recommendations."},
			{Role: "user", Content: prompt},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBytes, &chatResp); err != nil {
		return "", fmt.Errorf("gagal unmarshal response: %v; body: %s", err, string(respBytes))
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("tidak ada hasil dari Gemini API: %s", string(respBytes))
}
func ParseAIResponse(aiResponse string) ([]models.AIEvent, error) {
	aiResponse = strings.TrimSpace(aiResponse)
	aiResponse = strings.TrimPrefix(aiResponse, "```json")
	aiResponse = strings.TrimPrefix(aiResponse, "```")
	aiResponse = strings.TrimSuffix(aiResponse, "```")

	var events []models.AIEvent
	if err := json.Unmarshal([]byte(aiResponse), &events); err != nil {
		return nil, fmt.Errorf("gagal parsing AI response: %v", err)
	}

	return events, nil
}
