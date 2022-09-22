package main

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// es, err := elasticsearch.NewDefaultClient()
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Username: "elastic",
		Password: "Admin123",
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// create index
	index := "posts"
	post := &Post{
		ID:        "1",
		Title:     "Post 1",
		Text:      "This is a document",
		Tags:      []string{"#post"},
		CreatedAt: time.Now().Unix(),
	}

	IndexDocOperation(es, index, post)

	SearchOperation(es, index, post)
}

func IndexDocOperation(es *elasticsearch.Client, index string, post *Post) {
	CreateIndex(es, index)

	// add document
	AddDocument(es, index, post)

	// update document
	post.Tags = append(post.Tags, "#IT")
	UpdateDocument(es, index, post)

	// delete document
	//DeleteDocument(es, index, post.ID)
}

func SearchOperation(es *elasticsearch.Client, index string, post *Post) error {
	res, err := GetPostById(es, index, post.ID)
	if err != nil {
		return err
	}

	fmt.Print(res)

	posts, err := GetPostByTitle(es, index, post.Title)
	if err != nil {
		return err
	}
	fmt.Print(posts)

	// search bank account with balance greater than $10,000
	bankAccs, err := GetBankAccountGteBalance(es, "bank", 10000)
	if err != nil {
		return err
	}
	fmt.Print(bankAccs)

	return nil
}
