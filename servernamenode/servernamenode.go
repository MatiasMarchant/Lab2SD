package servernamenode

import (
	"golang.org/x/net/context"

	"fmt"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	fmt.Printf("Cliente envia: %s", message.Mensaje)
	mensajeDeServidor := "Servidor recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: mensajeDeServidor}, nil
}