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

		_, err_DN2 := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)

		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode1.\n")
		} else {
			//fmt.Printf("|Cliente| DataNode 1 responde: %s", respuesta_DN1.Mensaje)
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
		mensajetest_DN3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN3 := cDataNode3.EnvioMensajeTest(context.Background(), &mensajetest_DN3)

		if err_DN3 != nil {
			fmt.Printf("> Sin respuesta DataNode3.\n")
		} else {
			//fmt.Printf("|Cliente| DataNode 3 responde: %s", respuesta_DN3.Mensaje)
		}

	}
}

func main() {
	fmt.Printf("#### DataNode 2 ####\n\n")

	//##########################################################################
	// Escucha en el puerto 9002
	lis, err_s := net.Listen("tcp", ":9002")
	if err_s != nil {
		log.Fatalf("Error en DataNode 2 al escuchar en puerto 9002: %v", err_s)
	}
	fmt.Println("DataNode 2 escuchando en puerto 9002.\n")
	s := serverdatanode.Server{}

	//s.FlagLibroSubido = false

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	go func() {
		for {
			/*if s.FlagLibroSubido == true {
				// Llamar funcion
				NombreLibroSubido := s.NombreLibroSubido // Va sin ".pdf"
				s.FlagLibroSubido = false
				HacerPropuesta("distribuido", NombreLibroSubido)
			}*/
		}
	}()

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 2 en servidor gRPC en el puerto 9002: %v", err_s)
	}

}
