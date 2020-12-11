package main

import (
	"fmt"
	"os"
	"strconv"
)

func downloadPublication(id string) error {
	archive, err := fetchArchiveResponse(id)
	if err != nil {
		return err
	}

	issues, err := archive.fetchIssues()
	if err != nil {
		return err
	}

	for i, issue := range issues {
		for _, page := range issue.Issue.Pages {
			var issueAttr = archive.Archive.Issues[i].Attributes

			var sequentialPage, err = strconv.Atoi(page.Attributes.SequentialPage)
			if err != nil {
				return err
			}

			// target filePath to save the HTTP response to
			var filePath = fmt.Sprintf("issue_%s_(%s)_page_%03d.pdf", issueAttr.IssueId, issueAttr.IssueName, sequentialPage)

			fmt.Println(filePath)

			if err = fetch(page.Attributes.PDF, filePath); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// parse publication ID directly from os.Args
	// while the flags package can accomplish this, the argument is required, and since there's only one argument
	// 	it makes the most sense to simply parse (and require) len(os.Args) == 1
	if len(os.Args) != 2 {
		fmt.Printf("Usage: digimag-downloader <publication ID>")
		return
	}

	if err := downloadPublication(os.Args[1]); err != nil {
		panic(err)
	}
}
