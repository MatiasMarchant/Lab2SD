package servernamenode

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	log.Printf("Recibido message body from client: %s", in.Body)
	return message, nil
}