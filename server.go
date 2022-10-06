package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
    "github.com/mrphil2105/ntpSimulation/proto"
)

type server struct {
    ntpSimulation.UnimplementedNtpServer
}

func (s *server) GetTime(ctx context.Context, in *ntpSimulation.SendTime) (*ntpSimulation.SendTime, error) {
    t2 := timestamppb.Now()
    log.Printf("\nt1: %v\nt2: %v", in.Time.AsTime(), t2.AsTime())
    return &ntpSimulation.SendTime{Time: t2}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()

    ntpSimulation.RegisterNtpServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
