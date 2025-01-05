package githubcli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Execute() {
	// Флаги для owner и repo
	owner := flag.String("owner", "", "Owner of the repository (required)")
	repo := flag.String("repo", "", "Repository name (required)")

	// Остальные флаги
	list := flag.Bool("list", false, "List all issues")
	get := flag.Int("get", 0, "Get issue by number")
	create := flag.Bool("create", false, "Create a new issue")
	update := flag.Int("update", 0, "Update issue by number")
	close_ := flag.Int("close", 0, "Close issue by number")
	reopen := flag.Int("reopen", 0, "Reopen issue by number")

	// Дополнительные параметры для создания/обновления
	title := flag.String("title", "", "Title of the issue")
	body := flag.String("body", "", "Body of the issue")
	assignees := flag.String("assignees", "", "Comma-separated list of assignees")

	flag.Parse()

	//Проверяем, указаны ли owner и repo
	if *owner == "" || *repo == "" {
		fmt.Fprintln(os.Stderr, "Error: both --owner and --repo flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	// Проверяем, что хотя бы один из основных флагов установлен
	if !*list && *get == 0 && !*create && *update == 0 && *close_ == 0 && *reopen == 0 && *title == "" && *body == "" && *assignees == "" {
		fmt.Fprintln(os.Stderr, "Error: no valid action flag provided. Use -h for help.")
		os.Exit(1)
	}

	// Проверяем токен
	token := GetAuthToken()
	err := ValidateToken(token)
	if err != nil {
		fmt.Println("DEBUG: Token validation error:", err)
	}
	HandleError(err, "Invalid GitHub token")

	// Создаем объект Params
	params := Params{
		Owner: *owner,
		Repo:  *repo,
	}

	// Определяем действия
	switch {
	case *list:
		issues, err := params.GetIssues()
		if err != nil {
			fmt.Println("DEBUG: GetIssues error:", err)
		}
		HandleError(err, "Failed to list issues")
		for _, issue := range issues {
			fmt.Println(FormatIssueOutput(issue))
		}

	case *get > 0:
		params.Issue.Number = *get
		issue, err := params.GetIssue(*get)
		HandleError(err, "Failed to get issue")
		fmt.Println(FormatIssueOutput(*issue))

	case *create:
		params.Issue = Issue{
			Title:     *title,
			Body:      *body,
			Assignees: parseAssignees(*assignees),
		}
		issue, err := params.CreateIssue()
		HandleError(err, "Failed to create issue")
		fmt.Println("Issue created:")
		fmt.Println(FormatIssueOutput(*issue))

	case *update > 0:
		params.Issue = Issue{
			Number: *update,
			Title:  *title,
			Body:   *body,
		}
		issue, err := params.UpdateIssue(*update)
		HandleError(err, "Failed to update issue")
		fmt.Println("Issue updated:")
		fmt.Println(FormatIssueOutput(*issue))

	case *close_ > 0:
		issue, err := params.CloseIssue(*close_)
		HandleError(err, "Failed to close issue")
		fmt.Println("Issue closed:")
		fmt.Println(FormatIssueOutput(*issue))

	case *reopen > 0:
		issue, err := params.UpdateIssue(*reopen)
		HandleError(err, "Failed to reopen issue")
		fmt.Println("Issue reopened:")
		fmt.Println(FormatIssueOutput(*issue))

	default:
		fmt.Println("Please provide a valid flag. Use -h for help.")
		os.Exit(1)
	}
}

func parseAssignees(assignees string) []string {
	if assignees == "" {
		return nil
	}
	return strings.Split(assignees, ",")
}
