package serverdatanode

import (
	"golang.org/x/net/context"

	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Server struct {
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	//fmt.Printf("Se recibe: %s", message.Mensaje)
	respuestaDataNode := "DataNode recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaDataNode}, nil
}

func (s *Server) UploaderSubeLibro(ctx context.Context, eddChunkLibro *ChunkLibro) (*MensajeTest, error) {
	fmt.Printf("Se recibe chunk: %s", eddChunkLibro.Nombre)

	fileName := eddChunkLibro.Nombre
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error al crear archivo: %s", err)
	}
	ioutil.WriteFile(fileName, eddChunkLibro.Chunk, os.ModeAppend)

	return &MensajeTest{Mensaje: "retorno"}, nil
}

func (s *Server) DownloaderDescargaLibro(ctx context.Context, peticion_chunk *MensajeTest) (*ChunkLibro, error) {
	ChunkFileName := peticion_chunk.Mensaje
	fmt.Printf("Enviando chunk:  %s", ChunkFileName+"\n")

	newFileChunk, err := os.Open(ChunkFileName)
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