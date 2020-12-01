package main

import (
	//"golang.org/x/net/context"
	"google.golang.org/grpc"

	"fmt"
	"net"
	//"serverdatanode"
	"servernamenode"
	
	//"time"
)
/*
func enviar_a_DataNode1(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	if err_DN1 != nil {
		fmt.Printf("¡Sin conexión DataNode 1!\n")
	} else {
		defer conn_DN1.Close()

		cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
		mensajetest_DN1 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN1 := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)

		if err_DN1 != nil {
			fmt.Printf("> Sin respuesta DataNode1.\n")
		}

	}
}

func enviar_a_DataNode2(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 2
	var conn_DN2 *grpc.ClientConn
	conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
	if err_DN2 != nil {
		fmt.Printf("¡Sin conexión DataNode 2!\n")
	} else {
		defer conn_DN2.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		mensajetest_DN2 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)

		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode2.\n")
		}

	}
}

func enviar_a_DataNode3(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	if err_DN3 != nil {
		fmt.Printf("¡Sin conexión DataNode 3!\n")
	} else {
		defer conn_DN3.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_DN3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN3 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN3)

		if err_DN3 != nil {
			fmt.Printf("> Sin respuesta DataNode3.\n")
		}
	}
}*/

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
