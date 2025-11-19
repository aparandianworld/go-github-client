package main

import (
	"fmt"
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
	// TODO
	return nil, nil
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
