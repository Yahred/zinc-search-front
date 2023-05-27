package models

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"zinc-index-back.com/config"
)

type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}

type SearchEmail struct {
	SearchType string   `json:"search_type"`
	Query      Query    `json:"query"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

func SearchEmails(query string, page int, pageSize int) ([]Email, error) {
	queryObject := SearchEmail{
		SearchType: "match",
		Query: Query{
			Term:  query,
			Field: "body",
		},
		From:       (page - 1) * pageSize,
		MaxResults: pageSize,
		Source:     []string{"from", "to", "subject", "body"},
	}

	jsonQuery, err := json.Marshal(queryObject)

	if err != nil {
		return nil, err
	}

	url := config.ZincSearchURL + "/emails/_search"

	encodedZincSearchCredentials := config.GetZincSearchCredentials()

	headers := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Basic " + encodedZincSearchCredentials},
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonQuery))

	if err != nil {
		return nil, err
	}

	req.Header = headers

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var search ZincSearch

	err = json.NewDecoder(resp.Body).Decode(&search)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	emails := make([]Email, 0)

	for _, hit := range search.Hits.Hits {
		email := Email{
			MessageID: hit.Source.MessageID,
			Date:      hit.Source.Date,
			From:      hit.Source.From,
			To:        hit.Source.To,
			Subject:   hit.Source.Subject,
			Body:      hit.Source.Body,
		}

		emails = append(emails, email)
	}

	return emails, err

}
