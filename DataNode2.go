package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"log"
	"fmt"
	"net"
	"serverdatanode"
	"servernamenode"
	//"bufio"
	//"os"
)

func main() {
	fmt.Printf("#### DataNode 2 ####\n\n")

	// Ejecutar como Cliente o Servidor
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("¿Ejecutar 'cliente' o 'servidor'?")
	//fmt.Printf("> ")
	//ejecucion, _ = reader.ReadBytes('\n')


	// Escucha en el puerto 9002
	lis, err_s := net.Listen("tcp", ":9002")
	if err_s != nil {
		log.Fatalf("Error en DataNode 2 al escuchar en puerto 9002: %v", err_s)
	}
	fmt.Println("DataNode 2 escuchando en puerto 9002")
	s := serverdatanode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 2 en servidor gRPC en el puerto 9002: %v", err_s)
	}


	//--------------------------------------------------------------------
	// Conexion a NameNode
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		log.Fatalf("Error al conectar DataNode 2 como cliente a NameNode: %s", err_NN)
	}
	defer conn_NN.Close()

	cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

	mensajetestNN := servernamenode.MensajeTest{
		Mensaje: "Mensaje de prueba de Datanode 2\n",
	}

	respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)
	
	if err_NN != nil {
		fmt.Printf("Error: %s", err_NN)
	}

	fmt.Printf("NameNode responde : %s", respuestaNN.Mensaje)

	//#################################################################
	// Esperar para que se abran los otros servidores DataNode
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Servidores listos? (Si/No)")
	//fmt.Printf("> ")
	//_, _ = reader.ReadBytes('\n')

	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	if err_DN1 != nil {
		log.Fatalf("Error al conectar DataNode 2 como cliente a DataNode 1: %s", err_DN1)
	}
	defer conn_DN1.Close()

	cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
	mensajetest_DN1 := serverdatanode.MensajeTest{
		Mensaje: "Mensaje de prueba DataNode 2 a DataNode 1",
	}

	respuesta_DN1, _ := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)
	log.Printf("DataNode 1 responde: %s", respuesta_DN1.Mensaje)

	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	if err_DN3 != nil {
		log.Fatalf("Error al conectar DataNode 2 como cliente a DataNode 3: %s", err_DN3)
	}
	defer conn_DN3.Close()

	cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
	mensajetest_DN1 := serverdatanode.MensajeTest{
		Mensaje: "Mensaje de prueba DataNode 2 a DataNode 3",
	}

	respuesta_DN1, _ := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)
	log.Printf("DataNode 3 responde: %s", respuesta_DN1.Mensaje)

}
