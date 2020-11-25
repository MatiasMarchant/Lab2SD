package servernamenode

import (
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*Message, error) {
	mensajeDeServidor := "Servidor recibe: " + message.Mensaje
	return &Message{Body: mensajeDeServidor}, nil
}