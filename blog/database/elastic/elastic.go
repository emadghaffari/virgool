package elastic

import (
	"context"
	"sync"

	el "github.com/olivere/elastic/v7"

	"github.com/emadghaffari/virgool/blog/conf"
)

var (
	// Database is variable for elk
	Database elasticsearch = &elk{}
	// sync
	once sync.Once
)

type elasticsearch interface {
	Connect(config *conf.GlobalConfiguration) error
	GetClient() *el.Client
}

type elk struct {
	client *el.Client
}

func (e *elk) GetClient() *el.Client {
	return e.client
}

func (e *elk) Connect(conf *conf.GlobalConfiguration) (err error) {
	e.client, err = el.NewClient(el.SetURL(conf.ELK.URLs...))
	if err != nil {
		return err
	}

	ps := e.client.Ping(conf.ELK.URLs[0])
	if _, _, err = ps.Do(context.Background()); err != nil {
		return err
	}

	return err
}
