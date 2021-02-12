package elastic

import (
	"fmt"

	el "github.com/olivere/elastic/v7"

	"github.com/emadghaffari/virgool/blog/model"
)

// Option for Build Queries
type Option func(*el.BoolQuery)


// MustQuery create a BoolQuery
func MustQuery(must []*model.Query) Option {
	return func(bq *el.BoolQuery) {
		for _, m := range must {
			bq.Must(el.NewMatchQuery(fmt.Sprintf("%s.keyword", m.Name), m.Value))
		}
	}
}

// shouldQuery create a BoolQuery
func shouldQuery(should []*model.Query) Option {
	return func(bq *el.BoolQuery) {
		for _, s := range should {
			bq.Should(el.NewMatchQuery(fmt.Sprintf("%s.keyword", s.Name), s.Value))
		}
	}
}


// MustNotQuery create a BoolQuery
func MustNotQuery(not []*model.Query) Option {
	return func(bq *el.BoolQuery) {
		for _, n := range not {
			bq.MustNot(el.NewMatchQuery(fmt.Sprintf("%s.keyword", n.Name), n.Value))
		}
	}
}

// FilterQuery create a BoolQuery
func FilterQuery(filter []*model.Query) Option {
	return func(bq *el.BoolQuery) {
		for _, f := range filter {
			bq.Filter(el.NewMatchQuery(fmt.Sprintf("%s.keyword", f.Name), f.Value))
		}
	}
}