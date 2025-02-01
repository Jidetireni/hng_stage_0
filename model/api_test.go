package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewApiResponse(t *testing.T) {
	response, err := NewApiResponse()
	require.NoError(t, err)

	require.Equal(t, "jtirenipraise@gmail.com", response.Email)
	require.Equal(t, time.Now().UTC().Format(time.RFC3339), response.CurrentDateTime)
	require.Equal(t, "https://github.com/Jidetireni/hng_stage_0.git", response.GithubUrl)
}
