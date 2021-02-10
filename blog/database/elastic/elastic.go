package elastic

import (
	"context"
	"fmt"
	"sync"
	"time"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/model"
)

var (
	// Database is variable for elk
	Database elasticsearch = &elk{}
	// sync
	once sync.Once
)

// elasticsearch interface
type elasticsearch interface {
	Connect(conf *conf.GlobalConfiguration, logger el.Logger) (err error)
	Store(ctx context.Context, index string, docType string, id string, data interface{}) (*el.IndexResponse, error)
	Index(ctx context.Context, index string, docType string, doc interface{}) (*el.IndexResponse, error)
	IndexExists(ctx context.Context, index string) error
	Delete(ctx context.Context, index string, docType string, id string) (*el.DeleteResponse, error)
	Get(ctx context.Context, index string, docType string, id string) (*el.GetResult, error)
	Search(ctx context.Context, index string, query el.Query) (*el.SearchResult, error)
	Update(ctx context.Context, index string, docType string, id string, doc interface{}) (*el.UpdateResponse, error)
	BuildQuery(must, should, not, filter []*model.Query) (*el.BoolQuery, error)
	SetClient(client *el.Client)
	GetClient() *el.Client
}

// elk struct
type elk struct {
	client *el.Client
}

// SetClient method
// for set new client for elk
func (e *elk) SetClient(client *el.Client) {
	e.client = client
}

// GetClient method
func (e *elk) GetClient() *el.Client {
	return e.client
}

// Connect to elasticsearch service
func (e *elk) Connect(conf *conf.GlobalConfiguration, logger el.Logger) error {
	client, err := el.NewClient(
		el.SetURL(conf.ELK.URLs...),
		el.SetBasicAuth(conf.ELK.Username, conf.ELK.Password),
		el.SetErrorLog(logger),
		el.SetInfoLog(logger),
		el.SetHealthcheck(true),
		el.SetHealthcheckInterval(time.Second*50),
	)
	if err != nil {
		return err
	}

	e.SetClient(client)

	return err
}

func (e *elk) Store(ctx context.Context, index string, docType string, id string, data interface{}) (*el.IndexResponse, error) {

	put, err := e.client.Index().
		Index(index).
		Type(docType).
		Id(fmt.Sprintf("%s-%s", index, id)).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		logrus.Warn("Error in index new document into elasticsearch: %s", err)
		return nil, err
	}
	return put, nil
}

// Index meth
func (e *elk) Index(ctx context.Context, index string, docType string, doc interface{}) (*el.IndexResponse, error) {

	in, err := e.client.Index().Index(index).BodyJson(doc).Type(docType).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return nil, err
	}

	return in, nil
}

func (e *elk) IndexExists(ctx context.Context, index string) error {
	exists, err := e.client.IndexExists(index).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return err
	}
	if !exists {
		logrus.Warn(fmt.Sprintf("Index does not exist yet: %s", index), err)
		return fmt.Errorf(fmt.Sprintf("Index does not exist yet: %s", index))
	}

	return nil

}

func (e *elk) Delete(ctx context.Context, index string, docType string, id string) (*el.DeleteResponse, error) {
	resp, err := e.client.Delete().
		Id(fmt.Sprintf("%s-%s", index, id)).
		Index(index).
		Type(docType).
		Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return nil, err
	}

	return resp, nil
}

// Get meth
// index,doctype and id for get
func (e *elk) Get(ctx context.Context, index string, docType string, id string) (*el.GetResult, error) {
	elk, err := e.client.Get().
		Id(fmt.Sprintf("%s-%s", index, id)).
		Index(index).
		Type(docType).
		Do(ctx)
	if err != nil {
		logrus.Warn(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, err
	}
	return elk, nil
}

func (e *elk) Search(ctx context.Context, index string, query el.Query) (*el.SearchResult, error) {
	elk, err := e.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		logrus.Warn(fmt.Sprintf("error in Search data from elastic %v", query), err)
		return nil, fmt.Errorf(fmt.Sprintf("error in Search from elastic %s", err))
	}
	return elk, nil
}

func (e *elk) BuildQuery(must, should, not, filter []*model.Query) (*el.BoolQuery, error) {
	search := el.NewBoolQuery()

	// Create Must Query for elasticsearch
	for _, m := range must {
		search.Must(el.NewMatchQuery(fmt.Sprintf("%s.keyword", m.Name), m.Value))
	}

	// Create Should Query for elasticsearch
	for _, s := range should {
		search.Should(el.NewMatchQuery(fmt.Sprintf("%s.keyword", s.Name), s.Value))
	}

	// Create MustNot Query for elasticsearch
	for _, n := range not {
		search.MustNot(el.NewMatchQuery(fmt.Sprintf("%s.keyword", n.Name), n.Value))
	}

	// Create Filter Query for elasticsearch
	for _, f := range filter {
		search.Filter(el.NewMatchQuery(fmt.Sprintf("%s.keyword", f.Name), f.Value))
	}

	if _, err := search.Source(); err != nil {
		panic(err)
	}

	return search, nil
}

// Update meth
// index,Type,id and script query for Update
func (e *elk) Update(ctx context.Context, index string, docType string, id string, doc interface{}) (*el.UpdateResponse, error) {
	elk, err := e.client.Update().
		Index(index).
		Type(docType).
		Id(fmt.Sprintf("%s-%s", index, id)).
		Doc(doc).
		Do(ctx)
	if err != nil {
		logrus.Warn(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, fmt.Errorf(fmt.Sprintf("error in get from elastic: %s", err))
	}
	return elk, nil
}
