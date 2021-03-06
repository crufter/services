package ingester

import (
	"log"

	"github.com/micro/services/location/dao"
	"github.com/micro/services/location/domain"
	proto "github.com/micro/services/location/proto"
	"golang.org/x/net/context"
)

var (
	Topic = "geo.location"
)

type Geo struct{}

func (g *Geo) Handle(ctx context.Context, e *proto.Entity) error {
	log.Printf("Saving entity ID %s", e.Id)
	dao.Save(domain.ProtoToEntity(e))
	return nil
}
