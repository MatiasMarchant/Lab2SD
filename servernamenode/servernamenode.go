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
			listado += n+" "+p_linea+"\n"
		}		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return listado
}

func ubicacionLibro(tituloLibro string) string{
	file, err := os.Open("log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	libroEncontrado := false
	n_chunk := 0
	max_chunks := 0
	chunks_totales := ""

    for scanner.Scan() {
		linea := scanner.Text()
		splt_linea := strings.Split(linea, " ")
		i_linea := splt_linea[0]
		ii_linea := splt_linea[1]

		if !libroEncontrado && i_linea == tituloLibro{	
			libroEncontrado = true
			max_chunks, _ = strconv.Atoi(ii_linea)
		} else if libroEncontrado {
			if n_chunk < max_chunks{
				chunks_totales += linea+"\n"
				n_chunk += 1
			} else{
				break;
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	return chunks_totales
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {

	if message.Mensaje == "listadoLibros" {
		fmt.Printf("> Enviando listado de libros\n")
		respuestaDataNode := listaDeLibros()
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	} else if strings.Contains(message.Mensaje, "ubicacion") {
		tituloLibro := strings.Split(message.Mensaje, " ")[1]
		fmt.Printf("> Enviando ubicación de libro '"+ tituloLibro +"'\n")
		respuestaDataNode := ubicacionLibro(tituloLibro)
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	}
	

	//fmt.Printf("|Servidor| Se recibe: %s", message.Mensaje)
	respuestaNameNode := "Name node recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaNameNode}, nil
}