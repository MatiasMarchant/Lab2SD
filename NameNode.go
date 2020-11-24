package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"servernamenode"
)

func main() {
	// Conexion gRPC

	log.Printf("1")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("NameNode falla al escuchar puerto 9000: %v", err)
	}

	log.Printf("2")

	s := servernamenode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	servernamenode.RegisterNameNodeServiceServer(grpcServer, &s)

	log.Printf("3")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("NameNode falla siendo un servidor gRPC en el puerto 9000: %v", err)
	}
}
