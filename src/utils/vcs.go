package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
	Message string `json:"message"`
}

func FetchLatestTag() (string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/CrysoK/cli/releases/latest", nil)
	if err != nil {
		return "", err
	}

	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var release GithubRelease
	if err = json.Unmarshal(body, &release); err != nil {
		return "", err
	}

	if release.TagName == "" {
		return "", errors.New("GitHub response: " + release.Message)
	}

	return release.TagName[1:], nil
}
