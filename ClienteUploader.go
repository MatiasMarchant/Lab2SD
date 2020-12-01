package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math"
	"os"
	"serverdatanode"
	"strconv"
	"strings"
	"math/rand"
    "time"
)

func main() {
	fmt.Printf("#### ClienteUploader ####\n\n")
	for {
		// Conexion a un datanote random
		var dist [3]string
		dist[0] = "dist37:9001"
		dist[1] = "dist38:9002"
		dist[2] = "dist39:9003"

		s := rand.NewSource(time.Now().UnixNano())
		random := rand.New(s)
		valor_random := random.Intn(3)

		var conn_DN *grpc.ClientConn
		conn_DN, err_DN := grpc.Dial(dist[valor_random], grpc.WithInsecure())
		if err_DN != nil {
			log.Fatalf("Error al conectar cliente uploader al DataNode: %s", err_DN)
		}
		defer conn_DN.Close()

		cDataNode := serverdatanode.NewDataNodeServiceClient(conn_DN)

	

		fmt.Print("\nIngresar nombre de carpeta donde están libros\n")
		fmt.Print("> ")
		var carpeta_libros string
		_, err := fmt.Scanln(&carpeta_libros)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		carpeta_libros = strings.TrimRight(carpeta_libros, "\n")

		// https://stackoverflow.com/questions/14668850/list-directory-in-go
		files, err_files := ioutil.ReadDir("./" + carpeta_libros + "/")
		if err_files != nil {
			log.Fatal(err_files)
		}

		for indice_libro, f := range files {
			fmt.Print(indice_libro, ". ", f.Name(), "\n")
		}
		fmt.Print("---------------------------------------\n")
		fmt.Print("Ingresar numero de libro a subir\n")
		fmt.Print("> ")
		var indice_libro_a_subir string
		_, err2 := fmt.Scanln(&indice_libro_a_subir)
		if err2 != nil {
			log.Fatalf("Error al recibir variable indice_libro_a_subir: %s", err2)
		}

		integerdice_libro_a_subir, _ := strconv.Atoi(strings.TrimRight(indice_libro_a_subir, "\n"))
		fmt.Println("Libro escogido:", files[integerdice_libro_a_subir].Name())

		// https://www.socketloop.com/tutorials/golang-recombine-chunked-files-example
		fileToBeChunked := "./" + carpeta_libros + "/" + files[integerdice_libro_a_subir].Name()
		file, err3 := os.Open(fileToBeChunked)
		if err3 != nil {
			log.Fatalf("Error al abrir archivo (err3): %s", err3)
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		var fileSize int64 = fileInfo.Size()
		const fileChunk = 250 * (1 << 10) // (1 << 10 equivale a 1 Kilobyte), multiplicado 250 veces
		//const fileChunk = 250000 // 250000 bytes segun propiedades windows, README

		// Calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Dividiento en %d partes.\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {
			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			// Write to disk
			fileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_" + strconv.FormatUint(i, 10)
			_, err4 := os.Create(fileName)

			if err4 != nil {
				log.Fatalf("Error al crear archivo: %s", err4)
			}

			fmt.Println("> Dividiendo en: ", fileName)

			ChunkLibro := serverdatanode.ChunkLibro{
				Nombre: fileName,
				Chunk:  partBuffer,
			}
			respuesta, _ := cDataNode.UploaderSubeLibro(context.Background(), &ChunkLibro)
			fmt.Println(respuesta.Mensaje)

		}

		// Una vez termina de enviar un libro, le comunica al DN que terminó
		// para que así el DN proceda a hacer la propuesta.
		// Asume que el DN esta corriendo
		_, err = cDataNode.UploaderTerminoDeSubirLibro(context.Background(), &serverdatanode.MensajeTest{Mensaje: strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf")})
		if err != nil {
			fmt.Printf("Error al avisar al DN que se subio libro %s: %v\n", files[integerdice_libro_a_subir], err)
		}
	}

}
