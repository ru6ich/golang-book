package githubcli

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func FormatIssueOutput(issue Issue) string {
	var labelNames []string
	for _, label := range issue.Labels {
		labelNames = append(labelNames, label.Name)
	}
	return fmt.Sprintf(
		"Issue #%d: %s\nState: %s\nAssigned to: %s\nLabels: %s\nURL: %s\n",
		issue.Number,
		issue.Title,
		issue.State,
		strings.Join(issue.Assignees, ", "),
		strings.Join(labelNames, ", "),
		issue.HTMLURL,
	)
}

func HandleError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	}
	os.Exit(1)
}

func CheckResponse(resp *http.Response, expectedStatus int) error {
	if resp.StatusCode != expectedStatus {
		return fmt.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, expectedStatus)
	}
	return nil
}
