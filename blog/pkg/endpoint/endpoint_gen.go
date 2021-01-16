// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "github.com/emadghaffari/virgool/blog/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreatePostEndpoint endpoint.Endpoint
	UpdatePostEndpoint endpoint.Endpoint
	GetPostEndpoint    endpoint.Endpoint
	DeletePostEndpoint endpoint.Endpoint
	CreateTagEndpoint  endpoint.Endpoint
	GetTagEndpoint     endpoint.Endpoint
	UpdateTagEndpoint  endpoint.Endpoint
	DeleteTagEndpoint  endpoint.Endpoint
	UploadEndpoint     endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.BlogService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreatePostEndpoint: MakeCreatePostEndpoint(s),
		CreateTagEndpoint:  MakeCreateTagEndpoint(s),
		DeletePostEndpoint: MakeDeletePostEndpoint(s),
		DeleteTagEndpoint:  MakeDeleteTagEndpoint(s),
		GetPostEndpoint:    MakeGetPostEndpoint(s),
		GetTagEndpoint:     MakeGetTagEndpoint(s),
		UpdatePostEndpoint: MakeUpdatePostEndpoint(s),
		UpdateTagEndpoint:  MakeUpdateTagEndpoint(s),
		UploadEndpoint:     MakeUploadEndpoint(s),
	}
	for _, m := range mdw["CreatePost"] {
		eps.CreatePostEndpoint = m(eps.CreatePostEndpoint)
	}
	for _, m := range mdw["UpdatePost"] {
		eps.UpdatePostEndpoint = m(eps.UpdatePostEndpoint)
	}
	for _, m := range mdw["GetPost"] {
		eps.GetPostEndpoint = m(eps.GetPostEndpoint)
	}
	for _, m := range mdw["DeletePost"] {
		eps.DeletePostEndpoint = m(eps.DeletePostEndpoint)
	}
	for _, m := range mdw["CreateTag"] {
		eps.CreateTagEndpoint = m(eps.CreateTagEndpoint)
	}
	for _, m := range mdw["GetTag"] {
		eps.GetTagEndpoint = m(eps.GetTagEndpoint)
	}
	for _, m := range mdw["UpdateTag"] {
		eps.UpdateTagEndpoint = m(eps.UpdateTagEndpoint)
	}
	for _, m := range mdw["DeleteTag"] {
		eps.DeleteTagEndpoint = m(eps.DeleteTagEndpoint)
	}
	for _, m := range mdw["Upload"] {
		eps.UploadEndpoint = m(eps.UploadEndpoint)
	}
	return eps
}