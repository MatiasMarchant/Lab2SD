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
	//"strings"
)


func enviar_a_NameNode(mensaje_cliente string) {
		//##########################################################################			
		//--------------------------------------------------------------------
		// Conexion a NameNode
		var conn_NN *grpc.ClientConn
		conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
		if err_NN != nil {
			log.Fatalf("Error al conectar DataNode 1 como cliente a NameNode: %s", err_NN)
		}
		defer conn_NN.Close()

		cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

		mensajetestNN := servernamenode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)
		
		if err_NN != nil {
			fmt.Printf("Error: %s", err_NN)
		}

		fmt.Printf("(Cliente) NameNode responde : %s", respuestaNN.Mensaje)
	}

func enviar_a_DataNode2(mensaje_cliente string) {
		//--------------------------------------------------------------------
		// Conexion a DataNode 2
		var conn_DN2 *grpc.ClientConn
		conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if err_DN2 != nil {
			log.Fatalf("Error al conectar DataNode 1 como cliente a DataNode 2: %s", err_DN2)
		}
		defer conn_DN2.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		mensajetest_DN2 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_DN2, _ := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)
		fmt.Printf("(Cliente) DataNode 2 responde: %s", respuesta_DN2.Mensaje)
	}

func enviar_a_DataNode3(mensaje_cliente string) {
		//--------------------------------------------------------------------
		// Conexion a DataNode 3
		var conn_DN3 *grpc.ClientConn
		conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if err_DN3 != nil {
			log.Fatalf("Error al conectar DataNode 1 como cliente a DataNode 3: %s", err_DN3)
		}
		defer conn_DN3.Close()

		cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_3, _ := cDataNode3.EnvioMensajeTest(context.Background(), &mensajetest_3)
		fmt.Printf("(Cliente) DataNode 3 responde: %s", respuesta_3.Mensaje)
}

func main() {
	fmt.Printf("#### DataNode 1 ####\n\n")


	// Ejecutar como Cliente o Servidor
	/*fmt.Print("¿Ejecutar 'cliente' o 'servidor'?\n")
	fmt.Printf("> ")
    var ejecucion string
    _, err := fmt.Scanln(&ejecucion)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }*/

	//if strings.TrimRight(ejecucion, "\n") == "servidor" {

	//##########################################################################	
	// Escucha en el puerto 9001
	lis, err_s := net.Listen("tcp", ":9001")
	if err_s != nil {
		log.Fatalf("Error en DataNode 1 al escuchar en puerto 9001: %v", err_s)
	}
	fmt.Println("DataNode 1 escuchando en puerto 9001\n")
	s := serverdatanode.Server{}

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 1 en servidor gRPC en el puerto 9001: %v", err_s)
	}		

	
	go func() {
		for {
			fmt.Println("?\n")
			mensajeaNN := "Mensaje de prueba DataNode 1 a NameNode\n"
			mensajeaD2 := "Mensaje de prueba DataNode 1 a DataNode 2\n"
			mensajeaD3 := "Mensaje de prueba DataNode 1 a DataNode 3\n"
			enviar_a_NameNode(mensajeaNN)
			enviar_a_DataNode2(mensajeaD2)
			enviar_a_DataNode3(mensajeaD3)			
		}
	}()


    //} else {


			

    //}



}
