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

func (t TwoPCService) Create2PhaseCommit(context context.Context, request *protos.TwoPCRequest) (*protos.TwoPCReply, error) {
	leaseId, err := t.impl.QDB().Create2PhaseCommitWithLease(context, request.Txid)
	if err != nil {
		return nil, err
	}

	return &protos.TwoPCReply{
		Lease: leaseId,
	}, nil
}

func (t TwoPCService) Update2PhaseCommitStatus(context context.Context, request *protos.TwoPCRequest) (*emptypb.Empty, error) {
	err := t.impl.QDB().Update2PhaseCommit(context, request.Txid, request.Status)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

var _ protos.TwoPCServiceServer = &TwoPCService{}
