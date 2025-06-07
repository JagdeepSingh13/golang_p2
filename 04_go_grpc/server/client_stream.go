package main

import (
	"io"
	"log"

	pb "github.com/JagdeepSingh13/04_go_grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		messages = append(messages, "hello", req.Name)
	}
}
