package json_type

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	IssueURL = "https://api.github.com/search/issues"
)

type IssueSearchResult struct {
	TotalCount int `json1:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json1:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json1:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json1:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	fmt.Printf("IssueURL + \"?q=\" + q: %s\n", IssueURL+"?q="+q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search response %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
	}
	resp.Body.Close()
	return &result, nil
}
