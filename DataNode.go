package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"servernamenode"

	"fmt"
)

func main() {
	/*

	 */

	fmt.Printf("#### DataNode ####\n")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist38:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error en grpc.Dial?(...): %s", err)
	}
	defer conn.Close()

	cNameNode := servernamenode.NewNameNodeServiceClient(conn)

	mensajetest := servernamenode.MensajeTest{
		Mensaje: "Este es un mensaje de pruebac\n",
	}

	respuesta, err := cNameNode.EnvioMensajeTest(context.Background(), &mensajetest)
	
	if err != nil {
		fmt.Printf("Error al llamar: %s", err)
	}

	fmt.Printf("Servidor responde : %s", respuesta.Mensaje)
}
