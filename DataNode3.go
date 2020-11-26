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
	fmt.Printf("#### DataNode 1 ####\n\n")

	// Ejecutar como Cliente o Servidor
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("¿Ejecutar 'cliente' o 'servidor'?")
	//fmt.Printf("> ")
	//ejecucion, _ = reader.ReadBytes('\n')


	// Escucha en el puerto 9003
	lis, err_s := net.Listen("tcp", ":9003")
	if err_s != nil {
		log.Fatalf("Error en DataNode 3 al escuchar en puerto 9003: %v", err_s)
	}
	fmt.Println("DataNode 3 escuchando en puerto 9003")
	s := serverdatanode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 3 en servidor gRPC en el puerto 9003: %v", err_s)
	}



	//--------------------------------------------------------------------
	// Conexion a NameNode
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		log.Fatalf("Error al conectar DataNode 3 como cliente a NameNode: %s", err_NN)
	}
	defer conn_NN.Close()

	cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

	mensajetestNN := servernamenode.MensajeTest{
		Mensaje: "Mensaje de prueba de Datanode 3\n",
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
		log.Fatalf("Error al conectar DataNode 3 como cliente a DataNode 1: %s", err_DN1)
	}
	defer conn_DN1.Close()

	cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
	mensajetest_DN1 := serverdatanode.MensajeTest{
		Mensaje: "Mensaje de prueba DataNode 3 a DataNode 1",
	}

	respuesta_DN1, _ := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_DN1)
	log.Printf("DataNode 1 responde: %s", respuesta_DN1.Mensaje)
	

	//--------------------------------------------------------------------
	// Conexion a DataNode 2
	var conn_DN2 *grpc.ClientConn
	conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
	if err_DN2 != nil {
		log.Fatalf("Error al conectar DataNode 3 como cliente a DataNode 2: %s", err_DN2)
	}
	defer conn_DN2.Close()

	cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
	mensajetest_DN2 := serverdatanode.MensajeTest{
		Mensaje: "Mensaje de prueba DataNode 3 a DataNode 2",
	}

	respuesta_DN2, _ := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)
	log.Printf("DataNode 2 responde: %s", respuesta_DN2.Mensaje)
	


}
