package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
	"kalra.com/goProjects/contract"
)

func main() {
	var readback contract.ReqHeartbeat
	c := &contract.ReqHeartbeat{Id: 9, Done: true}
	flag.Parse()
	if flag.NArg() < 1 {
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("ReqHeatbeat(default) := %v", c))
	b, e := proto.Marshal(c)
	if e != nil {
		fmt.Println("Failed with error: ", e)
	} else {
		fmt.Println("marshalled data:", b)
	}
	if e := proto.Unmarshal(b, &readback); e != nil {
		fmt.Println("failed ot read back data from marshalled/raw data", e)
	}
	fmt.Println("read back value ->", fmt.Sprintf("data back %v", readback))
}
