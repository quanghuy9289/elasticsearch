package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
)

func GetPostById(es *elasticsearch.Client, index string, id string) (*Post, error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, fmt.Errorf("get document error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	if res.IsError() {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	var (
		post Post
		body DocumentById
	)
	body.Source = &post

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return &post, nil
}

func GetPostByTitle(es *elasticsearch.Client, index string, title string) ([]*Post, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": title,
			},
		},
	}

	response, err := ExecuteSearch(es, index, query)
	if err != nil {
		return nil, err
	}

	var posts []*Post
	for _, hit := range response.Hits.Hits {
		var post Post
		mapstructure.Decode(hit.Source, &post)
		posts = append(posts, &post)
	}

	return posts, nil
}
