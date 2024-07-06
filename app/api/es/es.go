package es

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"xx/utils/es"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var indexCreateJson = `
{
  "settings": {
    "number_of_shards": 3,
    "number_of_replicas": 1
  },
  "mappings": {
  "dynamic": "strict",
    "properties": {
      "id": {
        "type": "keyword",
        "doc_values": false,
        "norms": false,
        "similarity": "boolean"
      },
      "name": {
        "type": "text"
      },
      "age":{
        "type": "short"
      }
    }
  }
}
`

func Add(ctx *gin.Context) {
	indexName := "user111"
	ct := context.Background()
	esClient := es.GetClient(es.DefaultClient)
	err := esClient.CreateIndex(ct, indexName, indexCreateJson, true)
	if err != nil {
		log.Println(err)
	}
}
