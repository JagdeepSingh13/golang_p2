package main

import (
	"context"
	"log"
	"time"

	pb "github.com/JagdeepSingh13/04_go_grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending: %v", err)
		}

		log.Printf("sent the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()

	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}
	log.Printf("%v", res.Message)
}
