package main

import (
	"google.golang.org/grpc"

	"fmt"
	"net"
	"servernamenode"

)


func main() {
	// Conexion gRPC
	fmt.Printf("#### NameNode ####\n\n")

	// Escucha en el puerto 900
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("NameNode falla al escuchar puerto 9000: %v", err)
	}
	fmt.Println("NameNode escuchando en puerto 9000.\n")
	s := servernamenode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	servernamenode.RegisterNameNodeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("NameNode falla siendo un servidor gRPC en el puerto 9000: %v", err)
	}
}
