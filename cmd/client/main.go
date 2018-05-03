package main

import (
	"flag"
	"net/rpc"
	"os"

	"github.com/burxtx/gin-microservice-boilerplate/app/client/netrpc"
	"github.com/burxtx/gin-microservice-boilerplate/app/service"
)

func main() {
	var netrpcAddr = flag.String("netrpc.addr", "localhost:8003", "Address for net/rpc server")
	flag.Parse()
	cli, err := rpc.DialHttp("tcp", *netrpcAddr)
	if err != nil {
		os.Exit(1)
	}
	defer cli.Close()
	var svc service.AppService
	svc = netrpc.New()
}
