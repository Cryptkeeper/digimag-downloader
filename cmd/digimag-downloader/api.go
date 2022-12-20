package main

import (
	"encoding/json"
)

type issueResponse struct {
	Issue struct {
		Pages []struct {
			Attributes struct {
				PDF            string `json:"pdf"`
				SequentialPage string `json:"sequential_page"`
				Name           string `json:"name"`
			} `json:"@attributes"`
		} `json:"page"`
	} `json:"issue"`
}

type archiveResponse struct {
	Archive struct {
		Issues []struct {
			Attributes struct {
				Date      string `json:"date"`
				IssueName string `json:"issue_name"`
				IssueId   string `json:"issueid"`
			} `json:"@attributes"`
		} `json:"issue"`
	} `json:"archive"`
}

func (a archiveResponse) fetchIssues() ([]issueResponse, error) {
	var issues = make([]issueResponse, len(a.Archive.Issues))

	for i, issue := range a.Archive.Issues {
		url := "https://mydigimag.rrd.com/publication/globals.php?id_issue=" + issue.Attributes.IssueId + "&out=json"
		b, err := requestBody(url)

		if err = json.Unmarshal(b, &issues[i]); err != nil {
			return nil, err
		}
	}

	return issues, nil
}

func fetchArchiveResponse(id string) (archiveResponse, error) {
	var archive archiveResponse

	url := "https://mydigimag.rrd.com/publication/archive.php?id_publication=" + id + "&out=json"
	b, err := requestBody(url)

	if err != nil {
		return archive, err
	} else {
		return archive, json.Unmarshal(b, &archive)
	}
}
