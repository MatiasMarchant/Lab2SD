package serverdatanode

import (
	"golang.org/x/net/context"

	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bufio"
	"strings"
	"strconv"
)

type Server struct {
}


func listaDeLibros() string{
	listado := ""
	nLibro := 0
	file, err := os.Open("log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		linea := scanner.Text()
		p_linea := strings.Split(linea, " ")[0]
		if _, err := strconv.Atoi(v); err != nil {
			nLibro += 1
			n := strconv.Itoa(nLibro)
			listado += n+" "+linea+"\n"
		}		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return listado
}


func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {
	if message.Mensaje == "listadoLibros" {
		fmt.Printf("> Enviando listado de libros")
		respuestaDataNode := listaDeLibros
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	}
	
	fmt.Printf("|Servidor| Se recibe: %s", message.Mensaje)
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
