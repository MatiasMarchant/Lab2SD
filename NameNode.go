package main

import (
	"google.golang.org/grpc"
	"net"
	"servernamenode"

	"fmt"
)

func main() {
	// Conexion gRPC
	fmt.Printf("#### NameNode ####\n")

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println("NameNode falla al escuchar puerto 9001: %v", err)
	}

	s := servernamenode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	servernamenode.RegisterNameNodeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("NameNode falla siendo un servidor gRPC en el puerto 9001: %v", err)
	}
}
