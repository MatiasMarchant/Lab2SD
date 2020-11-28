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
	// peticion_chunk = "Dracula-Stoker_Bram_3" (ejemplo)
	// NO!!!!!!!
	
	partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
	partBuffer := make([]byte, partSize)

	file.Read(partBuffer)

	// Write to disk
	fileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_" + strconv.FormatUint(i, 10)
	_, err4 := os.Create(fileName)

	if err4 != nil {
		log.Fatalf("Error al crear archivo (err4): %s", err4)
	}


	fmt.Printf("Se envía chunk: %s", fileName)

	ChunkLibro := serverdatanode.ChunkLibro{
		Nombre: fileName,
		Chunk:  partBuffer,
	}

	return &ChunkLibro, nil
}