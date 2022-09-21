
# filter data from index
GET /bank/_search
{
  "query": {"match_all": {}},
  "size": 10,
  "from": 10,
  "sort": [
    {
      "balance": {
        "order": "desc"
      }
    }
  ]
}

# filter data and get some expected fields data
GET /bank/_search
{
  "query": {"match_all": {}},
  "_source": ["account_number", "balance", "email"]
}

# filter data that address contain "mill lane"
# => return all data that contain "mill" or "lane" or both "mill lane" with highest score record on the top
GET /bank/_search
{
  "query": {"match": {
    "address": "mill lane"
  }}
}


# filter data with logic: record that has address field contain both "mill" and "lane"
# logic AND
GET /bank/_search
{
  "query": {
    "bool": {
      "must": [ # or must_not
        {"match": {
          "address": "mill"
        }},
        {"match": {
          "address": "lane"
        }}
      
      ]
    }
  }
}

# filter data with logic: record that has address field contain or "mill" or "lane" or both of them
# logic OR
GET /bank/_search
{
  "query": {
    "bool": {
      "should": [
        {"match": {
          "address": "mill"
        }},
        {"match": {
          "address": "lane"
        }}
      
      ]
    }
  }
}

# filter data with logic: record that has age=40 and state NOT in "TN"
GET /bank/_search
{
  "query": {
    "bool": {
      "must": [
        {"match": {
          "age": 40
        }}
      ],
      "must_not": [
        {"match": {
          "state": "TN"
        }}
      ]
    }
  }
}