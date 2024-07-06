package es

import (
	"context"
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/olivere/elastic/v7"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var clients map[string]*Client

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Client struct {
	Name           string
	Addr           []string
	QueryLogEnable bool
	Username       string
	password       string
	BulkCfg        *BulkCfg
	Client         *elasticsearch.TypedClient
	BulkProcessor  esutil.BulkIndexer
	DebugMode      bool
	//本地缓存已经创建的索引，用于加速索引是否存在的判断
	CachedIndices sync.Map
	lock          sync.Mutex
}
type BulkCfg struct {
	Workers       int
	FlushInterval time.Duration
	ActionSize    int //每批提交的文档数
	RequestSize   int //每批提交的文档大小
	AfterFunc     elastic.BulkAfterFunc
	Ctx           context.Context
}

const (
	DefaultClient      = "es-default-client"
	DefaultReadClient  = "es-default-read-client"
	DefaultWriteClient = "es-default-write-client"
)

func InitClient(clientName string, addr []string, username string, password string) error {
	if clients == nil {
		clients = make(map[string]*Client, 0)
	}
	client := &Client{
		Addr:           addr,
		QueryLogEnable: false,
		Username:       username,
		password:       password,
		CachedIndices:  sync.Map{},
		lock:           sync.Mutex{},
	}

	cfg := getBaseCfg(username, password, addr)
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return err
	}
	esBulkClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: esBulkClient, // The Elasticsearch client
		//NumWorkers:    client.BulkCfg.Workers,     // The number of worker goroutines
		//FlushBytes:    client.BulkCfg.RequestSize, // The flush threshold in bytes
		FlushInterval: 3 * time.Second, // The periodic flush interval
		ErrorTrace:    true,
		OnError: func(ctx context.Context, err error) {
			if err != nil {
				log.Printf("Bulk error : %+v", err)
			}
		},
	})
	if err != nil {
		return err
	}
	client.BulkProcessor = bi
	client.Client = esClient
	clients[clientName] = client
	return nil
}
func getBaseCfg(username, password string, addr []string) elasticsearch.Config {
	cfg := elasticsearch.Config{
		Addresses: addr,
		Username:  username,
		Password:  password,
		Transport: &http.Transport{
			//DisableKeepAlives: true,
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}
				return d.DialContext(ctx, network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			//针对es7.x+版本的https的密码连接，需要设置TLSClientConfig
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		RetryOnStatus: []int{502, 503, 504, 429}, // Add 429 to the list of retryable statuses
		RetryBackoff:  func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
		MaxRetries:    3,
		EnableMetrics: true,
	}

	return cfg
}

func getDefaultClient() *http.Client {
	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func InitClientWithCfg(clientName string, cfg elasticsearch.Config, queryLogEnable bool, bulk BulkCfg) error {
	if clients == nil {
		clients = make(map[string]*Client, 0)
	}
	client := &Client{
		Addr:           cfg.Addresses,
		QueryLogEnable: false,
		Username:       cfg.Username,
		password:       cfg.Password,
		BulkCfg:        &bulk,
		CachedIndices:  sync.Map{},
		lock:           sync.Mutex{},
	}

	client.QueryLogEnable = queryLogEnable
	client.BulkCfg = &bulk
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return err
	}
	esBulkClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        esBulkClient,       // The Elasticsearch client
		NumWorkers:    bulk.Workers,       // The number of worker goroutines
		FlushBytes:    bulk.RequestSize,   // The flush threshold in bytes
		FlushInterval: bulk.FlushInterval, // The periodic flush interval
		ErrorTrace:    true,
		OnError: func(ctx context.Context, err error) {
			if err != nil {
				log.Printf("Bulk error : %+v", err)
			}
		},
	})
	if err != nil {
		return err
	}
	client.BulkProcessor = bi
	client.Client = esClient
	clients[clientName] = client
	return nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.BulkProcessor.Close(ctx)
}

func CloseAll() {
	for _, c := range clients {
		if c != nil {
			err := c.BulkProcessor.Close(context.Background())
			if err != nil {
				log.Print("bulk close error", err)
			}
		}
	}
}

func GetClient(name string) *Client {
	if client, exist := clients[name]; exist {
		return client
	}
	log.Print("call init", name, "before !!!")
	return nil
}
