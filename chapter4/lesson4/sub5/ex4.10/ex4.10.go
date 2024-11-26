package main

import (
	"fmt"
	"log"
	"main/chapter4/lesson4/sub5/github"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
	})
	fmt.Printf("%dтем:\n", result.TotalCount)
	currDate := time.Now()

	fmt.Println("Created less then a month ago")
	for _, item := range result.Items {
		if currDate.Sub(item.CreatedAt) < 30*24*time.Hour {
			fmt.Printf("| %-5d | %-15s | %-55s | %-20v |\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Println("Created less then a year ago")
	for _, item := range result.Items {
		if currDate.Sub(item.CreatedAt) < 12*30*24*time.Hour &&
			currDate.Sub(item.CreatedAt) > 30*24*time.Hour {
			fmt.Printf("| %-5d | %-15s | %-55s | %-20v |\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Println("Created more then a year ago")
	for _, item := range result.Items {
		if currDate.Sub(item.CreatedAt) > 365*24*time.Hour {
			fmt.Printf("| %-5d | %-15s | %-55s | %-20v |\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

}
