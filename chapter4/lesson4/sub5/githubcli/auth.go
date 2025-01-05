package githubcli

import (
	"fmt"
	"net/http"
	"os"
)

func GetAuthToken() string {
	fmt.Println("DEBUG: Fetching GITHUB_TOKEN from environment")
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Fprintln(os.Stderr, "Error: GITHUB_TOKEN is not set. Please set it as an environment variable.")
		os.Exit(1) // Логируем перед выходом
	}
	fmt.Println("DEBUG: Fetched GITHUB_TOKEN successfully")
	return githubToken
}

func ValidateToken(token string) error {
	fmt.Println("DEBUG: Validating token")
	url := "https://api.github.com/user"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("DEBUG: Token validation HTTP status:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid token: received status %d", resp.StatusCode)
	}

	fmt.Println("DEBUG: Token is valid")
	return nil
}
