package apihandler

import (
	"context"
	"time"

	"github.com/menwald/ctf/proto/github.com/menwald/ctf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GameMap struct {
	proto.UnimplementedMapAPIServer
}

func (g *GameMap) GetEntireMap(ctx context.Context, req *proto.GetEntireMapRequest) (*proto.GetEntireMapResponse, error) {
	var resp proto.GetEntireMapResponse
	resp.Timestamp = timestamppb.New(time.Now())
	return &resp, nil
}
