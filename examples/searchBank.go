package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
)

func GetBankAccountGteBalance(es *elasticsearch.Client, index string, balance int64) ([]*BankAccount, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"balance": map[string]interface{}{
							"gte": balance,
						},
					},
				},
			},
		},
	}

	response, err := ExecuteSearch(es, index, query)
	if err != nil {
		return nil, err
	}

	var res []*BankAccount
	for _, hit := range response.Hits.Hits {
		var item BankAccount
		mapstructure.Decode(hit.Source, &item)
		res = append(res, &item)
	}

	return res, nil
}
