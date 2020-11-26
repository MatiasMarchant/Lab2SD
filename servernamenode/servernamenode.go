package servernamenode

import (
	"golang.org/x/net/context"

	"fmt"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	fmt.Printf("(Servidor) Se recibe: %s", message.Mensaje)
	respuestaNameNode := "Name node recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaNameNode}, nil
}