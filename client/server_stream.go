package main

import (
	"context"
	"io"
	"log"

	pb "github.com/lyfeyvutha/grpc-demo-id/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: &v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}


/*

base) lyfey@Lyfeys-MacBook-Pro client % go run *.go
2023/05/15 15:39:27 Hello
(base) lyfey@Lyfeys-MacBook-Pro client % go run *.go
2023/05/15 15:53:30 Streaming started
2023/05/15 15:53:30 message:"HelloLyfey"
2023/05/15 15:53:32 message:"HelloLeBron"
2023/05/15 15:53:34 message:"HelloLeonard"
2023/05/15 15:53:36 Streaming finished
(base) lyfey@Lyfeys-MacBook-Pro client % go run *.go
2023/05/15 16:12:01 Client streaming started
2023/05/15 16:12:01 Send the request with name: Lyfey
2023/05/15 16:12:03 Send the request with name: LeBron
2023/05/15 16:12:05 Send the request with name: Leonard
2023/05/15 16:12:07 CLient streaming finished
2023/05/15 16:12:07 [Hello Lyfey Hello LeBron Hello Leonard]
(base) lyfey@Lyfeys-MacBook-Pro client % go run *.go
2023/05/15 16:29:17 Bidirectional Streaming Started
2023/05/15 16:29:17 message:"HelloLyfey"
2023/05/15 16:29:19 message:"HelloLeBron"
2023/05/15 16:29:21 message:"HelloLeonard"
2023/05/15 16:29:23 Bidirectional streaming finished

*/