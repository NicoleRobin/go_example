package main

import (
	"fmt"
	"go_example/packages/net/rpc/arith"
	"log"
	"net/rpc"
)

func synchronous_call(client *rpc.Client) {
	args := &arith.Args{7, 8}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Synchronous arith: %d*%d=%d\n", args.A, args.B, reply)
}

func asynchronous_call(client *rpc.Client) {
	args := &arith.Args{7, 8}
	quotient := new(arith.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Method:%s\n", replyCall.ServiceMethod)
	replyArgs := replyCall.Args.(*arith.Args)
	fmt.Printf("Args.A:%d, Args.B:%d\n", replyArgs.A, replyArgs.B)
	replyReply := replyCall.Reply.(*arith.Quotient)
	fmt.Printf("Reply.Quo:%d, Reply.Rem:%d\n", replyReply.Quo, replyReply.Rem)
	fmt.Printf("quotient.Quo:%d, quotient.Rem:%d\n", quotient.Quo, quotient.Rem)
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	synchronous_call(client)
	asynchronous_call(client)
}
