package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type DocumentById struct {
	Source interface{} `json:"_source"`
}

type HitItem struct {
	Index      string                 `json:"_index"`
	Type       string                 `json:"_type"`
	DocumentId string                 `json:"_id"`
	Score      float64                `json:"_score"`
	Source     map[string]interface{} `json:"_source"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type SearchHits struct {
	Total    Total     `json:"total"`
	MaxScore float64   `json:"max_score"`
	Hits     []HitItem `json:"hits"`
}

type SearchResponse struct {
	Took float64    `json:"took"`
	Hits SearchHits `json:"hits"`
}

func ExecuteSearch(es *elasticsearch.Client, index string, query map[string]interface{}) (*SearchResponse, error) {
	if query == nil {
		return nil, fmt.Errorf("query is not valid")
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, fmt.Errorf("search error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	if res.IsError() {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	var response SearchResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return &response, nil
}
