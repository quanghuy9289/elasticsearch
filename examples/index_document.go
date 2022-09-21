package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

func CreateIndex(es *elasticsearch.Client, index string) error {
	isExist, err := CheckIndexExist(es, index)
	if err != nil {
		return err
	}
	if !isExist {
		res, err := es.Indices.Create(index)
		if err != nil {
			return fmt.Errorf("cannot create index: %w", err)
		}
		if res.IsError() {
			return fmt.Errorf("error in index creation response: %s", res.String())
		}
	}
	return nil
}

func CheckIndexExist(es *elasticsearch.Client, index string) (bool, error) {
	res, err := es.Indices.Exists([]string{index})
	if err != nil {
		return false, fmt.Errorf("cannot check index existence: %w", err)
	}
	if res.StatusCode == 200 {
		return true, nil
	}
	if res.StatusCode != 404 {
		return false, fmt.Errorf("error in index existence response: %s", res.String())
	}

	return false, nil
}

func AddDocument(es *elasticsearch.Client, index string, post *Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.CreateRequest{
		Index:      index,
		DocumentID: post.ID,
		Body:       bytes.NewReader(body),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return fmt.Errorf("conflict")
	}

	if res.IsError() {
		return fmt.Errorf("insert: response: %s", res.String())
	}

	return nil
}

func UpdateDocument(es *elasticsearch.Client, index string, post *Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("update: marshall: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: post.ID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, body))),
		// Body: bytes.NewReader(body),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("update: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("not found")
	}

	if res.IsError() {
		return fmt.Errorf("update: response: %s", res.String())
	}

	return nil
}

func DeleteDocument(es *elasticsearch.Client, index string, id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("delete: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("not found")
	}

	if res.IsError() {
		return fmt.Errorf("delete: response: %s", res.String())
	}

	return nil
}
