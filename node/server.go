package node

import (
	"context"
	"fmt"

	"github.com/0xvesion/gocoin/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb.NodeServer = (*NodeServer)(nil)

type NodeServer struct {
	pb.UnimplementedNodeServer
	chain   Chain
	workers []chan interface{}
}

func (server *NodeServer) StreamTasks(_ *emptypb.Empty, stream pb.Node_StreamTasksServer) error {
	onUpdate := make(chan interface{})
	server.workers = append(server.workers, onUpdate)

	go func() { onUpdate <- nil }()

	for {
		<-onUpdate

		c := server.chain.Current

		task := pb.Task{
			Difficulty: uint32(server.chain.Difficulty),
			Parent:     c.Parent,
			Timestamp:  c.Timestamp,
			Data:       c.Data,
		}

		stream.Send(&task)
	}
}

func (server *NodeServer) SubmitTask(_ context.Context, submission *pb.Submission) (*emptypb.Empty, error) {
	ok := server.chain.AdvanceBlock(submission.GetNonce())

	if !ok {
		return nil, fmt.Errorf("invalid nonce: %v", submission.Nonce)
	}

	return &emptypb.Empty{}, nil
}

func NewServer() pb.NodeServer {
	s := &NodeServer{
		chain: *NewChain(),
	}

	s.chain.OnUpdate = func() {
		for _, w := range s.workers {
			w <- nil
		}
	}

	return s
}
