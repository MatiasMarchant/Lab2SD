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


func enviar_a_NameNode(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a NameNode
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		fmt.Printf("¡Sin conexión NameNode!\n")
	} else { 
		defer conn_NN.Close()

		cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

		mensajetestNN := servernamenode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)
		
		if err_NN != nil {
			fmt.Printf("> Sin respuesta NameNode.\n")
		} else {
			fmt.Printf("|Cliente| NameNode responde : %s", respuestaNN.Mensaje)
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

		respuesta_DN2, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)

		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode2.\n")
		} else {
			fmt.Printf("|Cliente| DataNode 2 responde: %s", respuesta_DN2.Mensaje)
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

		cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_3, err_DN3 := cDataNode3.EnvioMensajeTest(context.Background(), &mensajetest_3)

		if err_DN3 != nil {
			fmt.Printf("> Sin respuesta DataNode3.\n")
		} else {
			fmt.Printf("|Cliente| DataNode 3 responde: %s", respuesta_3.Mensaje)
		}

		
	}
}

func main() {
	fmt.Printf("#### DataNode 1 ####\n\n")

	// Escucha en el puerto 9001
	lis, err_s := net.Listen("tcp", ":9001")
	if err_s != nil {
		log.Fatalf("Error en DataNode 1 al escuchar en puerto 9001: %v", err_s)
	}
	fmt.Println("DataNode 1 escuchando en puerto 9001.\n")
	s := serverdatanode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)


	
	go func() {
		for {

			time.Sleep(10 * time.Second)

			mensajeaNN := "Mensaje de prueba DataNode 1 a NameNode\n"
			mensajeaD2 := "Mensaje de prueba DataNode 1 a DataNode 2\n"
			mensajeaD3 := "Mensaje de prueba DataNode 1 a DataNode 3\n"
			enviar_a_NameNode(mensajeaNN)
			enviar_a_DataNode2(mensajeaD2)
			enviar_a_DataNode3(mensajeaD3)			
		}
	}()



	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 1 en servidor gRPC en el puerto 9001: %v", err_s)
	}		



}
