package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kalra.com/goProjects/bidiGRPC"
)

func main() {

	var (
		addr   = flag.String("addr", "localhost:9090", "the address to connect to")
		id     = flag.Int64("n", 1234, "flag n")
		b      = flag.Bool("b", false, "flag end(bool)")
		epname = flag.String("e", "heartbeat", "heartbeat | alive")
	)
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to process connection %v", err)
	}
	fmt.Println("endpoint", *epname)

	c := bidiGRPC.NewHeartBeatClient(conn)
	ctx, cancel := context.WithCancel(context.Background()) //, 100*time.Second)
	defer cancel()
	req := bidiGRPC.Req{Id: *id, Stop: *b}
	r, err := c.Hello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if e := r.Send(&req); e == nil {
		for s, er := r.Recv(); ; {
			if er != nil {
				log.Printf("error: %v", er)
				return
			} else {
				log.Printf("Greeting: got back %d,%v \n", s.Id, s.Done)
			}
		}
	} else {
		log.Printf("got back error %v", e)
	}
}
