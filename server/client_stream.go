package main

import (
	"io"
	"log"
	"strings"

	pb "github.com/lyfeyvutha/grpc-demo-id/proto"
)

func (s *helloServer) SayHelloCLientStreaming(stream pb.GreetService_SayHelloCLientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
			
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Uppercase: ", strings.ToUpper(req.Name))
	}
}

/*

go run *.go
2023/05/15 16:57:26 Client streaming started
2023/05/15 16:57:26 Send the request with name: Lyfey
2023/05/15 16:57:28 Send the request with name: LeBron
2023/05/15 16:57:30 Send the request with name: Leonard
2023/05/15 16:57:32 CLient streaming finished
2023/05/15 16:57:32 [Uppercase:  LYFEY Uppercase:  LEBRON Uppercase:  LEONARD]

*/