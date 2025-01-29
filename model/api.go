package model

import (
	"fmt"
	"time"
)

type ApiResponse struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_datetime"`
	GithubUrl       string `json:"github_url"`
}

func NewApiResponse() (*ApiResponse, error) {

	location, err := timezone()
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().In(location)
	return &ApiResponse{
		Email:           "jtirenipraise@gmail.com",
		CurrentDateTime: currentTime.Format(time.RFC3339),
		GithubUrl:       "https://github.com/Jidetireni/hng_stage_0.git",
	}, nil
}

func timezone() (*time.Location, error) {
	location, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		return nil, fmt.Errorf("could not load location: %v", err)
	}

	return location, nil

}
