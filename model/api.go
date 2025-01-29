package model

import "time"

type ApiResponse struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}

func NewApiResponse() (*ApiResponse, error) {
	return &ApiResponse{
		Email:           "jtirenipraise@gmail.com",
		CurrentDateTime: time.Now().UTC().Format(time.RFC3339),
		GithubUrl:       "https://github.com/Jidetireni/hng_stage_0.git",
	}, nil
}
