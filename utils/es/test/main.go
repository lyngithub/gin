package main

import (
	"context"
	"log"
	"xx/utils/es"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
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

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	indexName := "user"
	ctx := context.Background()
	//err := es.InitClient(es.DefaultClient, []string{"http://110.40.153.208:9200"}, "", "")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	esClient := es.GetClient(es.DefaultClient)

	// create index
	err := esClient.CreateIndex(ctx, indexName, indexCreateJson, true)
	if err != nil {
		log.Println(err)
	}

	// delete index
	//res, err := esClient.DeleteIndex(ctx, indexName)
	//if err != nil {
	//	log.Println(res, err)
	//}

	// insert data
	//user := User{
	//	Id:   1,
	//	Name: "jack ma",
	//	Age:  25,
	//}

	//res, err := esClient.Insert(ctx, indexName, strconv.FormatInt(user.Id, 10), strconv.FormatInt(user.Id, 10), es.OptypeIndex, user)
	//if err != nil {
	//	log.Println(res, err)
	//}

	//update data
	/*user1ForUpdate := map[string]interface{}{
		"name": "update jack ma",
		"age":  5011,
	}
	res, err := esClient.Update(ctx, indexName, "1", "1", user1ForUpdate)
	if err != nil {
		log.Println(err, res)
	}*/

	//user2 := User{
	//	Id:   2,
	//	Name: "tom jerry",
	//	Age:  30,
	//}

	/*
		jsonDoc, err := json.Marshal(user2)
		err = esClient.BulkInsert(ctx, indexName, strconv.FormatInt(user2.Id, 10), strconv.FormatInt(user2.Id, 10), string(jsonDoc), func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			if err != nil {
				log.Println(err, item2)
			}
		})
		if err != nil {
			log.Println(err)
		}*/

	//修改部分字段
	//user2ForBulkUpdate := map[string]interface{}{
	//	"name": "update tom jerry",
	//	"age":  50,
	//}
	//
	//err = esClient.BulkUpdate(ctx, indexName, strconv.FormatInt(user2.Id, 10), strconv.FormatInt(user2.Id, 10), user2ForBulkUpdate, func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
	//	if err != nil {
	//		log.Println(err, item2)
	//	}
	//})
	//if err != nil {
	//	log.Println(err)
	//}

	//upsert data
	//user2ForUpdate := map[string]interface{}{
	//	"name": "update tom jerry",
	//	"age":  50,
	//}
	//userForUpsert := User{
	//	Id:   4,
	//	Name: "upsert tom jerry",
	//	Age:  50,
	//}
	//
	//res, err := esClient.Upsert(ctx, indexName, strconv.FormatInt(userForUpsert.Id, 10), strconv.FormatInt(userForUpsert.Id, 10), user2ForUpdate, userForUpsert)
	//if err != nil {
	//	log.Println(err, res)
	//}

	//query := &types.Query{
	//	Term: map[string]types.TermQuery{
	//		"age": {Value: 50},
	//	},
	//}
	//script := map[string]interface{}{
	//	"inline": "ctx._source.name = params.name",
	//	"params": map[string]interface{}{
	//		"name": "update-by-script",
	//	},
	//}
	//res, err := esClient.UpdateByQuery(ctx, indexName, "2", "2", query, script)
	//if err != nil {
	//	log.Println(err, res)
	//	return
	//}

	//err = esClient.BulkDelete(ctx, indexName, strconv.FormatInt(user2.Id, 10), strconv.FormatInt(user2.Id, 10), func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
	//	if err != nil {
	//		log.Println(err, item2)
	//	}
	//})
	//if err != nil {
	//	log.Println(err)
	//}

	//get data

	//res, err := esClient.Get(ctx, indexName, "1", "1", "1", nil)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//user := User{}
	//js, err := res.Source_.MarshalJSON()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//err = json.Unmarshal(js, &user)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(user)

	//search data

	//query := &types.Query{
	//	Bool: &types.BoolQuery{
	//		Must: []types.Query{
	//			{
	//				Term: map[string]types.TermQuery{
	//					"age": {Value: 5011},
	//				},
	//			},
	//			{
	//				MatchPhrase: map[string]types.MatchPhraseQuery{
	//					"name": {Query: "jack"},
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//res, err := esClient.Query(ctx, indexName, "1", query, 0, 10, nil)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(res.Hits.Total.Value)
	//for _, hit := range res.Hits.Hits {
	//	user := User{}
	//	err = json.Unmarshal(hit.Source_, &user)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println(user)
	//}

	//query := &types.Query{
	//	MatchAll: &types.MatchAllQuery{},
	//}

	//query := &types.Query{
	//	Term: map[string]types.TermQuery{
	//		"age": {Value: 30},
	//	},
	//}
	//
	//err = esClient.ScrollQuery(ctx, indexName, "", query, 2, func(res *scroll.Response, err error) {
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println(res.Hits.Total.Value)
	//	for _, hit := range res.Hits.Hits {
	//		user := User{}
	//		err = json.Unmarshal(hit.Source_, &user)
	//		if err != nil {
	//			log.Println(err)
	//			return
	//		}
	//		log.Println(user)
	//	}
	//})
	//if err != nil {
	//	log.Println(err)
	//}

	//等待bulk操作提交
	//select {}
}
