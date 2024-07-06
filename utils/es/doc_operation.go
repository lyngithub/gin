package es

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/updatebyquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"strings"
)

const (
	DefaultVersionType  = "external"
	VersionTypeInternal = "internal"
	DefaultRefresh      = "false"
	RefreshWaitFor      = "wait_for"
	RefreshTrue         = "true"
	DefaultScriptLang   = "painless"
	Bulk
	OptypeCreate = "create"
	OptypeIndex  = "index"
	OptypeUpsert = "upsert"
	OptypeUpdate = "update"
	OptypeDelete = "delete"
)

type BulkIndexerOnFailure func(context.Context, esutil.BulkIndexerItem, esutil.BulkIndexerResponseItem, error)

// 新建文档
func (c *Client) Insert(ctx context.Context, indexName, id, routing, action string, doc interface{}) (*index.Response, error) {
	opType := optype.OpType{Name: action}
	indexService := c.Client.Index(indexName).OpType(opType)
	if len(id) > 0 {
		indexService.Id(id)
	}
	if len(routing) > 0 {
		indexService.Routing(routing)
	}
	//Refresh setting
	//false 不采取任何与刷新有关的行动。这个请求所做的改变将在请求返回后的某个时间点变得可见
	//true; 在操作发生后立即刷新相关的主分片和副本分片（而不是整个索引），以便更新的文档立即出现在搜索结果中.对性能影响最大
	//wait_for 在操作响应之前，等待请求所做的改变通过刷新而变得可见。这并不强迫立即进行刷新，而是等待刷新的发生。
	//Elasticsearch每隔index.refresh_interval（默认为一秒）就会自动刷新发生变化的分片
	return indexService.Request(doc).Refresh(refresh.False).Do(ctx)

}

// 批量的方式新建文档，后台提交
func (c *Client) BulkInsert(ctx context.Context, indexName, id, routing, jsonDoc string, onFailure BulkIndexerOnFailure) error {
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	if len(id) > 0 {
		bulkIndexerItem.DocumentID = id
	}
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = OptypeCreate
	bulkIndexerItem.RetryOnConflict = esapi.IntPtr(3)
	bulkIndexerItem.Body = strings.NewReader(jsonDoc)
	bulkIndexerItem.OnFailure = onFailure

	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

// 批量的方式新建文档(覆盖写)，后台提交
func (c *Client) BulkIndex(ctx context.Context, indexName, id, routing, jsonDoc string, onFailure BulkIndexerOnFailure) error {
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	if len(id) > 0 {
		bulkIndexerItem.DocumentID = id
	}
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = OptypeIndex
	bulkIndexerItem.RetryOnConflict = esapi.IntPtr(3)
	bulkIndexerItem.Body = strings.NewReader(jsonDoc)
	bulkIndexerItem.OnFailure = onFailure

	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

func (c *Client) BulkUpdate(ctx context.Context, indexName, id, routing string, update map[string]interface{}, onFailure BulkIndexerOnFailure) error {
	if len(id) == 0 {
		return errors.New("_doc id is required")
	}
	updateDoc := map[string]interface{}{
		"doc": update,
	}
	jsonDoc, err := json.Marshal(updateDoc)
	if err != nil {
		return err
	}
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	bulkIndexerItem.DocumentID = id
	bulkIndexerItem.Body = strings.NewReader(string(jsonDoc))
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = OptypeUpdate
	bulkIndexerItem.RetryOnConflict = esapi.IntPtr(3)
	bulkIndexerItem.OnFailure = onFailure

	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

// seqNo primaryTerm 用于更新文档，通过SeqNo和PrimaryTerm来保证更新操作的顺序性，避免冲突
func (c *Client) BulkInsertWithSeqNo(ctx context.Context, id, routing, indexName, action, jsonDoc string, seqNo, primaryTerm *int64, onFailure BulkIndexerOnFailure) error {
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	//id为空时，自动生成id，在一定程度上可以提高写性能，适合日志场景。不合适find then update的场景
	if len(id) > 0 {
		bulkIndexerItem.DocumentID = id
	}
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = action
	bulkIndexerItem.IfSeqNo = seqNo
	bulkIndexerItem.IfPrimaryTerm = primaryTerm
	bulkIndexerItem.VersionType = DefaultVersionType
	bulkIndexerItem.Body = strings.NewReader(jsonDoc)
	bulkIndexerItem.OnFailure = onFailure
	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

func (c *Client) Delete(ctx context.Context, indexName, id, routing string) (*delete.Response, error) {
	if len(id) == 0 {
		return nil, errors.New("_doc id is required")
	}
	deleteService := c.Client.Delete(indexName, id)
	if len(routing) > 0 {
		deleteService.Routing(routing)
	}
	return deleteService.Do(ctx)
}

func (c *Client) DeleteWithSeqNo(ctx context.Context, indexName, id, routing, seqNo, primaryTerm string) (*delete.Response, error) {
	deleteService := c.Client.Delete(indexName, id).IfSeqNo(seqNo).IfPrimaryTerm(primaryTerm)
	if len(routing) > 0 {
		deleteService.Routing(routing)
	}
	return deleteService.Do(ctx)
}

// 删除过程中如果命中的文档被修改会导致删除出现版本冲突，为了彻底删除文档，需要指定ProceedOnVersionConflict
func (c *Client) DeleteByQuery(ctx context.Context, indexName, routing, preference string, query *types.Query) (*deletebyquery.Response, error) {
	deleteService := c.Client.DeleteByQuery(indexName).Conflicts(conflicts.Proceed)
	if len(routing) > 0 {
		deleteService.Routing(routing)
	}
	if len(preference) > 0 {
		deleteService.Preference(preference)
	}
	return deleteService.Query(query).Do(ctx)
}

func (c *Client) BulkDelete(ctx context.Context, indexName, id, routing string, onFailure BulkIndexerOnFailure) error {
	if len(id) == 0 {
		return errors.New("_doc id is required")
	}
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	if len(id) > 0 {
		bulkIndexerItem.DocumentID = id
	}
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = OptypeDelete
	bulkIndexerItem.RetryOnConflict = esapi.IntPtr(3)
	bulkIndexerItem.OnFailure = onFailure

	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

func (c *Client) BulkDeleteWithSeqNo(ctx context.Context, indexName, id, routing string, sqNo, primaryTerm *int64, onFailure BulkIndexerOnFailure) error {
	if len(id) == 0 {
		return errors.New("_doc id is required")
	}
	bulkIndexerItem := esutil.BulkIndexerItem{}
	bulkIndexerItem.Index = indexName
	bulkIndexerItem.VersionType = DefaultVersionType
	bulkIndexerItem.IfSeqNo = sqNo
	bulkIndexerItem.IfPrimaryTerm = primaryTerm
	bulkIndexerItem.DocumentID = id
	if len(routing) > 0 {
		bulkIndexerItem.Routing = routing
	}
	bulkIndexerItem.Action = OptypeDelete
	bulkIndexerItem.RetryOnConflict = esapi.IntPtr(3)
	bulkIndexerItem.OnFailure = onFailure

	return c.BulkProcessor.Add(ctx, bulkIndexerItem)
}

func (c *Client) Update(ctx context.Context, indexName, id, routing string, update map[string]interface{}) (*update.Response, error) {
	if len(id) == 0 {
		return nil, errors.New("_doc id is required")
	}
	updateService := c.Client.Update(indexName, id)
	if len(routing) > 0 {
		updateService.Routing(routing)
	}
	return updateService.Doc(update).Do(ctx)

}

func (c *Client) Upsert(ctx context.Context, indexName, id, routing string, update map[string]interface{}, upsertDoc interface{}) (*update.Response, error) {
	if len(id) == 0 {
		return nil, errors.New("_doc id is required")
	}
	updateService := c.Client.Update(indexName, id).Upsert(upsertDoc)
	if len(routing) > 0 {
		updateService.Routing(routing)
	}
	return updateService.Doc(update).Do(ctx)

}

func (c *Client) UpdateByQuery(ctx context.Context, indexName string, routing, preference string, query *types.Query, script types.Script) (*updatebyquery.Response, error) {
	updateService := c.Client.UpdateByQuery(indexName).Conflicts(conflicts.Abort).Script(script)
	if len(routing) > 0 {
		updateService.Routing(routing)
	}
	if len(preference) > 0 {
		updateService.Preference(preference)
	}
	return updateService.Query(query).Do(ctx)

}
