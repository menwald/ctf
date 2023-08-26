package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/menwald/ctf/common/config"
	"github.com/menwald/ctf/engine/internal/apihandler"
	"github.com/menwald/ctf/proto/github.com/menwald/ctf/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Basic initialization like logging and config.yaml
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	flag.Parse()
	cfg, err := config.Load()
	if err != nil {
		return
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.GetString("logging.level") == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Info().Str("configFile", cfg.Filename()).Msg("starting application")

	svrAddr := cfg.GetString("grpc.address")
	lis, err := net.Listen("tcp", svrAddr)
	if err != nil {
		log.Err(err)
		return
	}
	grpclog.V(2)
	gsvr := grpc.NewServer(
		//grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
		grpc.MaxRecvMsgSize(1024*1024),
		grpc.MaxSendMsgSize(1024*1024),
	)
	mapHandler := apihandler.GameMap{}
	proto.RegisterMapAPIServer(gsvr, &mapHandler)
	stop := make(chan os.Signal, 1)
	// Serve gRPC Server
	log.Info().Str("serverAddress", svrAddr).Msg("serving gRPC")
	go func() {
		err := gsvr.Serve(lis)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	/*
		addr := cfg.GetString("grpc.address")
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Err(err)
			return
		}
		grpclog.V(2)
		gsvr := grpc.NewServer(
			//grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
			grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
			grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
			grpc.MaxRecvMsgSize(1024*1024),
			grpc.MaxSendMsgSize(1024*1024),
		)
		//svr := grpc.Server
		log.Info().Str("address", addr).Msg("starting grpc server")
		err = gsvr.Serve(lis)
		if err != nil {
			log.Err(err)
			return
		}
	*/
	//cancel()
	signal.Notify(stop, os.Interrupt)

	<-stop

	log.Info().Msg("shutting down the server")
	cancel()
	gsvr.GracefulStop()
	log.Info().Msg("server gracefully stopped")
}
