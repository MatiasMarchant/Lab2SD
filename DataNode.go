package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"servernamenode"

	"fmt"
)

func main() {
	/*

	 */
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error en grpc.Dial(...): %s", err)
	}
	defer conn.Close()

	cNameNode := servernamenode.NewNameNodeServiceClient(conn)

	mensajetest := servernamenode.MensajeTest{
		Mensaje: "Este es un mensaje de pruebac",
	}

	fmt.Printf("????")

	respuesta, err := cNameNode.EnvioMensajeTest(context.Background(), &mensajetest)

	log.Printf("El mensaje de prueba es: %s", respuesta.Mensaje)
}
