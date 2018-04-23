package main

import (
	"go_example/packages/net/rpc/arith"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(arith.Arith)
	// 首先注册该对象
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	// go http.Serve(l, nil)
	log.Fatal(http.Serve(l, nil))
}
