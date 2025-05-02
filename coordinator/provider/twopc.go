package provider

import (
	"context"
	"github.com/pg-sharding/spqr/coordinator"
	protos "github.com/pg-sharding/spqr/pkg/protos"
)

func NewTwoPCServer(impl coordinator.Coordinator) *TwoPCService {
	return &TwoPCService{
		impl: impl,
	}
}

type TwoPCService struct {
	protos.UnimplementedTwoPCServiceServer

	impl coordinator.Coordinator
}

func (t TwoPCService) Create2PhaseCommit(context.Context, *protos.CreateTwoPCReply) (*protos.CreateTwoPCReply, error) {
	// todo доделать, чтобы координатор сохранях инфу в etcd
	return &protos.CreateTwoPCReply{
		Err: "",
	}, nil
}

var _ protos.TwoPCServiceServer = &TwoPCService{}
