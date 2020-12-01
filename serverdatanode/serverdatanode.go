package serverdatanode

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

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
	FlagOcupandoLog   bool
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	return &MensajeTest{Mensaje: "respuestaDataNode"}, nil
}

func (s *Server) UploaderSubeLibro(ctx context.Context, eddChunkLibro *ChunkLibro) (*MensajeTest, error) {
	fmt.Printf("> Se recibe chunk: %s\n", eddChunkLibro.Nombre)

	fileName := "Chunks/"+eddChunkLibro.Nombre
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error al crear archivo: %s", err)
	}
	ioutil.WriteFile(fileName, eddChunkLibro.Chunk, os.ModeAppend)

	return &MensajeTest{Mensaje: "Parte subida"}, nil
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
	// Revisar conflictos
	// Que quienes tengan chunks estén vivos
	// Que quienes no tengan chunks no estén vivos

	// ----------------------------------------------------------------------------
	FlagRespuestaDN1 := false
	FlagRespuestaDN2 := false
	FlagRespuestaDN3 := false
	if Propuesta.PartesDN1 != "" { // Si tiene chunks asignados, verificar if vivo
		var connDN1 *grpc.ClientConn
		connDN1, errDN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
		if errDN1 != nil {
			FlagRespuestaDN1 = false
		} else {
			defer connDN1.Close()
			cDataNode1 := NewDataNodeServiceClient(connDN1)
			_, errDN1 := cDataNode1.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN1 != nil {
				FlagRespuestaDN1 = false // Esta muerto
			} else {
				FlagRespuestaDN1 = true // Responde que esta vivo
			}
		}

	} else { // Si no tiene chunks asignados, verificar if not vivo
		var connDN1 *grpc.ClientConn
		connDN1, errDN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
		if errDN1 != nil {
			FlagRespuestaDN1 = true
		} else {
			defer connDN1.Close()
			cDataNode1 := NewDataNodeServiceClient(connDN1)
			_, errDN1 := cDataNode1.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN1 != nil {
				FlagRespuestaDN1 = true // Esta muerto, tonces esta bien que no tenga chunks
			} else {
				FlagRespuestaDN1 = false
			}
		}
	}

	// ----------------------------------------------------------------------------

	if Propuesta.PartesDN2 != "" {
		var connDN2 *grpc.ClientConn
		connDN2, errDN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if errDN2 != nil {
			FlagRespuestaDN2 = false
		} else {
			defer connDN2.Close()
			cDataNode2 := NewDataNodeServiceClient(connDN2)
			_, errDN2 := cDataNode2.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN2 != nil {
				FlagRespuestaDN2 = false
			} else {
				FlagRespuestaDN2 = true
			}
		}
	} else {
		var connDN2 *grpc.ClientConn
		connDN2, errDN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if errDN2 != nil {
			FlagRespuestaDN2 = true
		} else {
			defer connDN2.Close()
			cDataNode2 := NewDataNodeServiceClient(connDN2)
			_, errDN2 := cDataNode2.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN2 != nil {
				FlagRespuestaDN2 = true
			} else {
				FlagRespuestaDN2 = false
			}
		}
	}

	// ----------------------------------------------------------------------------

	if Propuesta.PartesDN3 != "" {
		var connDN3 *grpc.ClientConn
		connDN3, errDN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if errDN3 != nil {
			FlagRespuestaDN3 = false
		} else {
			defer connDN3.Close()
			cDataNode3 := NewDataNodeServiceClient(connDN3)
			_, errDN3 := cDataNode3.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN3 != nil {
				FlagRespuestaDN3 = false
			} else {
				FlagRespuestaDN3 = true
			}
		}

	} else {
		var connDN3 *grpc.ClientConn
		connDN3, errDN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if errDN3 != nil {
			FlagRespuestaDN3 = true
		} else {
			defer connDN3.Close()
			cDataNode3 := NewDataNodeServiceClient(connDN3)
			_, errDN3 := cDataNode3.EnvioMensajeTest(context.Background(), &MensajeTest{Mensaje: "Hola"})
			// fmt.Printf(">>> Mensaje enviado\n")
			if errDN3 != nil {
				FlagRespuestaDN3 = true
			} else {
				FlagRespuestaDN3 = false
			}

		}
	}
	// ----------------------------------------------------------------------------

	RespuestaBooleano := Booleano{Booleano: FlagRespuestaDN1 && FlagRespuestaDN2 && FlagRespuestaDN3}

	return &RespuestaBooleano, nil
}

func (s *Server) DownloaderDescargaLibro(ctx context.Context, peticion_chunk *MensajeTest) (*ChunkLibro, error) {
	ChunkFileName := peticion_chunk.Mensaje
	fmt.Printf("> Enviando chunk:  %s", ChunkFileName+"\n")

	newFileChunk, err := os.Open("Chunks/"+ChunkFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer newFileChunk.Close()
	chunkInfo, err := newFileChunk.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var chunkSize int64 = chunkInfo.Size()
	chunkBufferBytes := make([]byte, chunkSize)

	newFileChunk.Read(chunkBufferBytes)

	ChunkLibro := ChunkLibro{
		Nombre: ChunkFileName,
		Chunk:  chunkBufferBytes,
	}

	return &ChunkLibro, nil
}
