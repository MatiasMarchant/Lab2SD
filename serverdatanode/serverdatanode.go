package serverdatanode

import (
	"golang.org/x/net/context"

	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Propuesta struct {
	NombreLibroSubido string
	PartesDN1         []string
	PartesDN2         []string
	PartesDN3         []string // Strings para hacer la concatenacion mas facil
}

type Server struct {
	FlagLibroSubido   bool
	NombreLibroSubido string
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	fmt.Printf("|Servidor| Se recibe: %s", message.Mensaje)
	respuestaDataNode := "DataNode recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaDataNode}, nil
}

func (s *Server) UploaderSubeLibro(ctx context.Context, eddChunkLibro *ChunkLibro) (*MensajeTest, error) {
	fmt.Printf("Se recibe chunk: %s\n", eddChunkLibro.Nombre)

	fileName := eddChunkLibro.Nombre
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error al crear archivo: %s", err)
	}
	ioutil.WriteFile(fileName, eddChunkLibro.Chunk, os.ModeAppend)

	return &MensajeTest{Mensaje: "retorno"}, nil
}

func (s *Server) UploaderTerminoDeSubirLibro(ctx context.Context, NombreLibro *MensajeTest) (*MensajeTest, error) {
	fmt.Printf("Se termino de subir el libro: %s", NombreLibro.Mensaje)
	// Cambiar s.Flag y s.NombreLibroSubido para que en DataNode.go
	// se sepa cuando se termino de subir un libro y se debe comenzar
	// a hacer la propuesta/distribucion
	s.NombreLibroSubido = NombreLibro.Mensaje // Nombre va sin ".pdf"
	s.FlagLibroSubido = true
	return &MensajeTest{Mensaje: "retorno"}, nil
}

func (s *Server) Propuesta_Distribuido(ctx context.Context, Propuesta *Propuestagrpc) (*Booleano, error) {
	// Quizas implementar una probabilidad de que la rechace?
	// Porque cuando se llega aca, ya se sabe si esque los nodos están vivos o muertos (creo)

	return &Booleano{Booleano: false}, nil
}
