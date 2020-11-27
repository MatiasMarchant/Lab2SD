package servernamenode

import (
	"golang.org/x/net/context"

	"fmt"
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
		if _, err := strconv.Atoi(p_linea); err != nil {
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
	//if message.Mensaje == "listadoLibros" {
	if strings.Compare(message.Mensaje, "listadoLibros"){
		fmt.Printf("> Enviando listado de libros")
		respuestaDataNode := listaDeLibros()
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	}
	

	//fmt.Printf("|Servidor| Se recibe: %s", message.Mensaje)
	respuestaNameNode := "Name node recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaNameNode}, nil
}