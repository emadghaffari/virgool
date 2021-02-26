package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/opentracing/opentracing-go"

	el "github.com/emadghaffari/virgool/club/database/elastic"
	"github.com/emadghaffari/virgool/club/database/mysql"
	"github.com/emadghaffari/virgool/club/model"
)

// ClubService describes the service.
type ClubService interface {
	// Add your methods here
	Get(ctx context.Context, id string, token string) (result string, err error)
	Index(ctx context.Context, from, size int32, filter []*model.Query, token string) (results []model.Point, err error)
}

type basicClubService struct{}

func (b *basicClubService) Get(ctx context.Context, id string, token string) (result string, err error) {
	// start get-post trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("get-point")
	defer span.Finish()

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// find the user with username or email
	point := model.Point{}
	if err := tx.Table("points").Where("user_id = ?", id).First(&point).Error; err != nil {
		span.SetTag("Error", err.Error())
		return "", fmt.Errorf(err.Error())
	}
	tx.Commit()

	return fmt.Sprintf("%d", point.Point), err
}

// NewBasicClubService returns a naive, stateless implementation of ClubService.
func NewBasicClubService() ClubService {
	return &basicClubService{}
}

// New returns a ClubService with all of the expected middleware wired in.
func New(middleware []Middleware) ClubService {
	var svc ClubService = NewBasicClubService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicClubService) Index(ctx context.Context, from int32, size int32, filter []*model.Query, token string) (results []model.Point, err error) {
	// start get-post trace
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("get-point")
	defer span.Finish()

	// build a elastic query with filters
	query, err := el.Database.BuildQuery(
		el.FilterQuery(filter),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed To Create a Search Query: %s", err)
	}

	// search by query
	resp, err := el.Database.Search(ctx, "point", query)
	if err != nil {
		return nil, fmt.Errorf("Failed To Search: %s", err)
	}

	for _, r := range resp.Hits.Hits {
		var ptn model.Point
		if err := json.Unmarshal(r.Source, &ptn); err != nil {
			return nil, fmt.Errorf("Failed To unmarshal Searched: %s", err)
		}
		results = append(results, ptn)
	}

	return results, err
}
