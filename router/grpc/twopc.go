package grpc

import (
	"context"
	"errors"

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

func (t TwoPCService) Create2PhaseCommit(context context.Context, request *protos.TwoPCRequest) (*protos.CreateTwoPCReply, error) {
	return nil, errors.New("Not implemented Create2PhaseCommit on router")
}

func (t TwoPCService) Update2PhaseCommitStatus(context context.Context, request *protos.TwoPCRequest) (*emptypb.Empty, error) {
	return nil, errors.New("Not implemented Update2PhaseCommitStatus on router")
}

var _ protos.TwoPCServiceServer = &TwoPCService{}
