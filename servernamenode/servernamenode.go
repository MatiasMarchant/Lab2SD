package servernamenode

import (
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	mensajeDeServidor := "Servidor recibe: " + message.Mensaje
	return &MensajeTest{Body: mensajeDeServidor}, nil
}