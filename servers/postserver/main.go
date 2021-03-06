package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/junaozun/monitoring-service/servers/postserver/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/postserver/get_post", api.GetPostHandler)

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	server := &http.Server{
		Handler: mux,
	}
	// 本机上所有网卡的9090端口都会监听到
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Println("init listener fail ", err)
		return
	}
	err = server.Serve(listener)
	if err != nil {
		log.Println("init server fail ", err)
		return
	}
	defer server.Shutdown(context.TODO())

}
