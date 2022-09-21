# check healthy of elasticsearch server
GET /_cat/health?v

# get all indexes
GET /_cat/indices?pretty&v

# create index student
PUT /student

# delete index
# note that delete index, all docs will be deleted too
DELETE /product

# add data to index (add doc)
POST /student/_doc 
{
  "name": "ten 1",
  "age": 20
}

# update doc in student index
PUT /student/_doc/n6_SNYMB2r_H7pTMl3IA
{
  "name": "ten 1",
  "age": 21,
  "mon_hoc": ["Toan","Ly","Hoa"]
}

# get data of index with id
GET /student/_doc/n6_SNYMB2r_H7pTMl3IA

# get all data of index
GET /student/_search

# add bulk data to index
POST /_bulk 
{"index": {"_index":"student"}}
{"name": "ten 2", "age": 22}
{"index": {"_index":"student"}}
{"name": "ten 3","age": 23}

# add bulk data from json file
curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/_bulk?pretty&refresh" --data-binary "@accounts.json"
