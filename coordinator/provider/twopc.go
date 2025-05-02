package provider

import (
	"context"

	"github.com/pg-sharding/spqr/coordinator"
	protos "github.com/pg-sharding/spqr/pkg/protos"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (t TwoPCService) Create2PhaseCommit(context.Context, *protos.TwoPCRequest) (*protos.CreateTwoPCReply, error) {
	// todo доделать, чтобы координатор сохранях инфу в etcd
	return &protos.CreateTwoPCReply{
		Lease: "",
	}, nil
}

func (t TwoPCService) Update2PhaseCommitStatus(context.Context, *protos.TwoPCRequest) (*emptypb.Empty, error) {
	// todo доделать, чтобы координатор сохранях инфу в etcd
	return nil, nil
}

var _ protos.TwoPCServiceServer = &TwoPCService{}
