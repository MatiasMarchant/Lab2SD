package main

import (
	//"bufio"
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
)

func main() {
	fmt.Printf("#### ClienteUploader ####\n\n")

	// Conexion a un datanote (ahora lo hago especifico pero deberia ser a uno aleatorio)
	var conn_DN *grpc.ClientConn
	conn_DN, err_DN := grpc.Dial("dist37:9001", grpc.WithInsecure())
	if err_DN != nil {
		log.Fatalf("Error al conectar cliente uploader a (en este caso) DN1: %s", err_DN)
	}
	defer conn_DN.Close()

	cDataNode := serverdatanode.NewDataNodeServiceClient(conn_DN)

	fmt.Print("Ingresar nombre de carpeta donde están libros\n")
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
			log.Fatalf("Error al crear archivo (err4): %s", err4)
		}

		// Write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

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

	/*
		Hasta acá separa los libros en chunks de 250 kB y lo siguiente sería
		enviar chunks a un DataNode, pero voy a rearmarlo para probar que funcione
	*/

	/*
		newFileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_reconstruido" + ".pdf" // Estos trims y + no deberian ser necesario pq despues se reconstruye en otra carpeta
		_, err5 := os.Create(newFileName)

		if err5 != nil {
			log.Fatalf("Error al crear archivo (err5): %s", err5)
		}

		// Set the newFileName file to APPEND MODE!!
		// Open files r and w

		file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			log.Fatalf("Error al abrir archivo: %s", err)
		}

		// IMPORTANTE! do not defer a file.Close when opening a file for APPEND mode!
		// defer file.Close()

		// just information on which part of the new file we are appending
		var writePosition int64 = 0

		for j := uint64(0); j < totalPartsNum; j++ {

			// Read a chunk
			currentChunkFileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_" + strconv.FormatUint(j, 10)

			newFileChunk, err := os.Open(currentChunkFileName)

			if err != nil {
				log.Fatalf("Error al abrir chunk: %s", err)
			}
			defer newFileChunk.Close()

			chunkInfo, err := newFileChunk.Stat()
			if err != nil {
				log.Fatalf("Error al obtener info de newFileChunk: %s", err)
			}

			// Calculate the bytes size of each chunk
			// We are not going to rely on previous data and constant

			var chunkSize int64 = chunkInfo.Size()
			chunkBufferBytes := make([]byte, chunkSize)

			fmt.Println("Appending at position : [", writePosition, "] bytes")
			writePosition = writePosition + chunkSize
			// Read into chunkBufferBytes
			reader := bufio.NewReader(newFileChunk)
			_, err = reader.Read(chunkBufferBytes)

			if err != nil {
				log.Fatalf("Error al leer de chunk: %s", err)
			}

			// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
			// Write/save buffer to disk
			//ioutilWriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

			n, err := file.Write(chunkBufferBytes)

			if err != nil {
				log.Fatalf("Error al escribir en chunkBufferBytes: %s", err)
			}

			file.Sync() // flush to disk

			// Free up the buffer for next cycle
			// Should not be a problem if the chunk size is small, but
			// Can be resource hogging if the chunk size is huge.
			// Also a good practice to clean up your own plate after eating

			chunkBufferBytes = nil // Reset or empty our buffer
			fmt.Println("Written ", n, " bytes")

			fmt.Println("Recombining part [", j, "] into: ", newFileName)
		}


		// Now, we close the newFileName
		file.Close()
	*/
}
