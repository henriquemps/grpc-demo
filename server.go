package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type server struct {
	pb.UnimplementedProcessServiceServer
}

func (s *server) CallUnary(ctx context.Context, e *pb.Empty) (*pb.Process, error) {

	return &pb.Process{
		Id:   1,
		Time: "Tempo qualquer",
	}, nil
}

func (s *server) CallSStream(e *pb.Empty, stream pb.ProcessService_CallSStreamServer) error {

	for i := 0; i < 10; i++ {

		t := rand.Intn(5)

		time.Sleep(time.Duration(t) * time.Second)

		product := &pb.Process{
			Id:   int32(i),
			Time: strconv.Itoa(t),
		}

		stream.Send(product)
	}

	return nil
}

func (s *server) CallCStream(stream pb.ProcessService_CallCStreamServer) error {

	response := &pb.ListProcess{}

	for {
		p, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(response)
		}

		if err != nil {
			log.Fatalf("CallCStream failed: %v", err)
		}

		response.Process = append(response.Process, &pb.Process{
			Id:    p.GetId(),
			Time:  "Tempo zero",
			Label: "Processo: " + strconv.Itoa(int(p.GetId())) + " - OK",
		})
	}
}

func (s *server) CallBStream(stream pb.ProcessService_CallBStreamServer) error {

	for {
		p, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("CallCStream failed: %v", err)
		}

		err = stream.Send(&pb.Process{
			Id:    p.GetId(),
			Time:  p.GetTime(),
			Label: "Processo: " + strconv.Itoa(int(p.GetId())) + " - OK",
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	listen, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProcessServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
