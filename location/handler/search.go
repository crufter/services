package handler

import (
	"log"

	"github.com/micro/services/location/dao"
	"github.com/micro/services/location/domain"
	loc "github.com/micro/services/location/proto/location"

	"golang.org/x/net/context"
)

func (l *Location) Search(ctx context.Context, req *loc.SearchRequest, rsp *loc.SearchResponse) error {
	log.Print("Received Location.Search request")

	entity := &domain.Entity{
		Latitude:  req.Center.Latitude,
		Longitude: req.Center.Longitude,
	}

	entities := dao.Search(req.Type, entity, req.Radius, int(req.NumEntities))

	for _, e := range entities {
		rsp.Entities = append(rsp.Entities, e.ToProto())
	}

	return nil
}
