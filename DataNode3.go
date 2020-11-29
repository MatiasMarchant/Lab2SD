package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"fmt"
	"log"
	"net"
	"serverdatanode"
	"servernamenode"
	
	//"time"
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

		_, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)

		if err_NN != nil {
			fmt.Printf("> Sin respuesta NameNode.\n")
		} else {
			//fmt.Printf("|Cliente| NameNode responde : %s", respuestaNN.Mensaje)
		}

	}
}

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
		} else {
			//fmt.Printf("|Cliente| DataNode 1 responde: %s", respuesta_DN1.Mensaje)
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
		} else {
			//fmt.Printf("|Cliente| DataNode 2 responde: %s", respuesta_DN2.Mensaje)
		}

	}
}

func main() {
	fmt.Printf("#### DataNode 3 ####\n\n")

	//##########################################################################
	// Escucha en el puerto 9003
	lis, err_s := net.Listen("tcp", ":9003")
	if err_s != nil {
		log.Fatalf("Error en DataNode 3 al escuchar en puerto 9003: %v", err_s)
	}
	fmt.Println("DataNode 3 escuchando en puerto 9003.\n")
	s := serverdatanode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	go func() {
		for {

			//time.Sleep(10 * time.Second)

			//mensajeaNN := "Mensaje de prueba DataNode 3 a NameNode\n"
			//mensajeaD1 := "Mensaje de prueba DataNode 3 a DataNode 1\n"
			//mensajeaD2 := "Mensaje de prueba DataNode 3 a DataNode 2\n"
			//enviar_a_NameNode(mensajeaNN)
			//enviar_a_DataNode1(mensajeaD1)
			//enviar_a_DataNode2(mensajeaD2)			
		}
	}()

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 3 en servidor gRPC en el puerto 9003: %v", err_s)
	}

}
