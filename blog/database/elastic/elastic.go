package elastic

import (
	"context"
	"fmt"
	"sync"

	el "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/seeder/seeder"
	"github.com/emadghaffari/virgool/blog/conf"
)

var (
	// Database is variable for elk
	Database elasticsearch = &elk{}
	// sync
	once sync.Once
)

// elasticsearch interface
type elasticsearch interface {
	Connect(conf *conf.GlobalConfiguration,logger el.Logger) (err error)
	Store(ctx context.Context, index string, data interface{}) (*el.IndexResponse, error)
	SetClient(client *el.Client)
	GetClient() *el.Client
}

// elk struct
type elk struct {
	client *el.Client
}

// SetClient method
// for set new client for elk
func (e *elk) SetClient(client *el.Client){
	e.client = client
}

// GetClient method
func (e *elk) GetClient() *el.Client {
	return e.client
}

// Connect to elasticsearch service
func (e *elk) Connect(conf *conf.GlobalConfiguration,logger el.Logger) (err error) {
	client, err := el.NewClient(
		el.SetURL(conf.ELK.URLs...),
		el.SetBasicAuth(conf.ELK.Username,conf.ELK.Password),
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
func (e *elk) Index(ctx context.Context,index string,docType string,doc interface{}) (*el.IndexResponse,error)  {
	
	in,err := e.client.Index().Index(index).BodyJson(doc).Type(docType).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return nil,err
	}
	
	return in,nil
}

func (e *elk) IndexExists(ctx context.Context,index string) (error) {
	exists,err := e.client.IndexExists(index).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return err
	}
	if !exists{
		logrus.Warn(fmt.Sprintf("Index does not exist yet: %s", index), err)
		return fmt.Errorf(fmt.Sprintf("Index does not exist yet: %s", index))
	}

	return nil
	
}


func (e *elk) Delete(ctx context.Context,index string, docType string, id string) (*el.DeleteResponse,error) {
	resp,err := e.client.Delete().Id(id).Index(index).Type(docType).Do(ctx)
	if err != nil {
		logrus.Warn(err.Error())
		return nil,err
	}

	return resp,nil
}

