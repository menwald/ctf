package apihandler

import (
	"context"
	"time"

	pb "github.com/menwald/ctf/proto/github.com/menwald/ctf/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GameMap struct {
	pb.UnimplementedMapAPIServer
}

func (g *GameMap) GetEntireMap(ctx context.Context, req *pb.GetEntireMapRequest) (*pb.GetEntireMapResponse, error) {
	var resp pb.GetEntireMapResponse
	resp.Timestamp = timestamppb.New(time.Now())
	log.Debug().Str("response", protojson.Format(&resp)).Msg("handled GetEntireMap request")
	return &resp, nil
}
