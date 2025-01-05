package githubcli

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenEditor(initialContent string) (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "notepad"
	}

	tempFile, err := os.CreateTemp("", "githubcli_issue_*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(initialContent); err != nil {
		return "", fmt.Errorf("failed to write initial content to temp file: %w", err)
	}
	tempFile.Close()

	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run editor: %w", err)
	}

	updatedContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to read updated content from temp file: %w", err)
	}

	return string(updatedContent), nil
}

func GetContentFromEditor(initialContent string) (string, error) {
	updatedContent, err := OpenEditor(initialContent)
	if err != nil {
		return "", fmt.Errorf("failed to get content from editor: %w", err)
	}
	return updatedContent, nil
}
