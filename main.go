package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GitHubRepo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	CreatedAt   time.Time `json:"created_at"`
	Private     bool      `json:"private"`
}

func getGitHubRepos(url string) ([]GitHubRepo, error) {

	repos := []GitHubRepo{}

	if len(url) == 0 {
		return nil, fmt.Errorf("URL is empty")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repos: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return repos, nil
}

func main() {

	username := "aparandianworld"
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	repos, err := getGitHubRepos(url)

	if err != nil {
		fmt.Println("Error getting GitHub repos:", err)
		return
	}

	if len(repos) == 0 {
		fmt.Println("No GitHub repos found")
		return
	}

	fmt.Printf("Found %d public GitHub repos for user %s", len(repos), username)

	for i, repo := range repos {
		fmt.Printf("Repo %d: %s\n", i+1, repo.Name)
		fmt.Printf("Description: %s\n", repo.Description)
		fmt.Printf("Language: %s\n", repo.Language)
		fmt.Printf("Created At: %s\n", repo.CreatedAt)
		fmt.Printf("Private: %t\n", repo.Private)
		fmt.Println()
	}

}
