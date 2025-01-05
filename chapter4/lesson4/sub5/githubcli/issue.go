package githubcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Label struct {
	Name string `json:"name"`
}

type Issue struct {
	Number      int       `json:"number,omitempty"`
	Title       string    `json:"title,omitempty"`
	Body        string    `json:"body,omitempty"`
	Assignees   []string  `json:"assignees,omitempty"`
	HTMLURL     string    `json:"html_url,omitempty"`
	State       string    `json:"state,omitempty"`
	StateReason string    `json:"state_reason,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Labels      []Label   `json:"labels,omitempty"`
}

type Params struct {
	Owner string
	Repo  string
	Issue Issue
}

const baseURL = "https://api.github.com/repos"

var githubToken = os.Getenv("GITHUB_TOKEN")

func (p *Params) GetIssues() ([]Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, p.Owner, p.Repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+githubToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (p *Params) GetIssue(number int) (*Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, p.Owner, p.Repo, number)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token"+githubToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (p *Params) CreateIssue() (*Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, p.Owner, p.Repo)
	issueData := map[string]interface{}{
		"title":     p.Issue.Title,
		"body":      p.Issue.Body,
		"assignees": p.Issue.Assignees,
	}

	body, err := json.Marshal(issueData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+githubToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createdIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&createdIssue); err != nil {
		return nil, err
	}

	return &createdIssue, nil
}

func (p *Params) UpdateIssue(number int) (*Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, p.Owner, p.Repo, number)
	updateData := map[string]interface{}{}
	if p.Issue.Title != "" {
		updateData["title"] = p.Issue.Title
	}
	if p.Issue.Body != "" {
		updateData["body"] = p.Issue.Body
	}
	if len(p.Issue.Assignees) > 0 {
		updateData["assignees"] = p.Issue.Assignees
	}

	body, err := json.Marshal(updateData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+githubToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&updatedIssue); err != nil {
		return nil, err
	}

	return &updatedIssue, nil
}

func (p *Params) CloseIssue(number int) (*Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, p.Owner, p.Repo, number)
	updateIssue := map[string]interface{}{}
	if p.Issue.State != "" {
		updateIssue["state"] = p.Issue.State
	}

	body, err := json.Marshal(updateIssue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+githubToken)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var closedIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&closedIssue); err != nil {
		return nil, err
	}

	return &closedIssue, nil
}
