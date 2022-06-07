package main

import (
	"log"
	"net"
	"net/http"

	"github.com/junaozun/monitoring-service/services/postservice/impl"
	"github.com/junaozun/monitoring-service/services/postservice/pb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Println("init listener fail ", err)
		return
	}

	server := grpc.NewServer()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	pb.RegisterPostServiceServer(server, &impl.Service{})
	reflection.Register(server)
	server.Serve(listener)

}
