package servernamenode

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	mensajeDeServidor = "Servidor recibe: " + message
	return mensajeDeServidor, nil
}