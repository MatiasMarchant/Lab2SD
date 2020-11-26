package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"log"
	"fmt"
	"net"
	"serverdatanode"
	"servernamenode"
	
	"time"
)


func enviar_a_DataNode1(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	if err_DN1 != nil {
		fmt.Printf("¡Sin conexión DataNode 1!")
	} else { 
		defer conn_DN1.Close()

		cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
		mensajetest_DN1 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_DN1, err_DN1 := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)

		if err_DN1 != nil {
			fmt.Printf("Sin respuesta DataNode1")
		}

		fmt.Printf("|Cliente| DataNode 1 responde: %s", respuesta_DN1.Mensaje)
	}
}

func enviar_a_DataNode2(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 2
	var conn_DN2 *grpc.ClientConn
	conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
	if err_DN2 != nil {
		fmt.Printf("¡Sin conexión DataNode 2!")
	} else { 
		defer conn_DN2.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		mensajetest_DN2 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_DN2, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)

		if err_DN2 != nil {
			fmt.Printf("Sin respuesta DataNode2")
		}

		fmt.Printf("|Cliente| DataNode 2 responde: %s", respuesta_DN2.Mensaje)
	}
}

func enviar_a_DataNode3(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	if err_DN3 != nil {
		fmt.Printf("¡Sin conexión DataNode 3!")
	} else { 
		defer conn_DN3.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_DN3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_DN3, err_DN3 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN3)

		if err_DN3 != nil {
			fmt.Printf("Sin respuesta DataNode3")
		}
		
		fmt.Printf("|Cliente| DataNode 3 responde: %s", respuesta_DN3.Mensaje)
	}
}



func main() {
	// Conexion gRPC
	fmt.Printf("#### NameNode ####\n\n")

	// Escucha en el puerto 900
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("NameNode falla al escuchar puerto 9000: %v", err)
	}
	fmt.Println("NameNode escuchando en puerto 9000\n")
	s := servernamenode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	servernamenode.RegisterNameNodeServiceServer(grpcServer, &s)

			
	go func() {
		for {

			time.Sleep(2 * time.Second)

			mensajeaD1 := "Mensaje de prueba NameNode a DataNode 1\n"
			mensajeaD2 := "Mensaje de prueba NameNode a DataNode 2\n"
			mensajeaD3 := "Mensaje de prueba NameNode a DataNode 3\n"
			enviar_a_DataNode1(mensajeaD1)
			enviar_a_DataNode2(mensajeaD2)
			enviar_a_DataNode3(mensajeaD3)			
		}
	}()


	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("NameNode falla siendo un servidor gRPC en el puerto 9000: %v", err)
	}
}
