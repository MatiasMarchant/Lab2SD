package servernamenode

import (
	"golang.org/x/net/context"

	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Server struct {
}

func listaDeLibros() string {
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
			listado += n + " " + p_linea + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return listado
}

func ubicacionLibro(tituloLibro string) string {
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

		if !libroEncontrado && i_linea == tituloLibro {
			libroEncontrado = true
			max_chunks, _ = strconv.Atoi(ii_linea)
		} else if libroEncontrado {
			if n_chunk < max_chunks {
				chunks_totales += linea + "\n"
				n_chunk += 1
			} else {
				break
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
		fmt.Printf("> Enviando ubicación de libro '" + tituloLibro + "'\n")
		respuestaDataNode := ubicacionLibro(tituloLibro)
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	}

	respuestaNameNode := "Name node recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaNameNode}, nil
}

// https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *Server) EscribirEnLog(ctx context.Context, message *EscrituraLog) (*MensajeTest, error) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return &MensajeTest{Mensaje: "mal"}, nil
	}
	defer f.Close()

	// Checkear si el libro ya existe, o podemos asumir que el
	// mismo libro no se sube mas de una vez

	// Asumiendo q el mismo libro no se sube mas de una vez:
	// Se escribe Nombre Libro y cant partes
	f.WriteString(message.NombreLibro + " " + strconv.Itoa(int(message.CantPartes)) + "\n")
	PartesDN1 := strings.Split(message.PartesDN1, ",")
	PartesDN2 := strings.Split(message.PartesDN2, ",")
	PartesDN3 := strings.Split(message.PartesDN3, ",")
	// Sacarle el nombre a las partes, para solo escribir el indice
	// Y sortear de menor a mayor

	for indice, palabra := range PartesDN1 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN1[indice] = palabra
	}
	sort.Strings(PartesDN1)
	for indice, palabra := range PartesDN2 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN2[indice] = palabra
	}
	sort.Strings(PartesDN2)
	for indice, palabra := range PartesDN3 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN3[indice] = palabra
	}
	sort.Strings(PartesDN3)

	// Ahora si se puede escribir
	intmessageCantPartes := int(message.CantPartes)
	for i := 0; i < intmessageCantPartes; i++ {
		iastring := strconv.Itoa(i)
		if stringInSlice(iastring, PartesDN1) {
			f.WriteString(iastring + " dist37\n")
		}
		if stringInSlice(iastring, PartesDN2) {
			f.WriteString(iastring + " dist38\n")
		}
		if stringInSlice(iastring, PartesDN3) {
			f.WriteString(iastring + " dist39\n")
		}
	}

	return &MensajeTest{Mensaje: "bien"}, nil
}

func (s *Server) Propuesta_Centralizado(ctx context.Context, Propuesta *Propuestagrpc) (*Propuestagrpc, error) {
	return Propuesta, nil
}