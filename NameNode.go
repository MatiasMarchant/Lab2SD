package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"servernamenode"

	"fmt"
)

func main() {
	// Conexion gRPC
	fmt.Printf("NameNode")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("NameNode falla al escuchar puerto 9000: %v", err)
	}

	s := servernamenode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	servernamenode.RegisterNameNodeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("NameNode falla siendo un servidor gRPC en el puerto 9000: %v", err)
	}
}
