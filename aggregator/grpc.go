package main

import (
	"context"
	"toll-calculator/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> type.Distance
// Webpack -> types.WebPack -> types.Distance

// business layer -> business layer type (main type that everyone needs to be converted to)

func (s *GRPCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		Value: float64(req.ObuID),
		OBUID: int(req.ObuID),
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
