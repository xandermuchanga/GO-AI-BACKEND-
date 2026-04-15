package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"ai-customer-support/models"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewAIService(apiKey string) (*AIService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	
	// Set system instruction for the model
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("You are a professional customer support agent. Classify the user request and provide a helpful response. Respond in JSON with: response, category."),
		},
	}

	return &AIService{
		client: client,
		model:  model,
	}, nil
}

func (s *AIService) ProcessMessage(ctx context.Background(), message string) (*models.AIResponse, error) {
	prompt := fmt.Sprintf("User message: %s", message)
	
	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("no response from AI")
	}

	// Extract text from the response
	var responseText string
	for _, part := range resp.Candidates[0].Content.Parts {
		if text, ok := part.(genai.Text); ok {
			responseText += string(text)
		}
	}

	log.Printf("AI Raw Response: %s", responseText)

	// Clean up the response if it contains markdown code blocks
	cleanJSON := strings.TrimPrefix(strings.TrimSuffix(strings.TrimSpace(responseText), "```"), "```json")
	cleanJSON = strings.TrimSpace(cleanJSON)

	var aiResp models.AIResponse
	err = json.Unmarshal([]byte(cleanJSON), &aiResp)
	if err != nil {
		// Fallback if JSON parsing fails
		return &models.AIResponse{
			Response: responseText,
			Category: "general inquiry",
		}, nil
	}

	return &aiResp, nil
}

func (s *AIService) Close() {
	s.client.Close()
}
