package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"kalra.com/goProjects/bidiGRPC"
)

type MyHBImpl struct {
	bidiGRPC.UnimplementedHeartBeatServer
}

func (m MyHBImpl) Hello(inM bidiGRPC.HeartBeat_HelloServer) error {

	for in, e := inM.Recv(); e != nil; {
		for i := 0; i < 10; i++ {
			resp := bidiGRPC.Res{Id: 10, Done: true}
			if in == nil {
				fmt.Println("in is nil")
			} else {
				resp.Id = in.Id + 1
				resp.Done = i == 9
			}
			if e := inM.Send(&resp); e != nil {
				log.Println("failed sending with error %v", e)
				return e
			}
		}
	}
	return nil
}

func main() {
	mySrvObj := MyHBImpl{}
	//myOkObj := MyOkImpl{}
	lis, e := net.Listen("tcp", ":9090")
	if e != nil {
		log.Fatalf("failed to open net port 9090, error %v", e)
		os.Exit(-1)
	}

	srv := grpc.NewServer()

	bidiGRPC.RegisterHeartBeatServer(srv, mySrvObj)
	//bidiGRPC.RegisterAreyouOkServer(srv, myOkObj)
	srv.Serve(lis)
}
