package elastic

import (
	"context"
	"fmt"
	"sync"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/seeder/seeder"
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
	Store(ctx context.Context, index string, data interface{}) (*el.IndexResponse, error)
	Index(ctx context.Context, index string, docType string, doc interface{}) (*el.IndexResponse, error)
	IndexExists(ctx context.Context, index string) error
	Delete(ctx context.Context, index string, docType string, id string) (*el.DeleteResponse, error)
	Get(index string, docType string, id string) (*el.GetResult, error)
	Search(index string, query el.Query) (*el.SearchResult, error)
	Update(index string, docType string, id string, script *el.Script) (*el.UpdateResponse, error)
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
func (e *elk) Connect(conf *conf.GlobalConfiguration, logger el.Logger) (err error) {
	client, err := el.NewClient(
		el.SetURL(conf.ELK.URLs...),
		el.SetBasicAuth(conf.ELK.Username, conf.ELK.Password),
		el.SetErrorLog(logger),
		el.SetInfoLog(logger),
	)
	if err != nil {
		return err
	}

	ps := e.client.Ping(conf.ELK.URLs[0])
	if _, _, err = ps.Do(context.Background()); err != nil {
		return err
	}

	e.SetClient(client)

	return err
}

func (e *elk) Store(ctx context.Context, index string, data interface{}) (*el.IndexResponse, error) {

	put, err := e.client.Index().
		Index(index).
		Type(index).
		Id(seeder.RandomHash(25)).
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
	resp, err := e.client.Delete().Id(id).Index(index).Type(docType).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return nil, err
	}

	return resp, nil
}

// Get meth
// index,doctype and id for get
func (e *elk) Get(index string, docType string, id string) (*el.GetResult, error) {
	elk, err := e.client.Get().
		Id(id).
		Index(index).
		Type(docType).
		Do(context.Background())
	if err != nil {
		logrus.Warn(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, err
	}
	return elk, nil
}

func (e *elk) Search(index string, query el.Query) (*el.SearchResult, error) {
	elk, err := e.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(context.Background())
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
func (e *elk) Update(index string, docType string, id string, script *el.Script) (*el.UpdateResponse, error) {
	elk, err := e.client.Update().
		Index(index).
		Type(docType).
		Id(id).
		Script(script).
		Do(context.Background())
	if err != nil {
		logrus.Warn(fmt.Sprintf("error in Get data from elastic %s", id), err)
		return nil, fmt.Errorf(fmt.Sprintf("error in get from elastic", err))
	}
	return elk, nil
}
